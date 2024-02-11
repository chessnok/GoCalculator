package config

import (
	"github.com/chessnok/GoCalculator/agent/pkg/calculator"
	"github.com/labstack/echo/v4"
)

func SetConfigRequestHandler(config *calculator.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := c.Bind(config); err != nil {
			return c.JSON(400, err)
		}
		return c.JSON(200, config)
	}
}
