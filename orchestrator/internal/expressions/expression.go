package expressions

import (
	"github.com/chessnok/GoCalculator/orchestrator/internal/expressions/task"
	"github.com/google/uuid"
	"strings"
)

type Expression struct {
	// Id - unique identifier
	Id string `json:"id" xml:"id"`
	// Expression - expressions to calculate
	Expression string `json:"expression" xml:"expression"`
	// NormalizedExpression - expressions after normalization
	NormalizedExpression string `json:"normalized_expression" xml:"normalized_expression"`
	// IsValid - is expressions valid
	IsValid bool `json:"is_valid" xml:"is_valid"`
	//	ResultTaskId - id of the task that will be executed to calculate the expression
	ResultTaskId string `json:"-" xml:"-"`
	//	Tasks		- list of tasks that will be executed to calculate the expression
	Tasks []*task.Task `json:"-"`
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

func NewExpression(expression string) (*Expression, error) {
	normalizedExpression, err := normalizeExpression(expression)
	if err != nil {
		return nil, err
	}
	rpn, err := task.GetReversePolishNotation(normalizedExpression)
	if err != nil {
		return nil, err
	}
	tsks := task.GetTasks(strings.Split(rpn, " "))
	uid := uuid.New().String()
	for _, tsk := range tsks {
		tsk.ExprId = uid
	}
	return &Expression{
		Id:                   uid,
		Expression:           expression,
		NormalizedExpression: normalizedExpression,
		IsValid:              true,
		Tasks:                tsks,
		ResultTaskId:         tsks[len(tsks)-1].Id,
	}, err
}
