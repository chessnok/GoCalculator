package user

import (
	"github.com/chessnok/GoCalculator/orchestrator/internal/db/table"
	"github.com/chessnok/GoCalculator/orchestrator/internal/user"
	"github.com/labstack/echo/v4"
)

func NewUserHandler(pg *table.Users) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(user.User)
		if err := c.Bind(req); err != nil {
			return c.JSON(400, "Bad request")
		}
		if req.Username == "" || req.Password == "" {
			return c.JSON(400, map[string]string{"error": "Bad request"})
		}
		usr, err := pg.NewUser(req.Username, req.Password)
		if err != nil {
			return c.JSON(500, map[string]string{"error": "Internal server error"})
		}
		uc := usr
		uc.Password = ""
		return c.JSON(200, uc)
	}
}
