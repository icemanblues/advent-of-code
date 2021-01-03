package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "01"
	dayTitle = "Not Quite Lisp"
)

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	lines, _ := util.ReadInput("input.txt")
	c := 0
	p, f := 1, true
	for i, r := range lines[0] {
		switch r {
		case '(':
			c++
		case ')':
			c--
		}
		if f && c == -1 {
			p = i
			f = false
		}
	}
	fmt.Printf("Part 1: %v\n", c)
	fmt.Printf("Part 2: %v\n", p+1)
}
