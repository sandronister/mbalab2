package dto

type ResultDTO struct {
	TempCelsius    string `json:"temp_c"`
	TempFahrenheit string `json:"temp_f"`
	TempKelvin     string `json:"temp_k"`
}
