package dto

type CreateBranchRequest struct {
	CoachingID uint   `json:"coaching_id" validate:"required"`
	Name       string `json:"name" validate:"required,min=2,max=100"`
	Code       string `json:"code" validate:"required,min=2,max=50"`
	Phone      string `json:"phone" validate:"omitempty,max=20"`
	Email      string `json:"email" validate:"omitempty,email,max=150"`
	Address    string `json:"address" validate:"omitempty,max=500"`
	City       string `json:"city" validate:"omitempty,max=100"`
	Country    string `json:"country" validate:"omitempty,max=100"`
	Status     string `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`
}

type UpdateBranchRequest struct {
	Name    *string `json:"name" validate:"omitempty,min=2,max=100"`
	Code    *string `json:"code" validate:"omitempty,min=2,max=50"`
	Phone   *string `json:"phone" validate:"omitempty,max=20"`
	Email   *string `json:"email" validate:"omitempty,email,max=150"`
	Address *string `json:"address" validate:"omitempty,max=500"`
	City    *string `json:"city" validate:"omitempty,max=100"`
	Country *string `json:"country" validate:"omitempty,max=100"`
	Status  *string `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`
}