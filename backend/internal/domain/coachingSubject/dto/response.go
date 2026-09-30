package dto

type CoachingSubjectResponse struct {
	ID         uint   `json:"id"`
	CoachingID uint   `json:"coaching_id"`
	SubjectID  uint   `json:"subject_id"`
	Status     string `json:"status"`
}