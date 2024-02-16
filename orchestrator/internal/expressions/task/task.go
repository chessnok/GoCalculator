package task

import (
	"github.com/chessnok/GoCalculator/orchestrator/internal/expressions/structure"
	"github.com/google/uuid"
	"strconv"
)

type Operation struct {
	Id        string
	A         any
	B         any
	Operation string
	next      *Operation
}
type Task struct {
	Id           string  `json:"id"`
	Operation    string  `json:"operation"`
	A            float64 `json:"a"`
	B            float64 `json:"b"`
	ExprId       string  `json:"-"`
	AIsNumeral   bool    `json:"a_is_numeral"`
	BIsNumeral   bool    `json:"b_is_numeral"`
	NextTaskId   string  `json:"next_task_id"`
	NextTaskType bool    `json:"-"`
	IsFinal      bool    `json:"is_final"`
	Error        string  `json:"error"`
	Status       string  `json:"status"`
	Result       float64 `json:"result"`
}

func newOperation(a any, b any, operator string) *Operation {
	return &Operation{
		A:         a,
		B:         b,
		Operation: operator,
	}
}

func NewTask() *Task {
	return &Task{
		Id: uuid.New().String(),
	}
}
func tasksFromOperations(operations []*Operation) []*Task {
	var res = make([]*Task, 0, len(operations))
	ops := make(map[any]int)
	for i, operation := range operations {
		ops[operation] = i
	}
	for i, operation := range operations {
		res = append(res, NewTask())
		task := res[i]
		task.Operation = operation.Operation
		switch operation.A.(type) {
		case int:
			task.A = float64(operation.A.(int))
			task.AIsNumeral = true
		default:
			task.AIsNumeral = false
			res[ops[operation.A.(*Operation)]].NextTaskId = task.Id
			res[ops[operation.A.(*Operation)]].NextTaskType = false
		}
		switch operation.B.(type) {
		case int:
			task.B = float64(operation.B.(int))
			task.BIsNumeral = true
		default:
			task.BIsNumeral = false
			res[ops[operation.B.(*Operation)]].NextTaskId = task.Id
			res[ops[operation.B.(*Operation)]].NextTaskType = true
		}
	}
	res[len(res)-1].IsFinal = true
	return res
}
func GetTasks(expression []string) []*Task {
	LinkedListObj := structure.LinkedList{}
	var result []*Operation
	for i := 0; i < len(expression); i++ {
		if num, err := strconv.Atoi(expression[i]); err == nil {
			LinkedListObj.InsertAfter(LinkedListObj.Last, &structure.Node{Value: num})
		} else {
			LinkedListObj.InsertAfter(LinkedListObj.Last, &structure.Node{Value: expression[i]})
		}
	}

	for curN := LinkedListObj.Head; curN != nil; {
		switch curN.Value.(type) {
		case string:
			curN.Value = newOperation(curN.Prev.Prev.Value, curN.Prev.Value, curN.Value.(string))
			LinkedListObj.DeleteBefore(curN)
			LinkedListObj.DeleteBefore(curN)
			result = append(result, curN.Value.(*Operation))
		}
		curN = curN.Next
	}
	res := tasksFromOperations(result)
	return res
}
