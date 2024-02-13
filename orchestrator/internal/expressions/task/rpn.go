package task

import (
	"errors"
	"github.com/chessnok/GoCalculator/orchestrator/internal/expressions/structure"
)

var (
	priority = map[string]int{
		"+": 2,
		"-": 2,
		"*": 3,
		"/": 3,
		"(": 1,
		"4": 1,
	}
)

func getRpn(expression string) string {
	res := ""
	st := structure.NewStack()
	for _, symbol := range expression {
		switch symbol {
		case '(':
			st.Push(string(symbol))
		case ')':
			for st.Len() > 0 && st.Peek() != "(" {
				res += st.Pop().(string) + " "
			}
			st.Pop()
		case '+', '-', '*', '/':
			for st.Len() > 0 && priority[st.Peek().(string)] >= priority[string(symbol)] {
				res += st.Pop().(string) + " "
			}
			st.Push(string(symbol))
		default:
			res += string(symbol) + " "
		}
	}
	for st.Len() > 0 {
		if st.Peek() == "(" {
			panic(errors.New("parentheses error"))
		}
		res += st.Pop().(string) + " "
	}
	if res == "" {
		panic(errors.New("empty expression"))
	}
	return res[:len(res)-1]
}
func GetReversePolishNotation(expression string) (string, error) {
	res := getRpn(expression)
	if err := recover(); err != nil {
		return "", err.(error)
	}
	return res, nil
}

func main() {
	exp := "2+4+6*5"
	rpn, err := GetReversePolishNotation(exp)
	if err != nil {
		panic(err)
	}
	println(rpn)
}
