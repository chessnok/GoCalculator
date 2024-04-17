package server

import (
	"errors"
	"github.com/chessnok/GoCalculator/orchestrator/http/server/handler/agent"
	configHandler "github.com/chessnok/GoCalculator/orchestrator/http/server/handler/config"
	"github.com/chessnok/GoCalculator/orchestrator/http/server/handler/expression"
	"github.com/chessnok/GoCalculator/orchestrator/http/server/handler/user"
	"github.com/chessnok/GoCalculator/orchestrator/http/token"
	user2 "github.com/chessnok/GoCalculator/orchestrator/internal/user"
	"github.com/labstack/echo/v4"
	"os"
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
	tokenManager := token.NewTokenManager(os.Getenv("GO_SECRET_KEY"))
	s := echo.New()
	s.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})
	s.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie("session")
			if err != nil || cookie.Value == "" {
				c.Set("loggedIn", false)
				return next(c)
			}
			valid, userId, err := tokenManager.CheckToken(cookie.Value)
			if err != nil || !valid {
				c.Set("loggedIn", false)
				return next(c)
			}
			usr, err := config.DB.Users.GetUserById(userId)
			if err != nil {
				c.Set("loggedIn", false)
				return next(c)
			}
			c.Set("loggedIn", true)
			c.Set("user", usr)
			return next(c)
		}
	})
	adminMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !c.Get("loggedIn").(bool) {
				return c.JSON(401, map[string]string{"status": "Unauthorized"})
			}
			usr := c.Get("user").(*user2.User)
			if usr == nil || !usr.IsAdmin {
				return c.JSON(401, map[string]string{"status": "Unauthorized"})
			}
			return next(c)
		}
	}
	adminRoute := s.Group("/admin")
	adminRoute.Use(adminMiddleware)
	configRoute := adminRoute.Group("/config")
	configRoute.GET("/get", configHandler.GetConfigRequestHandler(config.CalculatorConfig))
	configRoute.PUT("/set", configHandler.SetConfigRequestHandler(config.CalculatorConfig))
	agentsRoute := s.Group("/agent")
	agentsRoute.GET("/list", agent.GetListAgentsHandler(config.DB.Agents))
	agentsAdminRoute := adminRoute.Group("/agent")
	agentsAdminRoute.DELETE("/delete", agent.DeleteAgentHandler(config.DB.Agents))
	agentsAdminRoute.DELETE("/kill", agent.KillAgentHandler(config.DB.Agents))
	agentsAdminRoute.POST("/new", agent.NewAgentHandler())
	expressionsRoute := s.Group("/expression")
	expressionsRoute.POST("/new", expression.NewExpressionHandler(config.DB))
	expressionsRoute.GET("/list", expression.GetListExpressionsHandler(config.DB.Expressions))
	userRoute := s.Group("/user")
	userRoute.GET("/me", user.MeHandler())
	userRoute.POST("/new", user.NewUserHandler(config.DB.Users))
	userRoute.POST("/login", user.LoginHandler(config.DB.Users, tokenManager))
	userRoute.POST("/set_is_admin", user.SetIsAdminHandler(config.DB.Users, os.Getenv("GO_API_KEY"))) // ONLY WITH API KEY
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
func (s *Server) Start() {
	if err := validatePort(s.config.Port); err != nil {
		return
	}
	s.server.Start(":" + strconv.Itoa(s.config.Port))
}
