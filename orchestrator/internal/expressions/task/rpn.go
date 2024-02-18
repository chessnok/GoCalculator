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

func splitDigits(expression string) []string {
	res := []string{}
	num := ""
	for _, symbol := range expression {
		if symbol >= '0' && symbol <= '9' {
			num += string(symbol)
		} else {
			if num != "" {
				res = append(res, num)
			}
			num = ""
			res = append(res, string(symbol))
		}
	}
	if num != "" {
		res = append(res, num)
	}
	return res
}

func getRpn(expression string) string {
	res := ""
	splitedExpression := splitDigits(expression)
	st := structure.NewStack()
	for _, symbol := range splitedExpression {
		switch symbol {
		case "(":
			st.Push(string(symbol))
		case ")":
			for st.Len() > 0 && st.Peek() != "(" {
				res += st.Pop().(string) + " "
			}
			st.Pop()
		case "+", "-", "*", "/":
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
