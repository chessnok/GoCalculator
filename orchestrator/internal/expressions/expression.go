package expressions

type Expression struct {
	// ID - unique identifier
	ID int
	// Expression - expressions to calculate
	Expression string
	// NormalizedExpression - expressions after normalization
	NormalizedExpression string
	// IsValid - is expressions valid
	IsValid bool
}
type Operation struct {
	ID           int
	A            float64
	B            float64
	Operator     string
	IsError      bool
	Next         *Operation
	NextTaskType bool
}

func NewExpression(id int, expression string) (*Expression, error) {
	normalizedExpression, err := normalizeExpression(expression)
	return &Expression{
		ID:                   id,
		Expression:           expression,
		NormalizedExpression: normalizedExpression,
		IsValid:              err == nil,
	}, err
}
