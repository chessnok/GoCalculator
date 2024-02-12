package agents

import (
	"github.com/chessnok/GoCalculator/orchestrator/internal/db/table/agents"
	"github.com/labstack/echo/v4"
)

func GetListAgentsHandler(ag *agents.Agents) echo.HandlerFunc {
	return func(c echo.Context) error {
		a, err := ag.GetList()
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}
		return c.JSON(200, a)
	}
}