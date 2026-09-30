package dto

type CreateCoachingSubjectRequest struct {
	CoachingID uint   `json:"coaching_id" validate:"required"`
	SubjectID  uint   `json:"subject_id" validate:"required"`
	Status     string `json:"status" validate:"omitempty,oneof=active inactive"`
}

type UpdateCoachingSubjectRequest struct {
	Status *string `json:"status,omitempty" validate:"omitempty,oneof=active inactive"`
}