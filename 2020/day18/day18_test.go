package main

import (
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		line     string
		expected int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
		{"(1 + 2) * (3 + 4) * (5 + 6)", 231},
	}

	for _, test := range tests {
		t.Run(test.line, func(t *testing.T) {
			actual, _ := Eval(test.line, 0, false)
			if actual != test.expected {
				t.Errorf("expected %v, actual %v", test.expected, actual)
			}
		})
	}
}

//6 * 7 + 7 * (9 * 6 + 6 + 4 * 6) * 5 * 4
//6 *  14   * (9 * 6 + 6 + 4 * 6) * 5 * 4
//6 *  14   * (9 *     16    * 6) * 5 * 4
//6 *  14   * (9 *     16    * 6) * 5 * 4
//6 * 14 * 864 * 5 * 4

// 5 * 7 + 2 + (8 + (8 + 4 + 3 * 9 + 8 + 7) + (4 * 2 * 6 + 2) * 8 + 3 + 9) * (4 * 9 + (5 + 7 + 9)) * 2
// 5 * 9 + (8 + (15 * 24) + (4 * 2 * 8) * 20) * (4 * 9 + (21)) * 2
// 5 * 9 + (8 + 360 + 64 * 20) * (4 * 30) * 2
// 5 * 9 + (432 * 20) * 120 * 2
// 5 * 8649 * 120 * 2

func TestEvalLookahead(t *testing.T) {
	tests := []struct {
		line     string
		expected int
	}{
		{"1 + 2 * 3 + 4 * 5 + 6", 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
		{"(1 + 2) * (3 + 4) * (5 + 6)", 231},
		{"6 * 7 + 7 * (9 * 6 + 6 + 4 * 6) * 5 * 4", 1451520},
		{"7 * 2 + 4", 42},
		{"5 * 7 + 2 + (8 + (8 + 4 + 3 * 9 + 8 + 7) + (4 * 2 * 6 + 2) * 8 + 3 + 9) * (4 * 9 + (5 + 7 + 9)) * 2", 10378800},
	}

	for _, test := range tests {
		t.Run(test.line, func(t *testing.T) {
			actual := SumExpressions([]string{test.line}, true)
			if actual != test.expected {
				t.Errorf("expected %v, actual %v", test.expected, actual)
			}
		})
	}
}

func TestStack(t *testing.T) {
	stack := Stack{}
	if !stack.IsEmpty() {
		t.Errorf("New Stack should be empty and yet its not. %v\n", stack.s)
	}

	stack.Push('r')
	if stack.IsEmpty() {
		t.Errorf("Just added to the stack, it should NOT be empty")
	}

	if a := stack.Peek(); a != 'r' {
		t.Errorf("Peeking at top of stack. Expecting r. Actual %c", a)
	}
	if stack.IsEmpty() {
		t.Errorf("Did a peek, so the stack should not be empty")
	}

	stack.Push('g')
	stack.Push('k')
	if stack.IsEmpty() {
		t.Errorf("Just added to the stack, it should NOT be empty")
	}
	if l := stack.Len(); l != 3 {
		t.Errorf("Expecting 3 when pushing 3 elements. actual: %v\n", l)
	}

	if a := stack.Pop(); a != 'k' {
		t.Errorf("expecting k, actual %v\n", a)
	}

	if b := stack.Pop(); b != 'g' {
		t.Errorf("expecting g, actual %v\n", b)
	}

	if c := stack.Pop(); c != 'r' {
		t.Errorf("expecting r, actual %v\n", c)
	}

	if !stack.IsEmpty() {
		t.Errorf("After push 3 and pop 3, the stack should be empty. %v\n", stack.s)
	}

	// another pop should result in a panic
}
