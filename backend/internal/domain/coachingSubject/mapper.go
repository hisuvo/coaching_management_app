package coachingsubject

import (
	"coaching_backend/internal/domain/coachingSubject/dto"
	"coaching_backend/internal/domain/subjects"
)

func ToCreateCoachingSubjectResponse(coachingSubject *CoachingSubject,) *dto.CreateCoachingSubjectResponse {

	if coachingSubject == nil {
		return nil
	}

	resporse := &dto.CreateCoachingSubjectResponse{
		ID:         coachingSubject.ID,
		Status:     coachingSubject.Status,

		Subject: subjects.ToSubjectResponse(&coachingSubject.Subject),
	}

	return resporse
}

func ToCoachingSubjectResponse(coachingSubject *CoachingSubject) *dto.CoachingSubjectResponse{
	return &dto.CoachingSubjectResponse{
		CoachingID: coachingSubject.CoachingID,
		SubjectID: coachingSubject.SubjectID,
		Status: coachingSubject.Status,
	}
}

func ToCoachingSubjectRes(coachingSubject *CoachingSubject) *dto.CoachingSubjectResponse{
		return &dto.CoachingSubjectResponse{
			Status: coachingSubject.Status,
			Subject: subjects.ToSubjectResponse(&coachingSubject.Subject),
		}
}

func ToCoachingSubjectResponses(CoachingSubjects []*CoachingSubject) []*dto.CoachingSubjectResponse{
	responses := make([]*dto.CoachingSubjectResponse, 0, len(CoachingSubjects))

	for i := range CoachingSubjects{
		responses = append(responses, ToCoachingSubjectRes(CoachingSubjects[i]))
	}

	return responses
}