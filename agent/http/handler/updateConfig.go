package handler

import (
	"github.com/chessnok/GoCalculator/agent/pkg/calculator"
	"github.com/labstack/echo"
)

func UpdateConfigHandler(config *calculator.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		nc := new(calculator.Config)
		if err := c.Bind(nc); err != nil {
			return c.JSON(400, map[string]string{"error": "invalid request"})
		}
		nc.ParallelWorkers = config.ParallelWorkers
		*config = *nc
		return c.JSON(200, map[string]string{"message": "config updated"})
	}
}
