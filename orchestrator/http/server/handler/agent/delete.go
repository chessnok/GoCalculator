package agent

import (
	"github.com/chessnok/GoCalculator/orchestrator/internal/agents"
	"github.com/chessnok/GoCalculator/orchestrator/internal/db/table"
	"github.com/labstack/echo/v4"
)

func DeleteAgentHandler(pg *table.Agents) echo.HandlerFunc {
	return func(c echo.Context) error {
		a := new(agents.Agent)
		if err := c.Bind(&a); err != nil || a.ID == "" {
			return c.JSON(404, map[string]string{"status": "Bad request"})
		}
		err := pg.DeleteAgent(a.ID)
		if err != nil {
			return c.JSON(500, map[string]string{"status": "Internal server error"})
		}
		return c.JSON(200, map[string]string{"status": "ok"})
	}
}
