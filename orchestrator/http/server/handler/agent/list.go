package agent

import (
	"github.com/chessnok/GoCalculator/orchestrator/internal/db/table"
	"github.com/labstack/echo/v4"
)

func GetListAgentsHandler(ag *table.Agents) echo.HandlerFunc {
	return func(c echo.Context) error {
		a, err := ag.GetAgentsList()
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}
		return c.JSON(200, a)
	}
}
