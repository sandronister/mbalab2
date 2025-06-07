package di

import (
	"github.com/sandronister/mbalab2/weather_service/config"
	"github.com/sandronister/mbalab2/weather_service/internal/infra/handler"
)

func NewWeatherHandler(env *config.EnviromentVar) *handler.WeatherByCep {
	httpService := newHttpService()
	return handler.NewWeatherByCep(NewCepUseCase(env, httpService), NewWeatherUseCase(env, httpService))
}
