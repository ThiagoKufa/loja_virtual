package utils

type ErrorResponse struct {
	Field string `json:"field"`
	Error string `json:"error"`
	Code  int    `json:"code"`
}
