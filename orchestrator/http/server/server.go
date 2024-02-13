package server

import (
	"errors"
	"github.com/chessnok/GoCalculator/orchestrator/http/server/handler/agents"
	configHandler "github.com/chessnok/GoCalculator/orchestrator/http/server/handler/config"
	"github.com/chessnok/GoCalculator/orchestrator/http/server/handler/expression"
	"github.com/labstack/echo/v4"
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
	s := echo.New()
	configRoute := s.Group("/config")
	configRoute.GET("/get", configHandler.GetConfigRequestHandler(config.CalculatorConfig))
	configRoute.PUT("/set", configHandler.SetConfigRequestHandler(config.CalculatorConfig))
	agentsRoute := s.Group("/agent")
	agentsRoute.GET("/list", agents.GetListAgentsHandler(config.DB.Agents))
	expressionsRoute := s.Group("/expression")
	expressionsRoute.POST("/new", expression.NewExpressionHandler(config.Producer, config.DB))
	return &Server{
		server: s,
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
	if err := validatePort(s.config.Port); err != nil {
		return err
	}
	return s.server.Start(":" + strconv.Itoa(s.config.Port))
}
