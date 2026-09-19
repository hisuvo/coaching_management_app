হ্যাঁ। তোমার **Coaching Management System**-এর জন্য আমি JWT authentication-কে একদম fresh design হিসেবে ধরছি এবং আগের token/JWT design-এর উপর নির্ভর করছি না।

তোমার project যেহেতু **Go + Echo + GORM + PostgreSQL** এবং future-এ **multi-coaching SaaS** হবে, তাই আমি এই architecture recommend করব:

- **Access Token** → short-lived JWT, যেমন 15 minutes
- **Refresh Token** → long-lived, random opaque token, যেমন 7 days
- Refresh token **database-এ hash করে** রাখা হবে
- Refresh token rotation থাকবে
- Logout করলে session revoke হবে
- Password → bcrypt/Argon2id hash
- JWT payload-এ password বা sensitive data থাকবে না
- `User` এবং `Auth/Session` আলাদা domain
- `CoachingID` দিয়ে tenant isolation
- Middleware protected route-এ authenticated user context inject করবে
- Access token → `Authorization: Bearer <token>`
- Refresh token → **HttpOnly + Secure cookie**
- Production-এ HTTPS বাধ্যতামূলক

---

# 1. Recommended folder structure

তোমার `internal/domain` structure ধরে আমি এভাবে করতাম:

```text
backend/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   │
│   ├── domain/
│   │
│   │   ├── auth/
│   │   │   ├── entity.go
│   │   │   ├── dto/
│   │   │   │   ├── request.go
│   │   │   │   └── response.go
│   │   │   ├── repository.go
│   │   │   ├── service.go
│   │   │   ├── handler.go
│   │   │   ├── middleware.go
│   │   │   ├── token.go
│   │   │   ├── password.go
│   │   │   ├── mapper.go
│   │   │   └── register.go
│   │   │
│   │   ├── users/
│   │   │   ├── entity.go
│   │   │   ├── dto/
│   │   │   ├── repository.go
│   │   │   ├── service.go
│   │   │   ├── handler.go
│   │   │   ├── mapper.go
│   │   │   └── register.go
│   │   │
│   │   ├── coachings/
│   │   ├── students/
│   │   ├── teachers/
│   │   ├── subjects/
│   │   ├── batches/
│   │   └── ...
│   │
│   ├── config/
│   │   └── config.go
│   │
│   ├── database/
│   │   └── database.go
│   │
│   └── server/
│       └── server.go
│
├── pkg/
│   ├── appresponse/
│   ├── apperror/
│   └── ...
│
├── .env
├── go.mod
└── go.sum
```

### Important separation

```text
users
   ↓
User identity / profile / role

auth
   ↓
Password authentication
JWT
Refresh token
Sessions
Login
Logout
Token refresh
Middleware
```

অর্থাৎ:

> **User = কে user?**
> **Auth = user কীভাবে authenticate করবে এবং session/token কীভাবে manage হবে?**

---

# 2. Dependencies

আমি JWT-এর জন্য `golang-jwt/jwt/v5` ব্যবহার করব।

```bash
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
```

---

# 3. Environment variables

`.env`

```env
APP_ENV=development

JWT_ACCESS_SECRET=replace-with-a-long-random-secret
JWT_ISSUER=coaching-management-api

JWT_ACCESS_EXPIRES=15m
JWT_REFRESH_EXPIRES=168h

AUTH_COOKIE_NAME=refresh_token
AUTH_COOKIE_SECURE=false
AUTH_COOKIE_DOMAIN=
```

Production:

```env
APP_ENV=production

JWT_ACCESS_SECRET=<strong-random-secret>
JWT_ISSUER=coaching-management-api

JWT_ACCESS_EXPIRES=15m
JWT_REFRESH_EXPIRES=168h

AUTH_COOKIE_NAME=refresh_token
AUTH_COOKIE_SECURE=true
AUTH_COOKIE_DOMAIN=.yourdomain.com
```

**JWT secret কখনো GitHub-এ commit করবে না।**

---

# 4. User entity

`internal/domain/users/entity.go`

```go
package users

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	RoleSuperAdmin UserRole = "SUPER_ADMIN"
	RoleAdmin      UserRole = "ADMIN"
	RoleTeacher    UserRole = "TEACHER"
	RoleStudent    UserRole = "STUDENT"
)

type UserStatus string

const (
	StatusActive  UserStatus = "ACTIVE"
	StatusInactive UserStatus = "INACTIVE"
	StatusBlocked  UserStatus = "BLOCKED"
)

type User struct {
	ID uint `gorm:"primaryKey"`

	CoachingID *uint `gorm:"index"`

	Name string `gorm:"type:varchar(100);not null"`

	Email string `gorm:"type:varchar(150);uniqueIndex;not null"`

	PasswordHash string `gorm:"type:varchar(255);not null"`

	Role UserRole `gorm:"type:varchar(30);not null;index"`

	Status UserStatus `gorm:"type:varchar(30);not null;default:'ACTIVE';index"`

	Phone string `gorm:"type:varchar(20)"`

	LastLoginAt *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

### কেন `PasswordHash`?

Database-এ কখনো:

```text
password = "123456"
```

রাখবে না।

বরং:

```text
password_hash = "$2a$..."
```

রাখবে।

---

# 5. Auth Session entity

এটা authentication-এর সবচেয়ে গুরুত্বপূর্ণ অংশগুলোর একটি।

`internal/domain/auth/entity.go`

```go
package auth

