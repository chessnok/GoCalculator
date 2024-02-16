package expression

import (
	"github.com/chessnok/GoCalculator/orchestrator/internal/db/table"
	"github.com/labstack/echo/v4"
)

func GetListExpressionsHandler(postgres *table.Expressions) echo.HandlerFunc {
	return func(c echo.Context) error {
		expressions, err := postgres.GetExpressionsList()
		if err != nil {
			return c.JSON(500, map[string]string{"error": err.Error()})
		}
		return c.JSON(200, expressions)
	}
}
