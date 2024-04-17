package config

import (
	agent_proto "github.com/chessnok/GoCalculator/proto"
	"github.com/labstack/echo/v4"
)

func SetConfigRequestHandler(config *agent_proto.Config) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := c.Bind(config); err != nil {
			return c.JSON(400, err)
		}
		return c.JSON(200, config)
	}
}
