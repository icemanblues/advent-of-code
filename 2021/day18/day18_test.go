package main

import (
	"testing"
)

// numberEqual gross way to test for equality
func numberEqual(a, b *Number) bool {
	return a.String() == b.String()
}

func TestAddNumbers(t *testing.T) {
	a := NewPair(1, 2)
	b := Pair(NewPair(3, 4), NewNumber(5))
	c := AddNumbers(a, b)

	expected := Pair(NewPair(1, 2), Pair(NewPair(3, 4), NewNumber(5)))
	if !numberEqual(c, expected) {
		t.Errorf("c: %v expected: %v", c, expected)
	}
}

func TestParseNumber(t *testing.T) {
	tests := []struct {
		number   string
		expected *Number
	}{
		{"[9,1]", NewPair(9, 1)},
		{"[[1,2],[[3,4],5]]", Pair(NewPair(1, 2), Pair(NewPair(3, 4), NewNumber(5)))},
	}

	for _, test := range tests {
		t.Run(test.number, func(t *testing.T) {
			actual := ParseNumber(test.number)
			if !numberEqual(actual, test.expected) {
				t.Errorf("actual: %v expected: %v", actual, test.expected)
			}
		})
	}
}

func TestMagnitude(t *testing.T) {
	tests := []struct {
		number   *Number
		expected int
	}{
		{
			number:   NewPair(9, 1),
			expected: 29,
		},
		{
			number:   NewPair(1, 9),
			expected: 21},
		{
			number:   Pair(NewPair(9, 1), NewPair(1, 9)),
			expected: 129,
		},
	}

	for _, test := range tests {
		t.Run(test.number.String(), func(t *testing.T) {
			actual := test.number.Magnitude()
			if actual != test.expected {
				t.Errorf("actual: %v expected: %v", actual, test.expected)
			}
		})
	}
}

func TestMagnitudeString(t *testing.T) {
	tests := []struct {
		number   string
		expected int
	}{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	}

	for _, test := range tests {
		t.Run(test.number, func(t *testing.T) {
			n := ParseNumber(test.number)
			actual := n.Magnitude()
			if actual != test.expected {
				t.Errorf("actual: %v expected: %v", actual, test.expected)
			}
		})
	}
}
