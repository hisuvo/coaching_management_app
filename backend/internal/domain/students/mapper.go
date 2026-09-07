package students

import (
	"coaching_backend/internal/domain/students/dto"
	"strconv"
)

func ToStudentResponse(student *Student) *dto.StudentResponse {
	return &dto.StudentResponse{
		Id: strconv.Itoa(int(student.ID)),
		Name: student.Name,
		Class: student.Class,
		Session: student.Session,
		Email: student.Email,
		Phone: student.Phone,
		Address: student.Address,
		SchoolName: student.SchoolName,
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
	}
}

func ToStudentResponses(students []*Student) []*dto.StudentResponse{

	responses := make([]*dto.StudentResponse, 0, len(students))

	for i := range students {
		responses = append(responses, ToStudentResponse(students[i]))
	}

	return responses
}