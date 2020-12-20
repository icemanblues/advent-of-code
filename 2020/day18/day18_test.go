package main

import (
	"testing"
)

func TestEvaluateLine(t *testing.T) {
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
			actual, _ := EvaluateLine(test.line, 0)
			if actual != test.expected {
				t.Errorf("expected %v, actual %v", test.expected, actual)
			}
		})
	}
}

func TestExpression(t *testing.T) {
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
	}

	for _, test := range tests {
		t.Run(test.line, func(t *testing.T) {
			actual := Express(test.line)
			if actual != test.expected {
				t.Errorf("expected %v, actual %v", test.expected, actual)
			}
		})
	}
}
