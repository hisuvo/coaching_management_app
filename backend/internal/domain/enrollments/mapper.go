package enrollments

import "coaching_backend/internal/domain/enrollments/dto"

func ToEnrollmentResponse(enrollment *Enrollment) *dto.EnrollmentResponse {
	return &dto.EnrollmentResponse{}
}

func ToEnrollmentResponses(enrollments []*Enrollment) []*dto.EnrollmentResponse {
	response := make([]*dto.EnrollmentResponse, 0, len(enrollments))

	for i := range enrollments {
		response = append(response, ToEnrollmentResponse(enrollments[i]))
	}
	
	return response
}