import (
	"time"

	"gorm.io/gorm"
)

type AuthSession struct {
	ID uint `gorm:"primaryKey"`

	UserID uint `gorm:"not null;index"`

	CoachingID *uint `gorm:"index"`

	RefreshTokenHash string `gorm:"type:varchar(255);uniqueIndex;not null"`

	UserAgent string `gorm:"type:text"`

	IPAddress string `gorm:"type:varchar(45)"`

	ExpiresAt time.Time `gorm:"not null;index"`

	RevokedAt *time.Time `gorm:"index"`

	CreatedAt time.Time

	UpdatedAt time.Time

	DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

### কেন database session দরকার?

শুধু JWT ব্যবহার করলে:

```text
Login
 ↓
JWT তৈরি
 ↓
JWT valid
```

কিন্তু server থেকে নির্দিষ্ট token/session revoke করা কঠিন।

Database session থাকলে:

```text
User
 ↓
Session
 ↓
Refresh Token
```

তুমি চাইলে:

```text
Logout
Force logout
All devices logout
Admin session revoke
Expired session cleanup
```

করতে পারবে।

---

# 6. Auth request DTO

`internal/domain/auth/dto/request.go`

```go
package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken,omitempty"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refreshToken,omitempty"`
}
```

যদি refresh token HttpOnly cookie-তে রাখো, তাহলে `refreshToken` body-তে পাঠানোর প্রয়োজন নেই।

---

# 7. Auth response DTO

`internal/domain/auth/dto/response.go`

```go
package dto

import "time"

