package rpn

var (
	priority = map[string]int{
		"+": 2,
		"-": 2,
		"*": 3,
		"/": 3,
		"(": 4,
		"4": 4,
	}
)

func getRpn(expression string) string {
	res := ""
	st := NewStack()
	for _, symbol := range expression {
		switch symbol {
		case '(':
			st.Push(string(symbol))
		case ')':
			for st.Len() > 0 && st.Peek() != "(" {
				res += st.Pop()
			}
			st.Pop()
		case '+', '-', '*', '/':
			for st.Len() > 0 && priority[st.Peek()] >= priority[string(symbol)] {
				res += st.Pop()
			}
			st.Push(string(symbol))
		default:
			res += string(symbol)
		}
	}
	return res
}
func GetReversePolishNotation(expression string) (string, error) {
	res := getRpn(expression)
	if err := recover(); err != nil {
		return "", err.(error)
	}
	return res, nil
}
