package schedules

import "coaching_backend/internal/domain/schedules/dto"

func ToSchedulesResponse(schedule *Schedule) *dto.ScheduleResponse {
	return &dto.ScheduleResponse{}
}


func ToSchedulesResponses(schedules []*Schedule) []*dto.ScheduleResponse {
	response := make([]*dto.ScheduleResponse, 0, len(schedules))

	for i := range schedules {
		return append(response, ToSchedulesResponse(schedules[i]))
	}
	
	return response
}


