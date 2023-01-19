package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "21"
	dayTitle = "Monkey Math"
)

type Math struct {
	Op   rune
	A, B string
}

func (m Math) String() string {
	return fmt.Sprintf("{%v %c %v}", m.A, m.Op, m.B)
}

type MonkeyMath struct {
	Math map[string]Math
	Nums map[string]int
}

func (mm *MonkeyMath) Clone() MonkeyMath {
	clone := MonkeyMath{make(map[string]Math), make(map[string]int)}
	for k, v := range mm.Nums {
		clone.Nums[k] = v
	}

	clone.Math = mm.Math // this never changes, so reuse it
	return clone
}

func (mm *MonkeyMath) DoMath(key string) (int, bool) {
	if v, ok := mm.Nums[key]; ok {
		return v, true
	}

	math, ok := mm.Math[key]
	if !ok {
		return 0, false // no formula found
	}
	a, aok := mm.Nums[math.A]
	b, bok := mm.Nums[math.B]
	if !aok || !bok {
		return -1, false // I don't have all of the proper values to solve this yet
	}

	var c int
	switch math.Op {
	case '+':
		c = a + b
	case '-':
		c = a - b
	case '*':
		c = a * b
	case '/':
		c = a / b
	default:
		fmt.Printf("Unknown operator: %v from Math %v\n", math.Op, math)
		return -2, false
	}

	mm.Nums[key] = c
	return c, true
}

func (mm *MonkeyMath) Solve() {
	for _, ok := mm.Nums[root]; !ok; _, ok = mm.Nums[root] {
		for m := range mm.Math {
			mm.DoMath(m)
		}
	}
}

func Parse(filename string) MonkeyMath {
	input, _ := util.ReadInput(filename)
	mathMap := make(map[string]Math)
	numMap := make(map[string]int)
	for _, line := range input {
		equation := strings.Split(line, ": ")
		fields := strings.Fields(equation[1])
		if len(fields) == 1 {
			n, _ := strconv.Atoi(fields[0])
			numMap[equation[0]] = n
		} else {
			mathMap[equation[0]] = Math{[]rune(fields[1])[0], fields[0], fields[2]}
		}
	}
	return MonkeyMath{mathMap, numMap}
}

var root string = "root"

func part1() {
	monkeyMath := Parse("input.txt")
	monkeyMath.Solve()
	fmt.Printf("Part 1: %v\n", monkeyMath.Nums[root])
}

func part2() {
	prototype := Parse("input.txt")
	max, min := 93813115694560, 0
	for true {
		humn := (max + min) / 2
		mm := prototype.Clone()
		mm.Nums["humn"] = humn
		mm.Solve()

		formula := mm.Math[root]
		a, b := mm.Nums[formula.A], mm.Nums[formula.B]
		if a == b {
			fmt.Printf("Part 2: %v\n", humn)
			break
		} else if a < b {
			max = humn - 1
		} else {
			min = humn + 1
		}
	}
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
