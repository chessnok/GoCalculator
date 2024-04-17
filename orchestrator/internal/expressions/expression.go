package expressions

import (
	"github.com/chessnok/GoCalculator/orchestrator/internal/expressions/task"
	"github.com/google/uuid"
	"strings"
	"time"
)

type Expression struct {
	Id                   string       `json:"id" xml:"id"`
	Expression           string       `json:"expression" xml:"expression"`
	NormalizedExpression string       `json:"normalized_expression" xml:"normalized_expression"`
	ResultTaskId         string       `json:"-" xml:"-"`
	Status               string       `json:"status" xml:"status"`
	Result               float64      `json:"result" xml:"result"`
	CreatedAt            time.Time    `json:"created_at" xml:"created_at"`
	Tasks                []*task.Task `json:"tasks" xml:"tasks"`
	UserId               string       `json:"user_id" xml:"user_id"`
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

func NewExpression(expression, userid string) (*Expression, error) {
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
		Tasks:                tsks,
		Status:               "pending",
		ResultTaskId:         tsks[len(tsks)-1].Id,
		UserId:               userid,
	}, err
}
