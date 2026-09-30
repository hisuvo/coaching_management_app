package coachingsubject

import "coaching_backend/internal/domain/coachingSubject/dto"

func ToCoachingSubjectResponse(coachingSubject *CoachingSubject) *dto.CoachingSubjectResponse{
	return &dto.CoachingSubjectResponse{
		CoachingID: coachingSubject.CoachingID,
		SubjectID: coachingSubject.SubjectID,
		Status: coachingSubject.Status,
	}
}