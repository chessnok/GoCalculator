package handler

import (
	"github.com/chessnok/GoCalculator/agent/pkg/calculator"
	"github.com/labstack/echo"
)

func LastOperationHandler(calc *calculator.Calculator) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, map[string]string{"operation": calc.LastOperationID})
	}
}
