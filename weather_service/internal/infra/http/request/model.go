package request

import (
	"context"
	"net/http"

	"github.com/sandronister/mbalab2/weather_service/internal/infra/http/types"
)

type model struct {
	ctx context.Context
}

func New(ctx context.Context) types.IRequestService {
	return &model{
		ctx: ctx,
	}
}

func (m *model) Get(url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(m.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
