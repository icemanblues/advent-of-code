package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "19"
	dayTitle = "Monster Messages"
)

type RuleMap map[int]Rule

type Rule interface {
	// apply int -> int (in bulk). if the current int passes, and what the next idx is
	apply(RuleMap, string, []int) []int
}

type ValueMatch struct {
	value string
}

func (vm ValueMatch) apply(rules RuleMap, msg string, indexes []int) []int {
	var idx []int
	for _, i := range indexes {
		if i < len(msg) && vm.value == msg[i:i+1] {
			idx = append(idx, i+1)
		}
	}
	return idx
}

type Ordered struct {
	ruleNums []int
}

func (o Ordered) apply(rules RuleMap, msg string, indexes []int) []int {
	var idx []int

	for _, i := range indexes {
		search := []int{i}
		for _, ruleNum := range o.ruleNums {
			rule := rules[ruleNum]
			bidx := rule.apply(rules, msg, search)
			search = nil
			for _, b := range bidx {
				search = append(search, b)
			}
		}
		// we finished running all of the rules
		for _, sidx := range search {
			idx = append(idx, sidx)
		}
	}
	return idx
}

type Or struct {
	left, right Ordered
}

func (or Or) apply(rules RuleMap, msg string, indexes []int) []int {
	var idx []int

	lidx := or.left.apply(rules, msg, indexes)
	idx = append(idx, lidx...)

	ridx := or.right.apply(rules, msg, indexes)
	idx = append(idx, ridx...)

	return idx
}

func strings2ints(rulers []string) []int {
	ruleNums := make([]int, 0, len(rulers))
	for _, s := range rulers {
		r, _ := strconv.Atoi(s)
		ruleNums = append(ruleNums, r)
	}
	return ruleNums
}

func parseRule(line string, ruleMap RuleMap) {
	// parse rule and add it
	parts := strings.Split(line, ": ")
	num, _ := strconv.Atoi(parts[0])
	// check value
	if parts[1][1:2] == "a" || parts[1][1:2] == "b" {
		ruleMap[num] = ValueMatch{parts[1][1:2]}
		return
	}

	// check or or ordered
	extender := strings.Split(parts[1], " | ")
	if len(extender) == 1 { // Ordered
		rulers := strings.Split(extender[0], " ")
		ruleNums := strings2ints(rulers)
		ruleMap[num] = Ordered{ruleNums}
		return
	}

	rulers := strings.Split(extender[0], " ")
	leftRuleNums := strings2ints(rulers)
	rulers = strings.Split(extender[1], " ")
	rightRuleNums := strings2ints(rulers)
	ruleMap[num] = Or{Ordered{leftRuleNums}, Ordered{rightRuleNums}}
	return
}

func parse(lines []string) (RuleMap, []string) {
	ruleMap := make(RuleMap)
	var msgs []string

	isRule := true
	for _, line := range lines {
		if len(line) == 0 {
			isRule = false
			continue
		}

		if isRule {
			parseRule(line, ruleMap)
			continue
		}

		msgs = append(msgs, line)
	}

	return ruleMap, msgs
}

func match(ruleMap RuleMap, msgs []string) int {
	count := 0
	start := []int{0}
	for _, msg := range msgs {
		idx := ruleMap[0].apply(ruleMap, msg, start)
		for _, j := range idx {
			if j == len(msg) {
				count++
			}
		}
	}

	return count
}

func updateRules(rules RuleMap) {
	parseRule("8: 42 | 42 8", rules)
	parseRule("11: 42 31 | 42 11 31", rules)
}

func part1() {
	lines, _ := util.ReadInput("input.txt")
	rules, msgs := parse(lines)
	fmt.Printf("Part 1: %v\n", match(rules, msgs))
}

func part2() {
	lines, _ := util.ReadInput("input.txt")
	rules, msgs := parse(lines)
	updateRules(rules)
	fmt.Printf("Part 2: %v\n", match(rules, msgs))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
