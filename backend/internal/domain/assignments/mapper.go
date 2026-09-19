package assignments

import "coaching_backend/internal/domain/assignments/dto"

func ToAssignmentResponse(assignment *Assignment) *dto.AssignmentResponse {
	return &dto.AssignmentResponse{
		ID: assignment.ID,
		Title: assignment.Title,
		Description: assignment.Description,
		SubjectID: assignment.SubjectID,
		CoachingID: assignment.CoachingID,
		DueDate: assignment.DueDate,
		Status: string(assignment.Status),
		CreatedBy: assignment.CreatedBy,
		CreatedAt: assignment.CreatedAt,
		UpdatedAt: assignment.UpdatedAt,
	}
}

func ToAssignmentResponses(assignments []*Assignment) []*dto.AssignmentResponse {
	responses := make([]*dto.AssignmentResponse, 0, len(assignments))
	
	for i := range assignments {
		responses = append(responses, ToAssignmentResponse(assignments[i]))
	}
	return responses
}