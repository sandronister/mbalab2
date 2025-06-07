package clienthttp

import (
	"context"
	"net/http"

	"githuc.com/sandronister/mbalab2/cep_service/internal/infra/client_http/types"
)

type model struct {
	ctx context.Context
}

func NewRequestClient(ctx context.Context) types.IRequestService {
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
