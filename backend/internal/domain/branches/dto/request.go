package dto

type CreateBranchRequest struct {
	CoachingID uint   `json:"coaching_id" validate:"required"`
	Name       string `json:"name" validate:"required,min=2,max=100"`
	Code       string `json:"code" validate:"required,min=2,max=50"`
	Phone      string `json:"phone" validate:"omitempty,max=20"`
	Email      string `json:"email" validate:"omitempty,email,max=150"`
	Address    string `json:"address" validate:"omitempty,max=500"`
	Status     string `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`

	AdminName     string `json:"admin_name" validate:"required"`
	AdminEmail    string `json:"admin_email" validate:"required,email,max=150"`
	AdminPassword string `json:"admin_password" validate:"required,min=8"`
}

type UpdateBranchRequest struct {
	Name    *string `json:"name" validate:"omitempty,min=2,max=100"`
	Code    *string `json:"code" validate:"omitempty,min=2,max=50"`
	Phone   *string `json:"phone" validate:"omitempty,max=20"`
	Email   *string `json:"email" validate:"omitempty,email,max=150"`
	Address *string `json:"address" validate:"omitempty,max=500"`
	Status  *string `json:"status" validate:"omitempty,oneof=ACTIVE INACTIVE"`

	AdminName     string `josn:"admin_name" validate:"omitempty"`
	AdminEmail    string `json:"admin_email" validate:"omitempty,email,max=150"`
	AdminPassword string `json:"admin_password" validate:"omitempty,min=8"`
}