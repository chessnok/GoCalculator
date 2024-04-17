package user

import (
	user2 "github.com/chessnok/GoCalculator/orchestrator/internal/user"
	"github.com/labstack/echo/v4"
)

func MeHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		if !c.Get("loggedIn").(bool) {
			return c.JSON(401, map[string]string{"status": "Unauthorized"})
		}
		usr := c.Get("user").(*user2.User)
		usr.Password = ""
		return c.JSON(200, usr)
	}
}
