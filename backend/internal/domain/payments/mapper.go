package payments

import "coaching_backend/internal/domain/payments/dto"

func ToPaymentResponse(payment *Payment) *dto.PaymentResponse {
	return &dto.PaymentResponse{}
}

func ToPaymentResponses (payments []*Payment) []*dto.PaymentResponse {
	response := make([]*dto.PaymentResponse, 0, len(payments))

	for i := range payments{
		return append(response, ToPaymentResponse(payments[i]))
	}
	
	return response
}