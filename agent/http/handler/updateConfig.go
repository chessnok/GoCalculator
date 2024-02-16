package handler

import (
	"github.com/chessnok/GoCalculator/agent/pkg/calculator"
	"github.com/labstack/echo"
)

func UpdateConfigHandler(calc *calculator.Calculator) echo.HandlerFunc {
	return func(c echo.Context) error {
		nc := new(calculator.Config)
		if err := c.Bind(&nc); err != nil {
			return c.JSON(400, map[string]string{"error": "invalid request"})
		}
		nc.ParallelWorkers = calc.Config.ParallelWorkers
		*calc.Config = *nc
		return c.JSON(200, map[string]string{"task": calc.LastOperationID})
	}
}
