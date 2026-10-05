Yes — if you mean **“I want the repository code to stay in the repository, and the service should call repository methods”**, then you should **not write `tx.Create()` directly inside `service.go`**.

For your architecture, make the repositories **transaction-aware**.

### Recommended structure

```text
Subject Service
      │
      ├── SubjectRepository
      │
      └── CoachingSubjectRepository
                │
                ▼
             GORM DB
```

The cleanest approach for your current project is to have the **service start the transaction**, while repositories receive the transaction DB.

### `subject/repository.go`

```go
type SubjectRepository interface {
	FindByCode(ctx context.Context, db *gorm.DB, code string) (*Subject, error)
	Create(ctx context.Context, db *gorm.DB, subject *Subject) error
}
```

Implementation:

```go
func (r *repository) FindByCode(
	ctx context.Context,
	db *gorm.DB,
	code string,
) (*Subject, error) {

	var subject Subject

	err := db.WithContext(ctx).
		Where("code = ?", code).
		First(&subject).Error

	if err != nil {
		return nil, err
	}

	return &subject, nil
}

func (r *repository) Create(
	ctx context.Context,
	db *gorm.DB,
	subject *Subject,
) error {

	return db.WithContext(ctx).
		Create(subject).Error
}
```

### `coaching_subject/repository.go`

```go
type CoachingSubjectRepository interface {
	Create(
		ctx context.Context,
		db *gorm.DB,
		coachingSubject *CoachingSubject,
	) error
}
```

Implementation:

```go
func (r *repository) Create(
	ctx context.Context,
	db *gorm.DB,
	coachingSubject *CoachingSubject,
) error {

	return db.WithContext(ctx).
		Create(coachingSubject).Error
}
```

---

## Then your `service.go`

Now the transaction belongs in the service:

```go
func (s *service) CreateSubject(
	ctx context.Context,
	req *dto.CreateSubjectRequest,
) (*dto.CoachingSubjectResponse, error) {

	var response *CoachingSubject

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		// Check whether subject already exists.
		subject, err := s.subjectRepository.FindByCode(
			ctx,
			tx,
			req.Code,
		)

		// Subject does not exist.
		if errors.Is(err, gorm.ErrRecordNotFound) {

			subject = &Subject{
				Name:        req.Name,
				Code:        req.Code,
				Description: req.Description,
			}

			// Create new Subject.
			if err := s.subjectRepository.Create(
				ctx,
				tx,
				subject,
			); err != nil {
				return err
			}

		} else if err != nil {
			// Unexpected database error.
			return err
		}

		// Subject exists now, either because it was found
		// or because we just created it.
		coachingSubject := &CoachingSubject{
			CoachingID: req.CoachingID,
			SubjectID:  subject.ID,
			Status:     "active",
		}

		// Create CoachingSubject.
		if err := s.coachingSubjectRepository.Create(
			ctx,
			tx,
			coachingSubject,
		); err != nil {
			return err
		}

		response = coachingSubject

		return nil
	})

	if err != nil {
		return nil, err
	}

	return ToCoachingSubjectResponse(response), nil
}
```

### The important part

Notice the difference:

```go
s.subjectRepository.Create(ctx, tx, subject)
```

and:

```go
s.coachingSubjectRepository.Create(ctx, tx, coachingSubject)
```

Here `tx` is the **same transaction**.

So:

```text
BEGIN TRANSACTION
       │
       ├── Find Subject
       │
       ├── Subject exists?
       │       │
       │       ├── YES → don't create Subject
       │       │
       │       └── NO  → create Subject
       │
       └── Create CoachingSubject
                │
          ┌─────┴─────┐
          │           │
       success      error
          │           │
       COMMIT      ROLLBACK
```

This keeps your **business logic in `service.go`** and your **database operations in repositories**, which is much cleaner than putting repository logic directly inside the service.

<!-- Bad way -->

Yes. In **your exact case**, the transaction should be used in the **service layer**, because you potentially create **two records**:

```text
Subject
   +
CoachingSubject
```

Both should succeed, or both should fail.

