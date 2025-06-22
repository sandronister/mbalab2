package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sandronister/mba-lab2/weather-service/internal/dto"
	"github.com/sandronister/mba-lab2/weather-service/internal/usecase"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

type WeatherHandler struct {
	weatherFinder usecase.Finder
	localeFinder  usecase.Finder
}

func NewWeatherHandler(weatherFinder usecase.Finder, localeFinder usecase.Finder) *WeatherHandler {
	return &WeatherHandler{
		weatherFinder: weatherFinder,
		localeFinder:  localeFinder,
	}
}

func (h *WeatherHandler) Handle(w http.ResponseWriter, r *http.Request) {
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	tracer := otel.Tracer("weather-service")
	_, span := tracer.Start(ctx, "weather-handler")
	defer span.End()

	w.Header().Set("Content-Type", "application/json")

	cep := r.PathValue("cep")
	if len(cep) != 8 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(w).Encode(&dto.ErrorOutput{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "invalid zipcode",
		})
		return
	}

	localeOutputRaw, err := h.localeFinder.Execute(ctx, cep)
	if err != nil {
		log.Println("Error while fetching locale data:", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&dto.ErrorOutput{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	localeOutput := localeOutputRaw.(*dto.LocaleOutput)
	if localeOutput.Localidade == "" {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(&dto.ErrorOutput{
			StatusCode: http.StatusNotFound,
			Message:    "can not find zipcode",
		})
		return
	}

	weatherOutputRaw, err := h.weatherFinder.Execute(ctx, localeOutput.Localidade)
	if err != nil {
		log.Println("Error while fetching weather data:", err)
		if err.Error() == "API key is invalid" || err.Error() == "API key is not provided" {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(&dto.ErrorOutput{
				StatusCode: http.StatusUnauthorized,
				Message:    err.Error(),
			})
			return
		}

		if err.Error() == "can not find zipcode" {
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(&dto.ErrorOutput{
				StatusCode: http.StatusNotFound,
				Message:    err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(&dto.ErrorOutput{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	weatherOutput := weatherOutputRaw.(*dto.WeatherOutput)

	w.WriteHeader(http.StatusOK)
	result := dto.ResultOutput{
		City:  localeOutput.Localidade,
		TempC: weatherOutput.Current.TempC,
		TempF: weatherOutput.Current.TempF,
		TempK: weatherOutput.Current.TempC + 273.15,
	}
	log.Println("Weather data fetched successfully:", result)

	_ = json.NewEncoder(w).Encode(result)
}
