package user

import (
	"github.com/chessnok/GoCalculator/orchestrator/http/token"
	"github.com/chessnok/GoCalculator/orchestrator/internal/db/table"
	"github.com/chessnok/GoCalculator/orchestrator/internal/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func LoginHandler(pg *table.Users, tokenManager *token.TokenManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(user.User)
		if err := c.Bind(&u); err != nil {
			return c.JSON(400, map[string]string{"status": "Bad request"})
		}
		if u.Username == "" || u.Password == "" {
			return c.JSON(400, map[string]string{"status": "Bad request"})
		}
		usr, err := pg.GetUserByUsername(u.Username)
		if err != nil {
			return c.JSON(500, map[string]string{"status": "Internal server error"})
		}
		if !table.ComparePasswords(usr.Password, u.Password) {
			return c.JSON(401, map[string]string{"status": "Unauthorized"})
		}
		tkn, err, exp := tokenManager.GenerateToken(usr)
		if err != nil {
			return c.JSON(500, map[string]string{"status": "Internal server error"})
		}
		cookie := new(http.Cookie)
		cookie.Name = "session"
		cookie.HttpOnly = true
		cookie.Expires = exp
		cookie.Value = tkn
		cookie.Path = "/"
		c.SetCookie(cookie)
		return c.JSON(200, map[string]string{"status": "sent cookie with token"})
	}
}
