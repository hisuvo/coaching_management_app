package dto

import (
	"coaching_backend/internal/domain/subjects/dto"
)

type CreateCoachingSubjectResponse struct {
	ID         uint                    `json:"id"`
	Status     string                  `json:"status"`
	Subject    *dto.SubjectResponse `json:"subject,omitempty"`
}

type CoachingSubjectResponse struct {
	ID         uint                    `json:"id,omitempty"`
	CoachingID uint                    `json:"coaching_id,omitempty"`
	SubjectID  uint                    `json:"subject_id,omitempty"`
	Status     string                  `json:"status"`
	Subject    *dto.SubjectResponse `json:"subject,omitempty"`
}