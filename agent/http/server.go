package server

import (
	"errors"
	"github.com/chessnok/GoCalculator/agent/http/handler"
	"github.com/labstack/echo"
	"strconv"
)

var (
	// ErrInvalidPort is returned when the port is invalid
	ErrInvalidPort = errors.New("invalid port")
)

// Server is a struct that contains an echo server and a configuration
type Server struct {
	server *echo.Echo
	config *Config
}

// NewServer creates a new server with the given configuration and returns a pointer to it
func NewServer(config *Config) *Server {
	return &Server{
		server: echo.New(),
		config: config,
	}
}
func validatePort(port int) error {
	if port < 0 || port > 65535 {
		return ErrInvalidPort
	}
	return nil
}

// Start starts the server and listens for incoming requests
func (s *Server) Start() error {
	s.server.POST("/updateConfig", handler.UpdateConfigHandler(s.config.Calculator))
	s.server.GET("/currentOperation", handler.LastOperationHandler(s.config.Calculator))
	if err := validatePort(s.config.Port); err != nil {
		return err
	}
	return s.server.Start(":" + strconv.Itoa(s.config.Port))
}
