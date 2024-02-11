package config

import (
	"github.com/chessnok/GoCalculator/agent/pkg/calculator"
	"github.com/labstack/echo/v4"
)

func GetConfigRequestHandler(config *calculator.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, config)
	}
}
