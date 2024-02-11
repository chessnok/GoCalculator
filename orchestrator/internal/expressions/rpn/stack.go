package rpn

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value string
		prev  *node
	}
)

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (this *Stack) Len() int {
	return this.length
}

func (this *Stack) Peek() string {
	if this.length == 0 {
		return ""
	}
	return this.top.value
}

func (this *Stack) Pop() string {
	if this.length == 0 {
		return ""
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

func (this *Stack) Push(value string) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}
