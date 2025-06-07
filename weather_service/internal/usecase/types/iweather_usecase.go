package types

import (
	"github.com/sandronister/mbalab2/weather_service/internal/dto"
	"github.com/sandronister/mbalab2/weather_service/internal/infra/http/types"
)

type IWeatherUseCase interface {
	Get(localidade string) (*dto.ResultDTO, *types.HttpError)
}
