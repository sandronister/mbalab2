package server

import (
	"fmt"

	"github.com/labstack/echo"
	"githuc.com/sandronister/mbalab2/cep_service/internal/infra/web/handler"
)

type model struct {
	WebPort string
	router  *echo.Echo
}

func NewServer(webPort string) *model {
	router := echo.New()
	server := &model{
		WebPort: webPort,
		router:  router,
	}

	return server
}

func (s *model) AddRouter(handler *handler.CepHandler) {
	s.router.POST("/", handler.GetWeather)
}

func (s *model) Start() error {
	return s.router.Start(fmt.Sprintf(":%s", s.WebPort))
}
