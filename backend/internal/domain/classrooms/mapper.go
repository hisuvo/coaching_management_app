package classrooms

import "coaching_backend/internal/domain/classrooms/dto"

func ToClassRoomResponse(classroom *ClassRoom) *dto.ClassRoomResponse{
	return &dto.ClassRoomResponse{}
}

func ToClassRoomResponses(classrooms []*ClassRoom) []*dto.ClassRoomResponse{
	response := make([]*dto.ClassRoomResponse, 0, len(classrooms))

	for i := range classrooms {
		return append(response, ToClassRoomResponse(classrooms[i]))
	}

	return response
}