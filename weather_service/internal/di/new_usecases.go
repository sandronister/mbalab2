package di

import (
	"github.com/sandronister/mbalab2/weather_service/config"
	"github.com/sandronister/mbalab2/weather_service/internal/infra/http/types"
	"github.com/sandronister/mbalab2/weather_service/internal/usecase"
	usecaseTypes "github.com/sandronister/mbalab2/weather_service/internal/usecase/types"
)

func NewCepUseCase(env *config.EnviromentVar, httpService types.IHttpService) usecaseTypes.ICep {
	return usecase.NewWeatherByCep(env.CepServicePath, httpService)
}

func NewWeatherUseCase(env *config.EnviromentVar, httpService types.IHttpService) usecaseTypes.IWeatherUseCase {
	return usecase.NewWeatherUseCase(env.WeatherServicePath, httpService)
}
