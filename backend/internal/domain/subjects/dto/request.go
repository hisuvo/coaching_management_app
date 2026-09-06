package dto

type CreateSubjectRequest struct {
	Name        string  `json:"name" validate:"required,min=2,max=100"`
	Code        string  `json:"code" validate:"required,min=2,max=50"`
	Description *string `json:"description" validate:"omitempty,max=500"`
	Status      string  `json:"status" validate:"omitempty,oneof=active inactive"`
}

type UpdateSubjectRequest struct {
	Name        *string `json:"name" validate:"omitempty,min=2,max=100"`
	Code        *string `json:"code" validate:"omitempty,min=2,max=50"`
	Description *string `json:"description" validate:"omitempty,max=500"`
	Status      *string `json:"status" validate:"omitempty,oneof=active inactive"`
}

type SubjectQuery struct {
	Class int `query:"class"`
	Limit int `query:"limit"`
	Page  int `query:"page"`
}