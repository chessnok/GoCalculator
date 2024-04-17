package user

import (
	"github.com/chessnok/GoCalculator/orchestrator/internal/db/table"
	"github.com/chessnok/GoCalculator/orchestrator/internal/user"
	"github.com/labstack/echo/v4"
)

func SetIsAdminHandler(pg *table.Users, apiKey string) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("Authorization") != apiKey {
			return c.JSON(401, map[string]string{"status": "Unauthorized"})
		}
		req := new(user.User)
		if err := c.Bind(req); err != nil {
			return c.JSON(400, map[string]string{"status": "Bad request"})
		}
		if req.Username == "" {
			return c.JSON(400, map[string]string{"status": "Bad request"})
		}
		err := pg.SetIsAdmin(req.Username, req.IsAdmin)
		if err != nil {
			return c.JSON(500, map[string]string{"status": "Internal server error"})
		}
		return c.JSON(200, map[string]string{"status": "Updated user if exists"})
	}

}