type UserResponse struct {
	ID         uint    `json:"id"`
	CoachingID *uint   `json:"coachingId,omitempty"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Role       string  `json:"role"`
	Status     string  `json:"status"`
	Phone      string  `json:"phone,omitempty"`
}

type LoginResponse struct {
	AccessToken string       `json:"accessToken"`
	TokenType   string       `json:"tokenType"`
	ExpiresIn   int64        `json:"expiresIn"`
	User        UserResponse `json:"user"`
}

type RefreshResponse struct {
	AccessToken string `json:"accessToken"`
	TokenType   string `json:"tokenType"`
	ExpiresIn   int64  `json:"expiresIn"`
}

type SessionResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"userId"`
	IPAddress string    `json:"ipAddress,omitempty"`
	UserAgent string    `json:"userAgent,omitempty"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
}
```

---

# 8. JWT Claims

`internal/domain/auth/token.go`

```go
package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaims struct {
	UserID     uint    `json:"uid"`
	CoachingID *uint   `json:"cid,omitempty"`
	Role       string  `json:"role"`
	SessionID  uint    `json:"sid"`
	TokenType  string  `json:"typ"`

	jwt.RegisteredClaims
}
```

### JWT payload হবে conceptually:

```json
{
  "uid": 10,
  "cid": 2,
  "role": "ADMIN",
  "sid": 25,
  "typ": "access",
  "iss": "coaching-management-api",
  "exp": 1789490000,
  "iat": 1789489100
}
```

Password এখানে থাকবে না।

---

# 9. Token manager

`internal/domain/auth/token.go`

```go
package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager struct {
	AccessSecret []byte

	Issuer        string

	AccessExpires time.Duration

	RefreshExpires time.Duration
}

func NewTokenManager(
	accessSecret string,
	issuer string,
	accessExpires time.Duration,
	refreshExpires time.Duration,
) *TokenManager {
	return &TokenManager{
		AccessSecret:   []byte(accessSecret),
		Issuer:         issuer,
		AccessExpires:  accessExpires,
		RefreshExpires: refreshExpires,
	}
}
```

---

# 10. Generate access token

Same file:

```go
func (tm *TokenManager) GenerateAccessToken(
	userID uint,
	coachingID *uint,
	role string,
	sessionID uint,
) (string, int64, error) {

	now := time.Now()
	expiresAt := now.Add(tm.AccessExpires)

	claims := AccessTokenClaims{
		UserID:     userID,
		CoachingID: coachingID,
		Role:       role,
		SessionID:  sessionID,
		TokenType:  "access",

		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    tm.Issuer,
			Subject:   string(rune(userID)),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString(tm.AccessSecret)
	if err != nil {
		return "", 0, err
	}

	return signedToken, int64(tm.AccessExpires.Seconds()), nil
}
```

একটা ছোট improvement: `Subject`-এর জন্য `strconv.FormatUint` ব্যবহার করা production-এ আরও পরিষ্কার।

তাই final version:

```go
import "strconv"
```

এবং:

```go
Subject: strconv.FormatUint(uint64(userID), 10),
```

ব্যবহার করবে।

---

# 11. Verify access token

```go
func (tm *TokenManager) VerifyAccessToken(
	tokenString string,
) (*AccessTokenClaims, error) {

	token, err := jwt.ParseWithClaims(
		tokenString,
		&AccessTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {

			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("unexpected signing method")
			}

			return tm.AccessSecret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid access token")
	}

	if claims.TokenType != "access" {
		return nil, errors.New("invalid token type")
	}

	return claims, nil
}
```

---

# 12. Generate refresh token

Refresh token JWT না করে random token করা better।

```go
func GenerateRefreshToken() (string, error) {

	bytes := make([]byte, 64)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(bytes), nil
}
```

Result:

```text
d9QkJ5...very-long-random-value...
```

---

# 13. Hash refresh token

`token.go`

```go
package auth

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashRefreshToken(token string) string {

	hash := sha256.Sum256([]byte(token))

	return hex.EncodeToString(hash[:])
}
```

Database-এ:

```text
raw refresh token
```

রাখবে না।

বরং:

```text
SHA256(refresh_token)
```

রাখবে।

---

# 14. Password service

`internal/domain/auth/password.go`

```go
package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	if password == "" {
		return "", errors.New("password is required")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}
```

Verify:

```go
func ComparePassword(
	hash string,
	password string,
) error {

	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}
```

---

# 15. Repository interface

`internal/domain/auth/repository.go`

```go
package auth

import (
	"context"
	"time"
)

type Repository interface {

	CreateSession(
		ctx context.Context,
		session *AuthSession,
	) error

	GetSessionByTokenHash(
		ctx context.Context,
		tokenHash string,
	) (*AuthSession, error)

	RevokeSession(
		ctx context.Context,
		sessionID uint,
	) error

	RevokeAllUserSessions(
		ctx context.Context,
		userID uint,
	) error

	UpdateRefreshToken(
		ctx context.Context,
		sessionID uint,
		tokenHash string,
		expiresAt time.Time,
	) error
}
```

---

# 16. Repository implementation

`repository.go`

```go
package auth

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateSession(
	ctx context.Context,
	session *AuthSession,
) error {

	return r.db.WithContext(ctx).
		Create(session).
		Error
}
```

Get session:

```go
func (r *repository) GetSessionByTokenHash(
	ctx context.Context,
	tokenHash string,
) (*AuthSession, error) {

	var session AuthSession

	err := r.db.WithContext(ctx).
		Where(
			"refresh_token_hash = ?",
			tokenHash,
		).
		First(&session).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &session, nil
}
```

---

# 17. Revoke session

```go
func (r *repository) RevokeSession(
	ctx context.Context,
	sessionID uint,
) error {

	now := time.Now()

	return r.db.WithContext(ctx).
		Model(&AuthSession{}).
		Where("id = ?", sessionID).
		Update("revoked_at", now).
		Error
}
```

---

# 18. Revoke all sessions

```go
func (r *repository) RevokeAllUserSessions(
	ctx context.Context,
	userID uint,
) error {

	now := time.Now()

	return r.db.WithContext(ctx).
		Model(&AuthSession{}).
		Where(
			"user_id = ? AND revoked_at IS NULL",
			userID,
		).
		Update("revoked_at", now).
		Error
}
```

---

# 19. User repository

`internal/domain/users/repository.go`

```go
package users

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {

	Create(
		ctx context.Context,
		user *User,
	) error

	GetByEmail(
		ctx context.Context,
		email string,
	) (*User, error)

	GetByID(
		ctx context.Context,
		id uint,
	) (*User, error)

	UpdateLastLogin(
		ctx context.Context,
		id uint,
	) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}
```

Get email:

```go
func (r *repository) GetByEmail(
	ctx context.Context,
	email string,
) (*User, error) {

	var user User

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
```

Get user:

```go
func (r *repository) GetByID(
	ctx context.Context,
	id uint,
) (*User, error) {

	var user User

	err := r.db.WithContext(ctx).
		First(&user, id).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
```

Last login:

```go
func (r *repository) UpdateLastLogin(
	ctx context.Context,
	id uint,
) error {

	now := time.Now()

	return r.db.WithContext(ctx).
		Model(&User{}).
		Where("id = ?", id).
		Update("last_login_at", now).
		Error
}
```

---

# 20. Auth service interface

`internal/domain/auth/service.go`

```go
package auth

import (
	"context"

	"github.com/your-project/internal/domain/auth/dto"
)

type Service interface {

	Login(
		ctx context.Context,
		req *dto.LoginRequest,
		userAgent string,
		ipAddress string,
	) (*LoginResult, error)

	Refresh(
		ctx context.Context,
		refreshToken string,
		userAgent string,
		ipAddress string,
	) (*RefreshResult, error)

	Logout(
		ctx context.Context,
		refreshToken string,
	) error

	LogoutAll(
		ctx context.Context,
		userID uint,
	) error
}
```

`your-project` তোমার actual module name দিয়ে replace করবে।

---

# 21. Login result

```go
type LoginResult struct {
	AccessToken string

	ExpiresIn int64

	RefreshToken string

	UserID uint

	CoachingID *uint

	Role string

	Name string

	Email string
}
```

Refresh:

```go
type RefreshResult struct {
	AccessToken string

	ExpiresIn int64

	RefreshToken string
}
```

---

# 22. Login service

```go
func (s *service) Login(
	ctx context.Context,
	req *dto.LoginRequest,
	userAgent string,
	ipAddress string,
) (*LoginResult, error) {

	user, err := s.userRepository.GetByEmail(
		ctx,
		req.Email,
	)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if user.Status != users.StatusActive {
		return nil, errors.New("user account is not active")
	}

	if err := ComparePassword(
		user.PasswordHash,
		req.Password,
	); err != nil {
		return nil, errors.New("invalid email or password")
	}

	refreshToken, err := GenerateRefreshToken()

	if err != nil {
		return nil, err
	}

	refreshTokenHash := HashRefreshToken(refreshToken)

	session := &AuthSession{
		UserID:            user.ID,
		CoachingID:        user.CoachingID,
		RefreshTokenHash:  refreshTokenHash,
		UserAgent:         userAgent,
		IPAddress:         ipAddress,
		ExpiresAt:         time.Now().Add(
			s.tokenManager.RefreshExpires,
		),
	}

	if err := s.authRepository.CreateSession(
		ctx,
		session,
	); err != nil {
		return nil, err
	}

	accessToken, expiresIn, err :=
		s.tokenManager.GenerateAccessToken(
			user.ID,
			user.CoachingID,
			string(user.Role),
			session.ID,
		)

	if err != nil {
		return nil, err
	}

	if err := s.userRepository.UpdateLastLogin(
		ctx,
		user.ID,
	); err != nil {
		return nil, err
	}

	return &LoginResult{
		AccessToken:  accessToken,
		ExpiresIn:    expiresIn,
		RefreshToken: refreshToken,
		UserID:       user.ID,
		CoachingID:   user.CoachingID,
		Role:         string(user.Role),
		Name:         user.Name,
		Email:        user.Email,
	}, nil
}
```

---

# 23. Service struct

```go
type service struct {
	userRepository users.Repository

	authRepository Repository

	tokenManager *TokenManager
}
```

Constructor:

```go
func NewService(
	userRepository users.Repository,
	authRepository Repository,
	tokenManager *TokenManager,
) Service {

	return &service{
		userRepository: userRepository,
		authRepository: authRepository,
		tokenManager:   tokenManager,
	}
}
```

---

# 24. Refresh token service

এখানেই **refresh token rotation** হবে।

```go
func (s *service) Refresh(
	ctx context.Context,
	refreshToken string,
	userAgent string,
	ipAddress string,
) (*RefreshResult, error) {

	if refreshToken == "" {
		return nil, errors.New("refresh token is required")
	}

	tokenHash := HashRefreshToken(refreshToken)

	session, err :=
		s.authRepository.GetSessionByTokenHash(
			ctx,
			tokenHash,
		)

	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	if session.RevokedAt != nil {
		return nil, errors.New("refresh session has been revoked")
	}

	if time.Now().After(session.ExpiresAt) {
		return nil, errors.New("refresh token has expired")
	}

	user, err := s.userRepository.GetByID(
		ctx,
		session.UserID,
	)

	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.Status != users.StatusActive {
		return nil, errors.New("user account is not active")
	}

	newRefreshToken, err := GenerateRefreshToken()

	if err != nil {
		return nil, err
	}

	newHash := HashRefreshToken(newRefreshToken)

	newExpiresAt := time.Now().Add(
		s.tokenManager.RefreshExpires,
	)

	if err := s.authRepository.UpdateRefreshToken(
		ctx,
		session.ID,
		newHash,
		newExpiresAt,
	); err != nil {
		return nil, err
	}

	accessToken, expiresIn, err :=
		s.tokenManager.GenerateAccessToken(
			user.ID,
			user.CoachingID,
			string(user.Role),
			session.ID,
		)

	if err != nil {
		return nil, err
	}

	return &RefreshResult{
		AccessToken:  accessToken,
		ExpiresIn:    expiresIn,
		RefreshToken: newRefreshToken,
	}, nil
}
```

এটা গুরুত্বপূর্ণ:

```text
Refresh Token A
       ↓
Refresh request
       ↓
Refresh Token A invalidated/replaced
       ↓
Refresh Token B
```

এটাই rotation।

---

# 25. Update refresh token

```go
func (r *repository) UpdateRefreshToken(
	ctx context.Context,
	sessionID uint,
	tokenHash string,
	expiresAt time.Time,
) error {

	return r.db.WithContext(ctx).
		Model(&AuthSession{}).
		Where("id = ?", sessionID).
		Updates(map[string]interface{}{
			"refresh_token_hash": tokenHash,
			"expires_at":         expiresAt,
		}).
		Error
}
```

---

# 26. Logout

```go
func (s *service) Logout(
	ctx context.Context,
	refreshToken string,
) error {

	if refreshToken == "" {
		return nil
	}

	tokenHash := HashRefreshToken(refreshToken)

	session, err :=
		s.authRepository.GetSessionByTokenHash(
			ctx,
			tokenHash,
		)

	if err != nil {
		return nil
	}

	return s.authRepository.RevokeSession(
		ctx,
		session.ID,
	)
}
```

---

# 27. Logout all devices

```go
func (s *service) LogoutAll(
	ctx context.Context,
	userID uint,
) error {

	return s.authRepository.RevokeAllUserSessions(
		ctx,
		userID,
	)
}
```

---

# 28. HTTP cookie configuration

`handler.go`

```go
func (h *handler) setRefreshCookie(
	c echo.Context,
	token string,
	expiresAt time.Time,
) {

	cookie := new(http.Cookie)

	cookie.Name = h.config.CookieName
	cookie.Value = token

	cookie.Path = "/"

	cookie.HttpOnly = true

	cookie.Secure = h.config.CookieSecure

	cookie.SameSite = http.SameSiteLaxMode

	cookie.Expires = expiresAt

	cookie.MaxAge = int(
		time.Until(expiresAt).Seconds(),
	)

	http.SetCookie(c.Response(), cookie)
}
```

Production HTTPS-এর জন্য:

```text
HttpOnly = true
Secure   = true
SameSite = Lax/None
```

---

# 29. Clear cookie

```go
func (h *handler) clearRefreshCookie(
	c echo.Context,
) {

	cookie := new(http.Cookie)

	cookie.Name = h.config.CookieName

	cookie.Value = ""

	cookie.Path = "/"

	cookie.HttpOnly = true

	cookie.Secure = h.config.CookieSecure

	cookie.MaxAge = -1

	cookie.Expires = time.Unix(0, 0)

	http.SetCookie(c.Response(), cookie)
}
```

---

# 30. Login handler

Echo v5-এ handler signature রাখবে:

```go
func (h *handler) Login(c echo.Context) error {
```

**`*echo.Context` ব্যবহার করবে না।**

```go
func (h *handler) Login(c echo.Context) error {

	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return apperror.BadRequest(
			"invalid request body",
		)
	}

	if err := c.Validate(&req); err != nil {
		return apperror.BadRequest(
			"validation failed",
		)
	}

	result, err := h.service.Login(
		c.Request().Context(),
		&req,
		c.Request().UserAgent(),
		c.RealIP(),
	)

	if err != nil {
		return apperror.Unauthorized(
			"invalid email or password",
		)
	}

	expiresAt := time.Now().Add(
		h.config.RefreshExpires,
	)

	h.setRefreshCookie(
		c,
		result.RefreshToken,
		expiresAt,
	)

	response := dto.LoginResponse{
		AccessToken: result.AccessToken,
		TokenType:   "Bearer",
		ExpiresIn:   result.ExpiresIn,

		User: dto.UserResponse{
			ID:         result.UserID,
			CoachingID: result.CoachingID,
			Name:       result.Name,
			Email:      result.Email,
			Role:       result.Role,
			Status:     "ACTIVE",
		},
	}

	return c.JSON(
		http.StatusOK,
		response,
	)
}
```

---

# 31. Login API

```http
POST /api/v1/auth/login
```

Request:

```json
{
  "email": "admin@coaching.com",
  "password": "StrongPassword123"
}
```

Response:

```json
{
  "accessToken": "eyJhbGciOiJIUzI1NiIs...",
  "tokenType": "Bearer",
  "expiresIn": 900,
  "user": {
    "id": 1,
    "coachingId": 1,
    "name": "Admin",
    "email": "admin@coaching.com",
    "role": "ADMIN",
    "status": "ACTIVE"
  }
}
```

আর browser-এ:

```text
HttpOnly Cookie
    refresh_token=....
```

---

# 32. Refresh handler

```go
func (h *handler) Refresh(c echo.Context) error {

	cookie, err := c.Cookie(
		h.config.CookieName,
	)

	if err != nil {
		return apperror.Unauthorized(
			"refresh token is missing",
		)
	}

	result, err := h.service.Refresh(
		c.Request().Context(),
		cookie.Value,
		c.Request().UserAgent(),
		c.RealIP(),
	)

	if err != nil {

		h.clearRefreshCookie(c)

		return apperror.Unauthorized(
			"invalid refresh token",
		)
	}

	expiresAt := time.Now().Add(
		h.config.RefreshExpires,
	)

	h.setRefreshCookie(
		c,
		result.RefreshToken,
		expiresAt,
	)

	return c.JSON(
		http.StatusOK,
		dto.RefreshResponse{
			AccessToken: result.AccessToken,
			TokenType:   "Bearer",
			ExpiresIn:   result.ExpiresIn,
		},
	)
}
```

Endpoint:

```http
POST /api/v1/auth/refresh
```

Frontend-কে refresh token manually read করতে হবে না।

---

# 33. Logout handler

```go
func (h *handler) Logout(c echo.Context) error {

	cookie, err := c.Cookie(
		h.config.CookieName,
	)

	if err == nil && cookie.Value != "" {

		_ = h.service.Logout(
			c.Request().Context(),
			cookie.Value,
		)
	}

	h.clearRefreshCookie(c)

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"success": true,
			"message": "logged out successfully",
		},
	)
}
```

---

# 34. JWT Middleware

এটা protected API-এর জন্য।

`middleware.go`

```go
package auth

import (
	"strings"

	"github.com/labstack/echo/v5"
)

const (
	ContextUserID     = "auth.user_id"
	ContextCoachingID = "auth.coaching_id"
	ContextRole       = "auth.role"
	ContextSessionID  = "auth.session_id"
)

func (h *handler) AuthMiddleware(
	next echo.HandlerFunc,
) echo.HandlerFunc {

	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get(
			"Authorization",
		)

		if authHeader == "" {
			return apperror.Unauthorized(
				"authorization header is required",
			)
		}

		parts := strings.SplitN(
			authHeader,
			" ",
			2,
		)

		if len(parts) != 2 ||
			!strings.EqualFold(parts[0], "Bearer") {

			return apperror.Unauthorized(
				"invalid authorization header",
			)
		}

		claims, err :=
			h.tokenManager.VerifyAccessToken(parts[1])

		if err != nil {
			return apperror.Unauthorized(
				"invalid or expired access token",
			)
		}

		c.Set(
			ContextUserID,
			claims.UserID,
		)

		c.Set(
			ContextCoachingID,
			claims.CoachingID,
		)

		c.Set(
			ContextRole,
			claims.Role,
		)

		c.Set(
			ContextSessionID,
			claims.SessionID,
		)

		return next(c)
	}
}
```

---

# 35. Protected route

```go
authGroup := e.Group("/api/v1")

authGroup.Use(h.AuthMiddleware)

authGroup.GET(
	"/profile",
	h.Profile,
)
```

Request:

```http
GET /api/v1/profile
Authorization: Bearer eyJhbGciOiJIUzI1Ni...
```

---

# 36. Get authenticated user ID

একটা helper রাখো।

`middleware.go`

```go
func GetUserID(c echo.Context) (uint, bool) {

	value := c.Get(ContextUserID)

	userID, ok := value.(uint)

	return userID, ok
}
```

Coaching:

```go
func GetCoachingID(
	c echo.Context,
) (*uint, bool) {

	value := c.Get(ContextCoachingID)

	if value == nil {
		return nil, true
	}

	coachingID, ok := value.(*uint)

	return coachingID, ok
}
```

Role:

```go
func GetUserRole(c echo.Context) (string, bool) {

	value := c.Get(ContextRole)

	role, ok := value.(string)

	return role, ok
}
```

---

# 37. Role middleware

Authentication আর authorization এক জিনিস না।

### Authentication

```text
Who are you?
```

### Authorization

```text
What are you allowed to do?
```

তাই role middleware আলাদা রাখবে।

```go
func RequireRoles(
	roles ...string,
) echo.MiddlewareFunc {

	allowed := make(map[string]struct{})

	for _, role := range roles {
		allowed[role] = struct{}{}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			role, ok := GetUserRole(c)

			if !ok {
				return apperror.Unauthorized(
					"authentication required",
				)
			}

			if _, exists := allowed[role]; !exists {
				return apperror.Forbidden(
					"you do not have permission",
				)
			}

			return next(c)
		}
	}
}
```

---

# 38. Example admin route

```go
adminGroup := e.Group(
	"/api/v1/admin",
)

adminGroup.Use(
	h.AuthMiddleware,
)

adminGroup.Use(
	RequireRoles("SUPER_ADMIN", "ADMIN"),
)

adminGroup.GET(
	"/dashboard",
	h.Dashboard,
)
```

এখন:

```text
Unauthenticated
       ↓
401 Unauthorized

Authenticated but STUDENT
       ↓
403 Forbidden

Authenticated ADMIN
       ↓
200 OK
```

---

# 39. Register route

`register.go`

```go
package auth

import (
	"github.com/labstack/echo/v5"
)

func RegisterRoutes(
	e *echo.Echo,
	handler *handler,
) {

	auth := e.Group("/api/v1/auth")

	auth.POST(
		"/login",
		handler.Login,
	)

	auth.POST(
		"/refresh",
		handler.Refresh,
	)

	auth.POST(
		"/logout",
		handler.Logout,
	)
}
```

Protected auth routes:

```go
protected := e.Group("/api/v1/auth")

protected.Use(handler.AuthMiddleware)

protected.POST(
	"/logout-all",
	handler.LogoutAll,
)
```

---

# 40. Complete authentication flow

তোমার system-এ flow হবে:

```text
                 ┌──────────────────┐
                 │      LOGIN       │
                 └────────┬─────────┘
                          │
                          ▼
                 Validate email/password
                          │
                          ▼
                  Check password hash
                          │
                          ▼
                   Create DB Session
                          │
              ┌───────────┴───────────┐
              ▼                       ▼
       Access JWT               Refresh Token
       15 minutes                  7 days
              │                       │
              │                       ▼
              │                SHA256 Hash
              │                       │
              │                       ▼
              │                  PostgreSQL
              │
              ▼
       Authorization Header
              │
              ▼
        Protected APIs
```

---

# 41. Access token expiration

আমি recommend করব:

```text
Access Token
    ↓
15 minutes
```

কেন?

যদি access token leak হয়, attacker maximum short period পর্যন্ত ব্যবহার করতে পারবে।

---

# 42. Refresh token

```text
Refresh Token
    ↓
7 days
```

তবে তোমার business requirement অনুযায়ী:

```text
7 days
14 days
30 days
```

করতে পারো।

আমি default হিসেবে:

```text
Access  = 15m
Refresh = 7d
```

নেব।

---

# 43. Database migration

তোমার `database.go`-তে:

```go
err := db.AutoMigrate(
	&users.User{},
	&auth.AuthSession{},
	&coachings.Coaching{},
	&students.Student{},
	&teachers.Teacher{},
	&subjects.Subject{},
)
```

Auth session-এর জন্য table হবে:

```text
auth_sessions
```

Conceptually:

```text
auth_sessions
--------------------------------
id
user_id
coaching_id
refresh_token_hash
user_agent
ip_address
expires_at
revoked_at
created_at
updated_at
deleted_at
```

---

# 44. Multi-tenant Coaching SaaS-এর জন্য সবচেয়ে গুরুত্বপূর্ণ বিষয়

তোমার system:

```text
Coaching A
    ├── Users
    ├── Students
    ├── Teachers
    └── Subjects

Coaching B
    ├── Users
    ├── Students
    ├── Teachers
    └── Subjects
```

তাই JWT-তে:

```json
{
  "uid": 15,
  "cid": 2,
  "role": "TEACHER"
}
```

থাকতে পারে।

এখানে:

```text
uid = User ID
cid = Coaching ID
role = User role
```

---

# 45. Tenant isolation

ধরো:

```text
Teacher A
CoachingID = 1
```

সে যেন:

```http
GET /students
```

দিয়ে Coaching B-এর students না দেখতে পারে।

Repository query হবে:

```go
func (r *repository) GetStudents(
	ctx context.Context,
	coachingID uint,
) ([]*Student, error) {

	var students []*Student

	err := r.db.WithContext(ctx).
		Where(
			"coaching_id = ?",
			coachingID,
		).
		Find(&students).
		Error

	return students, err
}
```

অর্থাৎ authentication থেকে পাওয়া:

```text
coachingID
```

দিয়ে every tenant query filter করবে।

এটা তোমার SaaS architecture-এর জন্য **critical**।

---

# 46. Don't trust client-provided CoachingID

খুব গুরুত্বপূর্ণ।

❌ এভাবে করা যাবে না:

```json
{
  "name": "Rahim",
  "coachingId": 999
}
```

তারপর:

```go
student.CoachingID = req.CoachingID
```

কারণ user অন্য coaching-এর ID পাঠাতে পারে।

বরং:

```go
coachingID, ok := auth.GetCoachingID(c)

if !ok || coachingID == nil {
	return apperror.Forbidden(
		"coaching context is required",
	)
}

student.CoachingID = *coachingID
```

অর্থাৎ:

```text
JWT
 ↓
Authenticated User
 ↓
CoachingID
 ↓
Server decides tenant
```

Client নয়।

---

# 47. Auth route design

আমি তোমার জন্য এই API design রাখব:

```text
POST   /api/v1/auth/login
POST   /api/v1/auth/refresh
POST   /api/v1/auth/logout
POST   /api/v1/auth/logout-all

GET    /api/v1/auth/me

GET    /api/v1/auth/sessions
DELETE /api/v1/auth/sessions/:sessionId
```

Future:

```text
POST /api/v1/auth/change-password
POST /api/v1/auth/forgot-password
POST /api/v1/auth/reset-password
POST /api/v1/auth/verify-email
```

---

# 48. `/me`

```go
func (h *handler) Me(c echo.Context) error {

	userID, ok := GetUserID(c)

	if !ok {
		return apperror.Unauthorized(
			"authentication required",
		)
	}

	user, err := h.userService.GetByID(
		c.Request().Context(),
		userID,
	)

	if err != nil {
		return apperror.NotFound(
			"user not found",
		)
	}

	return c.JSON(
		http.StatusOK,
		user,
	)
}
```

Route:

```go
protected.GET(
	"/me",
	handler.Me,
)
```

---

# 49. Frontend request flow

তোমার Next.js frontend হলে:

### Login

```text
POST /auth/login
        ↓
Access Token response
        ↓
Refresh token HttpOnly cookie
```

Access token memory/state-এ রাখতে পারো।

Protected API:

```http
Authorization: Bearer <access-token>
```

যখন:

```text
Access Token expired
       ↓
API returns 401
       ↓
POST /auth/refresh
       ↓
New Access Token
       ↓
Retry original request
```

---

# 50. কেন access token localStorage-এ না রাখাই ভালো

আমি recommend করব:

```text
❌ localStorage
❌ sessionStorage
```

এর বদলে:

```text
Access Token
    ↓
memory/state

Refresh Token
    ↓
HttpOnly Cookie
```

কারণ HttpOnly cookie JavaScript থেকে directly read করা যায় না।

তবে cookie-based authentication করলে CSRF protections, SameSite policy এবং CORS configuration ঠিকভাবে করতে হবে।

---

# 51. CORS

যদি frontend:

```text
http://localhost:3000
```

এবং backend:

```text
http://localhost:5000
```

হয়:

```go
e.Use(middleware.CORSWithConfig(
	middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
		},

		AllowMethods: []string{
			http.MethodGET,
			http.MethodPOST,
			http.MethodPUT,
			http.MethodPATCH,
			http.MethodDELETE,
			http.MethodOPTIONS,
		},

		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},

		AllowCredentials: true,
	},
))
```

Production-এ:

```text
AllowOrigins = exact frontend domain
```

দেবে।

এভাবে:

```go
AllowOrigins: []string{"*"}
```

এবং:

```go
AllowCredentials: true
```

একসাথে ব্যবহার করা উচিত নয়।

---

# 52. Security checklist

Production deployment-এর আগে:

```text
✓ HTTPS
✓ Short-lived access JWT
✓ Refresh token rotation
✓ Refresh token hashed in DB
✓ HttpOnly refresh cookie
✓ Secure cookie in production
✓ SameSite configured
✓ Password hashing
✓ No password in JWT
✓ No sensitive data in JWT
✓ Session revocation
✓ Logout
✓ Logout all devices
✓ Account status checking
✓ Role authorization
✓ Coaching/tenant isolation
✓ CORS restricted
✓ Rate limit login
✓ Generic invalid credentials message
✓ Token expiry validation
✓ Signing algorithm validation
```

---

# 53. One important improvement: refresh-token reuse detection

আরও production-level করতে চাইলে session table-এ রাখবে:

```go
ReplacedBySessionID *uint
```

অথবা token-family design করবে।

Flow:

```text
Refresh Token A
       ↓
Refresh
       ↓
Token B
       ↓
Token A becomes invalid
```

যদি attacker পুরনো Token A দিয়ে আবার refresh করে:

```text
Token A reused
       ↓
Security event
       ↓
Revoke entire token family
       ↓
Force login
```

এটা high-security production system-এর জন্য আরও ভালো।

---

# 54. Final architecture

তোমার project-এ authentication layer আমি এভাবে রাখব:

```text
                    CLIENT
                       │
                       ▼
              ┌────────────────┐
              │   Login API    │
              └───────┬────────┘
                      │
          email/password validation
                      │
                      ▼
                Password Hash
                      │
                      ▼
                  PostgreSQL
                      │
          ┌───────────┴───────────┐
          │                       │
          ▼                       ▼
    Access JWT             Auth Session
      15 min                    │
                                ▼
                        Refresh Token Hash
                                │
                                ▼
                           PostgreSQL


Protected Request
       │
       ▼
Authorization: Bearer JWT
       │
       ▼
JWT Middleware
       │
       ├── UserID
       ├── CoachingID
       ├── Role
       └── SessionID
       │
       ▼
Authentication
       │
       ▼
Authorization
       │
       ▼
Tenant Isolation
       │
       ▼
Repository
       │
       ▼
PostgreSQL
```

## আমার recommended final module boundary

```text
internal/domain/users
        │
        ├── entity.go
        ├── repository.go
        ├── service.go
        ├── handler.go
        └── dto/

internal/domain/auth
        │
        ├── entity.go              ← AuthSession
        ├── token.go               ← JWT + refresh token
        ├── password.go            ← password hashing
        ├── repository.go          ← session DB operations
        ├── service.go             ← login/refresh/logout
        ├── middleware.go          ← authentication
        ├── handler.go             ← HTTP endpoints
        ├── mapper.go
        ├── register.go
        └── dto/
             ├── request.go
             └── response.go
```

**সবচেয়ে গুরুত্বপূর্ণ architectural rule:** `auth` module authentication/session/token-এর দায়িত্ব নেবে, আর `users` module user identity/profile-এর দায়িত্ব নেবে। অন্যান্য module (`students`, `teachers`, `subjects`, `payments` ইত্যাদি) কখনো নিজে JWT parse করবে না—তারা middleware থেকে authenticated `UserID`, `CoachingID` এবং `Role` নেবে। Սա তোমার Coaching Management SaaS-কে পরে বড় করার জন্য অনেক cleaner এবং safer architecture হবে।
