package expressions

import (
	"testing"
)

type OperationsExample struct {
	A          float64
	B          float64
	Operation  string
	AIsNumeral bool
	BIsNumeral bool
}

type ExpressionExample struct {
	name       string
	in         string
	operations []OperationsExample
}

func TestExpression(t *testing.T) {
	tests := []ExpressionExample{
		{"test1", "1+1", []OperationsExample{{1, 1, "+", true, true}}},
		{"test2", "1+ 2", []OperationsExample{{1, 2, "+", true, true}}},
		{"test3", "1+2-3", []OperationsExample{{1, 2, "+", true, true}, {0, 3, "-", false, true}}},
		{"test4", "1+ 2*3", []OperationsExample{{2, 3, "*", true, true}, {1, 0, "+", true, false}}},
		{"test5", "1+2*3 /4", []OperationsExample{{2, 3, "*", true, true}, {0, 4, "/", false, true}, {1, 0, "+", true, false}}},
		{"test6", "1 + 0", []OperationsExample{{1, 0, "+", true, true}}}}

	for _, c := range tests {
		got, _ := NewExpression(c.in, "1")
		for i, task := range got.Tasks {
			if task.A != c.operations[i].A {
				t.Errorf("Test %s, expression %s: A got %f, want %f", c.name, c.in, task.A, c.operations[i].A)
			} else if task.B != c.operations[i].B {
				t.Errorf("Test %s, expression %s: B got %f, want %f", c.name, c.in, task.B, c.operations[i].B)
			} else if task.Operation != c.operations[i].Operation {
				t.Errorf("Test %s, expression %s: Operation got %s, want %s", c.name, c.in, task.Operation, c.operations[i].Operation)
			} else if task.AIsNumeral != c.operations[i].AIsNumeral {
				t.Errorf("Test %s, expression %s: AIsNumeral got %t, want %t", c.name, c.in, task.AIsNumeral, c.operations[i].AIsNumeral)
			} else if task.BIsNumeral != c.operations[i].BIsNumeral {
				t.Errorf("Test %s, expression %s: BIsNumeral got %t, want %t", c.name, c.in, task.BIsNumeral, c.operations[i].BIsNumeral)
			}
		}
	}
}
