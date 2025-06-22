package usecase

import (
	"context"

	"github.com/sandronister/mba-lab2/cep-service/internal/dto"
)

type Finder interface {
	Execute(ctx context.Context, cep string) (*dto.LocaleOutput, error)
}
