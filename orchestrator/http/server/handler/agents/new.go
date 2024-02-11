package agents

type NewAgentRequest struct {
	WorkersCount int `json:"workers_count" validate:"required,min=1,max=1000"`
}

// NewAgent returns a handler that creates a new orchestrator
//func NewAgent(am *agents.Manager) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		r := new(NewAgentRequest)
//		if err := c.Bind(r); err != nil {
//			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
//		}
//		if err := c.Validate(r); err != nil {
//			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
//		}
//		a := am.CreateAgent(r.WorkersCount)
//		return c.JSON(http.StatusCreated, a)
//	}
//
//}
