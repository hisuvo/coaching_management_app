package subjects

import (
	"coaching_backend/internal/domain/subjects/dto"
	"strconv"
)

func ToSubjectResponse(subject *Subject) *dto.SubjectResponse{
	return &dto.SubjectResponse{
		Id: strconv.Itoa(int(subject.ID)),
		Name: subject.Name,
		Code: subject.Code,
		Description: subject.Description,
		Status: string(subject.Status),
	}
}

func ToSubjectResponses(subjects []*Subject)[]*dto.SubjectResponse{
	responses := make([]*dto.SubjectResponse, 0, len(subjects))

	for i := range subjects {
		response := ToSubjectResponse(subjects[i])
		responses = append(responses, response)
	}

	return responses
}