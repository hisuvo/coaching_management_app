package dto

type TeacherResponse struct {
	ID          uint   `json:"id"`
	CoachingID  uint   `json:"coaching_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Subject     string `json:"subject"`
	Designation string `json:"designation"`
	Address     string `json:"address"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}