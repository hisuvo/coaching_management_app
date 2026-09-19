package submissions

import "coaching_backend/internal/domain/submissions/dto"

func ToSubmissionResponse(submission *Submission) *dto.SubmissionResponse{
	return &dto.SubmissionResponse{
		ID: submission.ID,
		AssignmentID: submission.AssignmentID,
		StudentID: submission.StudentID,
		TextAnswer: submission.TextAnswer,
		FileURL: &submission.FileURL,
		SubmittedAt: submission.SubmittedAt,
		Status: submission.Status,
		ObtainedMarks: submission.ObtainedMarks,
		Feedback: submission.Feedback,
		ReviewedBy: submission.ReviewedBy,
		ReviewedAt: submission.ReviewedAt,
		CreatedAt: submission.CreatedAt,
		UpdatedAt: submission.UpdatedAt,
	}
}

func ToSubmissionResponses(submissions []*Submission) []*dto.SubmissionResponse {
	responses := make([]*dto.SubmissionResponse, 0, len(submissions))

	for i := range submissions {
		response := ToSubmissionResponse(submissions[i])

		if response != nil {
			responses = append(responses, response)
		}
		// return append(responses, ToSubmissionResponse(submissions[i]))
	}

	return responses
}