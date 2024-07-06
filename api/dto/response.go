package dto

type SuccessResponse[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
