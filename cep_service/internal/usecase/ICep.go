package usecase

import (
	"context"

	"github.com/sandronister/mbalab2/cep_service/internal/dto"
)

type ICep interface {
	Do(ctx context.Context, cep string) (*dto.Result, *dto.HttpError)
}
