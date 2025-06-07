package usecase

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/sandronister/mbalab2/weather_service/internal/dto"
	"github.com/sandronister/mbalab2/weather_service/internal/infra/http/types"
)

type WeatherUseCase struct {
	path string
	http types.IHttpService
}

func NewWeatherUseCase(path string, http types.IHttpService) *WeatherUseCase {
	return &WeatherUseCase{
		path: path,
		http: http,
	}
}

func (w *WeatherUseCase) requestWeatherData(url string) (*dto.WeatherResultDTO, *types.HttpError) {
	var weatherResult dto.WeatherResultDTO

	resp, err := w.http.Do(url)
	if err != nil {
		return nil, types.NewHttpError(422, "invalid city name")
	}

	if len(resp) == 0 {
		return nil, types.NewHttpError(404, "city not found")
	}

	err = json.Unmarshal(resp, &weatherResult)

	if err != nil {
		return nil, types.NewHttpError(422, "invalid city name")
	}

	if len(weatherResult.NearestArea) == 0 {
		return nil, types.NewHttpError(404, "city not found")
	}

	if len(weatherResult.NearestArea[0].AreaName) == 0 {
		return nil, types.NewHttpError(404, "city not found")
	}

	if weatherResult.NearestArea[0].AreaName[0].Value == "" && weatherResult.CurrentCondition == nil {
		return nil, types.NewHttpError(404, "city not found")
	}

	return &weatherResult, nil

}

func (w *WeatherUseCase) Get(city string) (*dto.ResultDTO, *types.HttpError) {
	if w.path == "" {
		return nil, types.NewHttpError(500, "invalid path")
	}

	url := fmt.Sprintf(w.path, city)

	result, err := w.requestWeatherData(url)
	if err != nil {
		return nil, err
	}
	feelsLikeC, parseErr := strconv.ParseFloat(result.CurrentCondition[0].FeelsLikeC, 64)
	if parseErr != nil {
		return nil, types.NewHttpError(500, "invalid temperature value")
	}

	feelsLikeF, parseErr := strconv.ParseFloat(result.CurrentCondition[0].FeelsLikeF, 64)
	if parseErr != nil {
		return nil, types.NewHttpError(500, "invalid temperature value")
	}
	Kelvin := feelsLikeC + 273.15

	resultDTO := &dto.ResultDTO{
		TempCelsius:    fmt.Sprintf("%.2f", feelsLikeC),
		TempFahrenheit: fmt.Sprintf("%.2f", feelsLikeF),
		TempKelvin:     fmt.Sprintf("%.2f", Kelvin),
	}

	return resultDTO, nil
}
