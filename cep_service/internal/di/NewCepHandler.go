package di

import (
	"context"

	"githuc.com/sandronister/mbalab2/cep_service/config"
	clienthttp "githuc.com/sandronister/mbalab2/cep_service/internal/infra/client_http"
	"githuc.com/sandronister/mbalab2/cep_service/internal/infra/web/handler"
	"githuc.com/sandronister/mbalab2/cep_service/internal/usecase"
)

func NewCepHandler(env *config.EnviromentVar) *handler.CepHandler {
	ctx := context.Background()
	requestService := clienthttp.NewRequestClient(ctx)
	httpService := clienthttp.NewHttpService(requestService)
	usecase := usecase.NewUCep(httpService, env.ServiceURL)
	return handler.NewCepHandler(usecase)
}
