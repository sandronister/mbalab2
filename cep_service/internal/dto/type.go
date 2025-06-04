package dto

import "fmt"

type Cep struct {
	Cep string `json:"cep"`
}

type Result struct {
	TempCelsius    string `json:"temp_c"`
	TempFahrenheit string `json:"temp_f"`
	TempKelvin     string `json:"temp_k"`
}

type HttpError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.StatusCode, e.Message)
}

func NewHttpError(statusCode int, message string) *HttpError {
	return &HttpError{
		StatusCode: statusCode,
		Message:    message,
	}
}
