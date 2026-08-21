package apperror

import "net/http"

type AppError struct {
	Status  int
	Code	string
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func New(status int, code, message string) *AppError {
	return &AppError{
		Status:status,
		Code: code,
		Message: message,
	}
}

func BadRequest(message string) *AppError {
	return New(http.StatusBadRequest, "BAD_REQUEST", message)
}

func Unauthorized(message string) *AppError {
	return New(http.StatusUnauthorized, "UNAUTHORIZED", message)
}

func Forbidden(message string) *AppError {
	return New(http.StatusForbidden, "FORBIDDEN", message)
}

func NotFound(message string) *AppError {
	return New(http.StatusNotFound, "NOT_FOUND", message)
}

func Conflict(message string) *AppError {
	return New(http.StatusConflict, "CONFLICT", message)
}

func UnprocessableEntity(message string) *AppError {
	return New(http.StatusUnprocessableEntity, "VALIDATION_ERROR", message)
}

func InternalServerError(message string) *AppError {
	return New(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", message)
}

func ServiceUnavailable(message string) *AppError {
	return New(http.StatusServiceUnavailable, "SERVICE_UNAVAILABLE", message)
}