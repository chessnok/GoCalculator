package agent

import (
	"github.com/labstack/echo/v4"
	"os/exec"
)

// NewAgentHandler returns a handler that creates a new orchestrator
func NewAgentHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		cmd := exec.Command("/agent")
		if err := cmd.Start(); err != nil {
			return c.JSON(500, map[string]string{
				"status": "Internal server error",
			})
		}
		return c.JSON(200, map[string]string{"status": "ok"})
	}
}
