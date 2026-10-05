package dto

type CreateCoachingSubjectRequest struct {
	CoachingID  uint    `json:"coaching_id" validate:"required"`
	Name        string  `json:"name" validate:"required,min=2,max=100"`
	Code        string  `json:"code" validate:"required,min=2,max=50"`
	Description *string `json:"description" validate:"omitempty,max=500"`
}

type UpdateCoachingSubjectRequest struct {
	Status *string `json:"status,omitempty" validate:"omitempty,oneof=active inactive"`
}