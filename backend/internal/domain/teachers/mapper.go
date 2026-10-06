package teachers

import "coaching_backend/internal/domain/teachers/dto"

func ToTeacherResponse(teacher *Teacher) *dto.TeacherResponse {
	return &dto.TeacherResponse{
		UserID: teacher.UserID,
		BranchID: teacher.BranchID,
		EmployeeNo: teacher.EmployeeNo,
		Designation: teacher.Designation,
		Qualification: teacher.Qualification,
		JoiningDate: teacher.JoiningDate,
		Status: string(teacher.Status),
		CreatedAt: teacher.CreatedAt,
		UpdatedAt: teacher.UpdatedAt,
	}
}

func ToTeacherResponses(teachers []*Teacher) []*dto.TeacherResponse {
	responses := make([]*dto.TeacherResponse, 0, len(teachers))

	for _, teacher := range teachers{
		responses = append(responses,ToTeacherResponse(teacher))
	}

	return responses
}