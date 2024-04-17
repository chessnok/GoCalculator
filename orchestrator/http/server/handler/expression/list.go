package expression

import (
	"github.com/chessnok/GoCalculator/orchestrator/internal/db/table"
	expressions2 "github.com/chessnok/GoCalculator/orchestrator/internal/expressions"
	user2 "github.com/chessnok/GoCalculator/orchestrator/internal/user"
	"github.com/labstack/echo/v4"
)

func GetListExpressionsHandler(postgres *table.Expressions) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !c.Get("loggedIn").(bool) {
			return c.JSON(401, map[string]string{"error": "Unauthorized"})
		}
		var expressions []expressions2.Expression
		var err error
		usr := c.Get("user").(*user2.User)
		if usr.IsAdmin {
			expressions, err = postgres.GetExpressionsList()
		} else {
			expressions, err = postgres.GetExpressionsListByUserId(usr.ID)
		}
		if err != nil {
			return c.JSON(500, map[string]string{"status": err.Error()})
		}
		return c.JSON(200, expressions)
	}
}
