package agent

import (
	"github.com/chessnok/GoCalculator/orchestrator/internal/agents"
	"github.com/chessnok/GoCalculator/orchestrator/internal/db/table"
	"github.com/labstack/echo/v4"
	"os/exec"
)

func KillAgentHandler(pg *table.Agents) echo.HandlerFunc {
	return func(c echo.Context) error {
		a := new(agents.Agent)
		if err := c.Bind(&a); err != nil {
			return c.JSON(404, map[string]string{"status": "Bad Request"})
		}
		if a.ID == "" {
			return c.JSON(404, map[string]string{"status": "Bad Request"})
		}
		agent, err := pg.GetAgentById(a.ID)
		if err != nil {
			return c.JSON(500, map[string]string{
				"status": "Internal server error",
			})
		}
		if agent == nil {
			return c.JSON(404, map[string]string{"status": "Bad Request"})
		}

		cmd := exec.Command("kill", "-9", agent.Pid)
		if err := cmd.Start(); err != nil {
			return c.JSON(500, map[string]string{
				"status": "Internal server error",
			})
		}
		return c.JSON(200, map[string]string{"status": "ok"})
	}
}
