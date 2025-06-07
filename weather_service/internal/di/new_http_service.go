package di

import (
	"context"

	"github.com/sandronister/mbalab2/weather_service/internal/infra/http"
	"github.com/sandronister/mbalab2/weather_service/internal/infra/http/request"
	"github.com/sandronister/mbalab2/weather_service/internal/infra/http/types"
)

func newHttpService() types.IHttpService {
	requestService := request.New(context.Background())
	return http.NewService(requestService)
}
