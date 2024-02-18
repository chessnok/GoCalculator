package task

import (
	"github.com/chessnok/GoCalculator/orchestrator/internal/expressions/structure"
	"github.com/google/uuid"
	"strconv"
)

type operation struct {
	Id        string
	A         any
	B         any
	Operation string
	next      *operation
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

func newOperation(a any, b any, operator string) *operation {
	return &operation{
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
func tasksFromOperations(operations []*operation) []*Task {
	var res = make([]*Task, 0, len(operations))
	ops := make(map[any]int)
	for i, operation := range operations {
		ops[operation] = i
	}
	for i, oprtn := range operations {
		res = append(res, NewTask())
		task := res[i]
		task.Operation = oprtn.Operation
		switch oprtn.A.(type) {
		case int:
			task.A = float64(oprtn.A.(int))
			task.AIsNumeral = true
		default:
			task.AIsNumeral = false
			res[ops[oprtn.A.(*operation)]].NextTaskId = task.Id
			res[ops[oprtn.A.(*operation)]].NextTaskType = false
		}
		switch oprtn.B.(type) {
		case int:
			task.B = float64(oprtn.B.(int))
			task.BIsNumeral = true
		default:
			task.BIsNumeral = false
			res[ops[oprtn.B.(*operation)]].NextTaskId = task.Id
			res[ops[oprtn.B.(*operation)]].NextTaskType = true
		}
	}
	res[len(res)-1].IsFinal = true
	return res
}
func GetTasks(expression []string) []*Task {
	LinkedListObj := structure.LinkedList{}
	var result []*operation
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
			result = append(result, curN.Value.(*operation))
		}
		curN = curN.Next
	}
	res := tasksFromOperations(result)
	return res
}
