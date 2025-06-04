package server

import (
	"fmt"

	"github.com/labstack/echo"
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

func (s *model) Start() error {
	return s.router.Start(fmt.Sprintf(":%s", s.WebPort))
}
