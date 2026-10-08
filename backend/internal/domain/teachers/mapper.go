package teachers

import (
	branchdto "coaching_backend/internal/domain/branches/dto"
	"coaching_backend/internal/domain/teachers/dto"
	usersdto "coaching_backend/internal/domain/users/dto"
)

func ToTeacherResponse(teacher *Teacher) *dto.TeacherResponse {

	if teacher == nil {
		return nil
	}

	response := &dto.TeacherResponse{
		UserID: teacher.UserID,
		CoachingID: teacher.CoachingID,
		BranchID: teacher.BranchID,
		EmployeeNo: teacher.EmployeeNo,
		Designation: teacher.Designation,
		JoiningDate: &teacher.JoiningDate,
	}

	if teacher.User.ID != 0 {
		response.User = &usersdto.UserResponse{
			Id: teacher.User.ID,
			Name: teacher.User.Name,
			Email: teacher.User.Email,
			Role: string(teacher.User.Role),
			Phone: teacher.User.Phone,
			Status: string(teacher.User.Status),
		}
	}

	if teacher.Branch.ID != 0 {
		response.Branch = &branchdto.BranchResponse{
			ID: teacher.Branch.ID,
			Name: teacher.Branch.Name,
			Code: teacher.Branch.Code,
			Phone: teacher.Branch.Phone,
			Email: teacher.Branch.Email,
			Status: string(teacher.Branch.Status),
		}
	}
	return response
}

func ToTeacherResponses(teachers []*Teacher) []*dto.TeacherResponse {
	responses := make([]*dto.TeacherResponse, 0, len(teachers))

	for _, teacher := range teachers{
		responses = append(responses,ToTeacherResponse(teacher))
	}

	return responses
}