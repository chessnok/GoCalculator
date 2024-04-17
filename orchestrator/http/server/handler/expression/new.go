package expression

import (
	"github.com/chessnok/GoCalculator/orchestrator/internal/db"
	"github.com/chessnok/GoCalculator/orchestrator/internal/expressions"
	user2 "github.com/chessnok/GoCalculator/orchestrator/internal/user"
	"github.com/labstack/echo/v4"
)

type newExpressionRequest struct {
	Expression string `json:"expression"`
}

func NewExpressionHandler(postgres *db.Postgres) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !c.Get("loggedIn").(bool) {
			return c.JSON(401, map[string]string{"error": "Unauthorized"})
		}
		req := newExpressionRequest{}
		uid := c.Get("user").(*user2.User).ID
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, map[string]string{"error": "invalid request"})
		}
		expr, err := expressions.NewExpression(req.Expression, uid)
		if err != nil {
			return c.JSON(400, map[string]string{"error": "Invalid expression, " + err.Error()})
		}
		err = postgres.Expressions.New(expr)
		if err != nil {
			return c.JSON(500, map[string]string{"error": "Internal server error"})
		}
		err = postgres.Tasks.New(expr.Tasks)
		if err != nil {
			return c.JSON(500, map[string]string{"error": "Internal server error"})
		}
		return c.JSON(200, expr)
	}
}