## 1. Repository methods

Keep repositories simple:

```go
func (r *repository) FindByCode(
	ctx context.Context,
	code string,
) (*Subject, error) {

	var subject Subject

	err := r.db.WithContext(ctx).
		Where("code = ?", code).
		First(&subject).Error

	if err != nil {
		return nil, err
	}

	return &subject, nil
}
```

```go
func (r *repository) Create(
	ctx context.Context,
	subject *Subject,
) error {

	return r.db.WithContext(ctx).
		Create(subject).Error
}
```

And `CoachingSubjectRepository`:

```go
func (r *repository) Create(
	ctx context.Context,
	coachingSubject *CoachingSubject,
) error {

	return r.db.WithContext(ctx).
		Create(coachingSubject).Error
}
```

---

# 2. Where transaction goes

If you simply do this:

```go
Create Subject()
Create CoachingSubject()
```

and the second operation fails, the first record remains.

So put the **whole operation inside one transaction**:

```go
func (s *service) CreateSubject(
	ctx context.Context,
	req *dto.CreateSubjectRequest,
) (*dto.CoachingSubjectResponse, error) {

	var result *CoachingSubject

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		// 1. Find existing subject
		var subject Subject

		err := tx.WithContext(ctx).
			Where("code = ?", req.Code).
			First(&subject).Error

		// 2. Subject doesn't exist
		if errors.Is(err, gorm.ErrRecordNotFound) {

			subject = Subject{
				Name:        req.Name,
				Code:        req.Code,
				Description: req.Description,
			}

			// Create Subject
			if err := tx.WithContext(ctx).
				Create(&subject).Error; err != nil {
				return err
			}

		} else if err != nil {
			// Some unexpected database error
			return err
		}

		// 3. Subject now exists.
		// Create CoachingSubject.
		coachingSubject := &CoachingSubject{
			CoachingID: req.CoachingID,
			SubjectID:  subject.ID,
			Status:     "active",
		}

		if err := tx.WithContext(ctx).
			Create(coachingSubject).Error; err != nil {
			return err
		}

		result = coachingSubject

		return nil
	})

	if err != nil {
		return nil, err
	}

	return ToCoachingSubjectResponse(result), nil
}
```

You need:

```go
import (
	"context"
	"errors"

	"gorm.io/gorm"
)
```

---

# 3. But there is an architecture issue

If your service currently has:

```go
type service struct {
	repository SubjectRepository
}
```

then you can't directly do:

```go
s.db.Transaction(...)
```

unless the service has access to the DB.

For your project, I'd structure this operation around a **unit-of-work / transaction-aware repository**, rather than exposing `*gorm.DB` everywhere.

For learning and your current project, however, you can initially do:

```go
type service struct {
	db                     *gorm.DB
	subjectRepository      SubjectRepository
	coachingSubjectRepository CoachingSubjectRepository
}
```

Then:

```go
err := s.db.Transaction(func(tx *gorm.DB) error {
	// all database operations here
	return nil
})
```

### The important concept

Inside the transaction, **don't accidentally use the normal `r.db`**:

❌ Wrong:

```go
s.subjectRepository.Create(ctx, subject)
```

if that repository internally uses:

```go
r.db.Create(subject)
```

because that may use the original DB connection rather than your transaction.

Instead, the operations inside the transaction must use:

```go
tx.Create(...)
```

or transaction-aware repositories.

---

## Your final flow

```text
POST /coaching-subjects
          │
          ▼
       Service
          │
          ▼
      BEGIN TX
          │
          ▼
   Find Subject by Code
       │          │
    EXISTS     NOT EXISTS
       │          │
       │          ▼
       │      Create Subject
       │          │
       └────┬─────┘
            ▼
   Create CoachingSubject
            │
       ┌────┴────┐
       │         │
     SUCCESS   ERROR
       │         │
     COMMIT    ROLLBACK
```

So if `CoachingSubject` creation fails after creating a new `Subject`, the newly created `Subject` is also **rolled back**. This is exactly where a transaction is useful.
