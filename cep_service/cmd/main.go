package main

import (
	"github.com/sandronister/mbalab2/cep_service/config"
	"github.com/sandronister/mbalab2/cep_service/internal/di"
	"github.com/sandronister/mbalab2/cep_service/internal/infra/web/server"
)

func main() {
	config, err := config.LoadEnviromentVars(".")
	if err != nil {
		panic(err)
	}

	handler := di.NewCepHandler(config)

	server := server.NewServer(config.WebPort)
	server.AddRouter(handler)
	if err := server.Start(); err != nil {
		panic(err)
	}

	err = server.Start()

	if err != nil {
		panic(err)
	}

}
