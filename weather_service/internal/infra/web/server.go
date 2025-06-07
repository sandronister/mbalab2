package web

import (
	"fmt"

	"github.com/labstack/echo"
)

type Server struct {
	Port   string
	router *echo.Echo
}

func NewServer(port string) *Server {
	return &Server{
		Port:   port,
		router: echo.New(),
	}
}

func (s *Server) Start() error {
	return s.router.Start(fmt.Sprintf(":%s", s.Port))
}
