package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "11"
	dayTitle = "Monkey in the Middle"
)

type Monkey struct {
	Items       []int
	Worry_Op    string
	Worry_Value string
	D, T, F     int
}

func parse(filename string) []*Monkey {
	input, _ := util.ReadInput(filename)
	var monkeys []*Monkey
	var m *Monkey
	for _, line := range input {
		if line == "" {
			monkeys = append(monkeys, m)
			continue
		}

		fields := strings.Fields(line)
		switch fields[0] {
		case "Monkey":
			m = &Monkey{nil, "", "", 0, 0, 0}
		case "Starting":
			numbers := strings.Split(line, ": ")
			nums := strings.Split(numbers[1], ", ")
			for _, n := range nums {
				i, _ := strconv.Atoi(n)
				m.Items = append(m.Items, i)
			}
		case "Operation:":
			m.Worry_Op = fields[4]
			m.Worry_Value = fields[5]
		case "Test:":
			m.D, _ = strconv.Atoi(fields[3])
		case "If":
			if fields[1] == "true:" {
				m.T, _ = strconv.Atoi(fields[5])
			} else {
				m.F, _ = strconv.Atoi(fields[5])
			}
		}
	}
	monkeys = append(monkeys, m)
	return monkeys
}

func worry(op, value string, old int) int {
	i, err := strconv.Atoi(value)
	if err != nil {
		i = old
	}

	switch op {
	case "+":
		return old + i
	case "*":
		return old * i
	default:
		fmt.Printf("Unknown operator: %v\n", op)
		return old
	}
}

func part1() {
	monkeys := parse("input.txt")
	mCount := make([]int, len(monkeys), len(monkeys))
	for round := 0; round < 20; round++ {
		for i := range monkeys {
			m := monkeys[i]
			for len(m.Items) > 0 {
				item := m.Items[0]
				m.Items = m.Items[1:]

				new := worry(m.Worry_Op, m.Worry_Value, item) / 3
				mCount[i]++

				target := monkeys[m.F]
				if new%m.D == 0 {
					target = monkeys[m.T]
				}
				target.Items = append(target.Items, new)
			}
		}
	}
	sort.Ints(mCount)
	fmt.Printf("Part 1: %v\n", mCount[len(mCount)-1]*mCount[len(mCount)-2])
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
}
