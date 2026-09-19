package notices

import "coaching_backend/internal/domain/notices/dto"

func ToNoticeResponse(notice *Notice) *dto.NoticeResponse{
	return &dto.NoticeResponse{}
}

func ToNoticeResponses(notices []*Notice) []*dto.NoticeResponse{
	response := make([]*dto.NoticeResponse, 0, len(notices))

	for i := range notices {
		return append(response, ToNoticeResponse(notices[i]))
	}

	return response
}