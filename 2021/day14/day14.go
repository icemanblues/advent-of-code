package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "14"
	dayTitle = "Extended Polymerization"
)

type Rule struct {
	reagents string
	results  string
}

func Parse(filename string) (string, []Rule) {
	input, _ := util.ReadInput(filename)
	start := input[0]

	var rules []Rule
	for _, line := range input[2:] {
		parts := strings.Split(line, " -> ")
		r := Rule{parts[0], parts[1]}
		rules = append(rules, r)
	}
	return start, rules
}

func step(start string, rules []Rule) string {
	builder := strings.Builder{}
	for i := 0; i < len(start)-1; i++ {
		pair := start[i : i+2]
		for _, rule := range rules { // can make this a hashmap for better lookup
			if rule.reagents == pair {
				builder.WriteString(pair[:1])
				builder.WriteString(rule.results)
				break
			}
		}
	}
	builder.WriteString(start[len(start)-1:])

	return builder.String()
}

func score(polymer string) int {
	counts := make(map[rune]int)
	for _, r := range polymer {
		counts[r]++
	}
	min, max := -1, -1
	for _, v := range counts {
		if min == -1 {
			min = v
		}
		if v < min {
			min = v
		}

		if max == -1 {
			max = v
		}
		if v > max {
			max = v
		}
	}

	return max - min
}

func part1() {
	start, rules := Parse("input.txt")
	for i := 0; i < 10; i++ {
		start = step(start, rules)
	}
	fmt.Printf("Part1: %v\n", score(start))
}

func part2() {}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
