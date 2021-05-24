package utils

type ApplicationError struct {
	Message    string `json:"massage"`
	StatusCode int    `json:"satus"`
	Code       string `json:"code"`
}
