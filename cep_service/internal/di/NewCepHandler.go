package di

import (
	"context"

	"github.com/sandronister/mbalab2/cep_service/config"
	clienthttp "github.com/sandronister/mbalab2/cep_service/internal/infra/client_http"
	"github.com/sandronister/mbalab2/cep_service/internal/infra/web/handler"
	"github.com/sandronister/mbalab2/cep_service/internal/usecase"
)

func NewCepHandler(env *config.EnviromentVar) *handler.CepHandler {
	ctx := context.Background()
	requestService := clienthttp.NewRequestClient(ctx)
	httpService := clienthttp.NewHttpService(requestService)
	usecase := usecase.NewUCep(httpService, env.ServiceURL)
	return handler.NewCepHandler(usecase)
}
