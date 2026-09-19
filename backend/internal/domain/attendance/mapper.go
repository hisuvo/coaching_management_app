package attendance

import "coaching_backend/internal/domain/attendance/dto"

func ToAttendanceResponse(attendance *Attendance) *dto.AttendanceResponse {
	return &dto.AttendanceResponse{}
}

func ToAttendanceResponses(attendances []*Attendance) []*dto.AttendanceResponse{
	response := make([]*dto.AttendanceResponse, 0, len(attendances))

	for i := range attendances {
		return append(response, ToAttendanceResponse(attendances[i]))
	}

	return response
}