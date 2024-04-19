package calculator

import "time"

func (c *Calculator) calc(operation string, a, b float64) (float64, error) {
	switch operation {
	case "+":
		time.Sleep(time.Duration(c.Config.AddExecutionTime) * time.Second)
		return a + b, nil
	case "-":
		time.Sleep(time.Duration(c.Config.SubExecutionTime) * time.Second)
		return a - b, nil
	case "*":
		time.Sleep(time.Duration(c.Config.MulExecutionTime) * time.Second)
		return a * b, nil
	case "/":
		time.Sleep(time.Duration(c.Config.DivExecutionTime) * time.Second)
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil
	default:
		return 0, ErrInvalidOperation
	}
}
