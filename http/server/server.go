package server

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	server *echo.Echo
	config Config
}

func NewServer(config Config) *Server {
	return &Server{
		server: echo.New(),
		config: config,
	}

}

func (s *Server) Start() error {
	return s.server.Start(":8080")
}
