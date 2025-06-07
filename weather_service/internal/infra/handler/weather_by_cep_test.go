// go
package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/sandronister/mbalab2/weather_service/internal/dto"
	"github.com/sandronister/mbalab2/weather_service/internal/infra/http/types"

	"github.com/stretchr/testify/assert"
)

type mockCepUsecase struct {
	GetFunc func(string) (*dto.CepResultDTO, *types.HttpError)
}

func (m *mockCepUsecase) Get(cep string) (*dto.CepResultDTO, *types.HttpError) {
	return m.GetFunc(cep)
}

type mockWeatherUseCase struct {
	GetFunc func(string) (*dto.ResultDTO, *types.HttpError)
}

func (m *mockWeatherUseCase) Get(localidade string) (*dto.ResultDTO, *types.HttpError) {
	return m.GetFunc(localidade)
}

func TestWeatherByCep_Get_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/weather/12345678", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("cep")
	c.SetParamValues("12345678")

	mockCep := &mockCepUsecase{
		GetFunc: func(cep string) (*dto.CepResultDTO, *types.HttpError) {
			return &dto.CepResultDTO{Localidade: "Sao Paulo"}, nil
		},
	}
	mockWeather := &mockWeatherUseCase{
		GetFunc: func(localidade string) (*dto.ResultDTO, *types.HttpError) {
			return &dto.ResultDTO{
				TempCelsius:    "25.0",
				TempFahrenheit: "77",
				TempKelvin:     "298.15",
			}, nil
		},
	}

	handler := NewWeatherByCep(mockCep, mockWeather)
	err := handler.Get(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.Contains(t, resp, "data")
	data := resp["data"].(map[string]interface{})
	assert.Equal(t, "25.0", data["temp_c"])
}

func TestWeatherByCep_Get_CepError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/weather/00000000", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("cep")
	c.SetParamValues("00000000")

	mockCep := &mockCepUsecase{
		GetFunc: func(cep string) (*dto.CepResultDTO, *types.HttpError) {
			return nil, &types.HttpError{StatusCode: 404, Message: "CEP not found"}
		},
	}
	mockWeather := &mockWeatherUseCase{}

	handler := NewWeatherByCep(mockCep, mockWeather)
	err := handler.Get(c)
	assert.NoError(t, err)
	assert.Equal(t, 404, rec.Code)

	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NotEmpty(t, resp["error"])
}

func TestWeatherByCep_Get_WeatherError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/weather/12345678", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("cep")
	c.SetParamValues("12345678")

	mockCep := &mockCepUsecase{
		GetFunc: func(cep string) (*dto.CepResultDTO, *types.HttpError) {
			return &dto.CepResultDTO{Localidade: "Sao Paulo"}, nil
		},
	}
	mockWeather := &mockWeatherUseCase{
		GetFunc: func(localidade string) (*dto.ResultDTO, *types.HttpError) {
			return nil, &types.HttpError{StatusCode: 500, Message: "Weather service error"}
		},
	}

	handler := NewWeatherByCep(mockCep, mockWeather)
	err := handler.Get(c)
	assert.NoError(t, err)
	assert.Equal(t, 500, rec.Code)

	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.NotEmpty(t, resp["error"])
}
