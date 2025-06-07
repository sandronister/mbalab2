package types

import (
	"github.com/sandronister/mbalab2/weather_service/internal/dto"
	"github.com/sandronister/mbalab2/weather_service/internal/infra/http/types"
)

type ICep interface {
	Get(cep string) (*dto.CepResultDTO, *types.HttpError)
}
