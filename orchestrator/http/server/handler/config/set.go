package config

import (
	"github.com/chessnok/GoCalculator/agent/pkg/calculator"
	"github.com/chessnok/GoCalculator/orchestrator/internal/agents/manager"
	"github.com/labstack/echo/v4"
)

func SetConfigRequestHandler(mangr *manager.AgentManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		var newConfig calculator.Config
		if err := c.Bind(&newConfig); err != nil {
			return c.JSON(400, err)
		}
		newConfig.ParallelWorkers = mangr.CalculatorConfig.ParallelWorkers
		*mangr.CalculatorConfig = newConfig
		mangr.UpdateConfig()
		return c.JSON(200, newConfig)
	}
}
