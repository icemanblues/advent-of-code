package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "18"
	dayTitle = "Operation Order"
)

func operate(op rune, a, b int) int {
	switch op {
	case '+':
		return a + b
	case '*':
		return a * b
	}
	fmt.Printf("oops! %c %v %v\n", op, a, b)
	panic("unknown operator: " + string(op))
}

// Eval evaluates the math expression from left to right. No operator precedence
func Eval(line string, i int, look bool) (int, int) {
	runes := []rune(line)
	var e int

	if runes[i] == '(' {
		e, i = Eval(line, i+1, look)

	} else {
		e = util.MustAtoi(string(runes[i]))
	}
	i++

	var op rune
	for i < len(line) {
		r := runes[i]

		switch r {
		case '(':
			exp, j := Eval(line, i+1, look)
			e = operate(op, e, exp)
			i = j
		case ')':
			return e, i
		case '+':
			op = r
		case '*':
			op = r
		case ' ':
		default: // number
			// do a look ahead if current op is multiply and the next op is +
			// * 5 + 8
			//   ^^^
			if look && op == '*' {
				// need to find the next operator. parens dont count

			}

			if look && op == '*' && i+2 < len(runes) && runes[i+2] == '+' {
				exp, j := Eval(line, i, look)
				e = operate(op, e, exp)
				i = j
			} else {
				right := util.MustAtoi(string(r))
				e = operate(op, e, right)
			}
		}

		i++
	}

	return e, i
}

type Stack struct {
	s []rune
}

func (s *Stack) Len() int {
	return len(s.s)
}

func (s *Stack) IsEmpty() bool {
	return len(s.s) == 0
}

func (s *Stack) Push(r rune) {
	s.s = append(s.s, r)
}

func (s *Stack) Pop() rune {
	if !s.IsEmpty() {
		r := s.s[len(s.s)-1]
		s.s = s.s[:len(s.s)-1]
		return r
	}
	panic("Cannot pop on an empty stack")
}

func (s *Stack) Peek() rune {
	if !s.IsEmpty() {
		return s.s[len(s.s)-1]
	}
	panic("Cannot peek on an empty stack")
}

// EvalPrecedence evaluates the mathematical expression using Shunting Yard Algorithm
func EvalPrecedence(line string) []rune {
	numbers := []rune{}
	operators := Stack{}
	for _, r := range line {
		switch r {
		case '+':
			operators.Push(r)
		case '*':
			// this is lower precedence, so we need to peek to see if the previous operator was a +
			if !operators.IsEmpty() && operators.Peek() == '+' {
				// pop all operators and put them into the numbers (output)
				// and put back the left parens
				for !operators.IsEmpty() {
					op := operators.Pop()
					if op == '(' {
						operators.Push('(')
						break
					}
					numbers = append(numbers, op)
				}
			}
			operators.Push(r)
		case ' ':
			// skip
		case '(':
			operators.Push(r)
		case ')':
			// pop all operators and put them into the numbers (output) until you see the open paren
			for !operators.IsEmpty() {
				o := operators.Pop()
				if o == '(' {
					break
				}
				numbers = append(numbers, o)
			}
		default: // number
			numbers = append(numbers, r)
		}
	}

	// pop all operators to finish
	for !operators.IsEmpty() {
		op := operators.Pop()
		if op == '(' {
			fmt.Printf("This op %v shouldn't be here. %v\n", op, line)
		}
		numbers = append(numbers, op)
	}

	return numbers
}

func PostFixEval(expression []rune) int {
	var numbers []int
	for _, e := range expression {
		switch e {
		case '+':
			one := numbers[len(numbers)-1]
			two := numbers[len(numbers)-2]
			numbers = numbers[:len(numbers)-2]
			result := one + two
			numbers = append(numbers, result)
		case '*':
			one := numbers[len(numbers)-1]
			two := numbers[len(numbers)-2]
			numbers = numbers[:len(numbers)-2]
			result := one * two
			numbers = append(numbers, result)
		default: // number
			n := util.MustAtoi(string(e))
			numbers = append(numbers, n)
		}
	}
	return numbers[0]
}

func SumExpressions(expressions []string, look bool) int {
	sum := 0
	for _, exp := range expressions {
		if look {
			postfix := EvalPrecedence(exp)
			n := PostFixEval(postfix)
			sum += n
		} else {
			n, _ := Eval(exp, 0, false)
			sum += n
		}
	}
	return sum
}

func part1() {
	lines, _ := util.ReadInput("input.txt")
	fmt.Printf("Part 1: %v\n", SumExpressions(lines, false))
}

func part2() {
	lines, _ := util.ReadInput("input.txt")
	fmt.Printf("Part 2: %v\n", SumExpressions(lines, true))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
