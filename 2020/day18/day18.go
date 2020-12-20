package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "18"
	dayTitle = "Operation Order"
)

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

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

func EvaluateLine(line string, i int) (int, int) {
	runes := []rune(line)
	var e int

	if runes[i] == '(' {
		e, i = EvaluateLine(line, i+1)

	} else {
		e = MustAtoi(string(runes[i]))
	}
	i++

	var op rune
	for i < len(line) {
		r := runes[i]

		switch r {
		case '(':
			exp, j := EvaluateLine(line, i+1)
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
			right := MustAtoi(string(r))
			e = operate(op, e, right)
		}
		i++
	}

	return e, i
}

type Expression interface {
	Evaluate() int
}

type Value struct {
	v int
}

func (value Value) Evaluate() int {
	return value.v
}

type BinOp struct {
	op    rune
	left  Expression
	right Expression
}

func (b BinOp) Evaluate() int {
	return operate(b.op, b.left.Evaluate(), b.right.Evaluate())
}

func Express(line string) int {
	newline := strings.ReplaceAll(line, " * ", ") * (")
	newline = "(" + newline + ")"
	fmt.Println(line)
	fmt.Println(newline)
	e, _ := EvaluateLine(newline, 0)
	return e
}

func part1() {
	sum := 0
	lines, _ := util.ReadInput("input.txt")
	for _, line := range lines {
		n, _ := EvaluateLine(line, 0)
		sum += n
	}
	fmt.Printf("Part 1: %v\n", sum)
}

func part2() {
	sum := 0
	lines, _ := util.ReadInput("input.txt")
	for _, line := range lines {
		n, _ := EvaluateLine(line, 0)
		sum += n
	}
	fmt.Printf("Part 2: %v\n", sum)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
