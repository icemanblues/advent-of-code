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
	apply(string, int, RuleMap) (bool, int)
}

type ValueMatch struct {
	value string
}

func (vm ValueMatch) apply(msg string, i int, rules RuleMap) (bool, int) {
	if i >= len(msg) {
		return false, i
	}
	return vm.value == msg[i:i+1], i + 1
}

type Ordered struct {
	ruleNums []int
}

func (o Ordered) apply(msg string, i int, rules RuleMap) (bool, int) {
	/*
		if i >= len(msg) {
			return false, i
		}
	*/

	valid := true
	for _, ruleNum := range o.ruleNums {
		rule := rules[ruleNum]
		b, j := rule.apply(msg, i, rules)
		if !b {
			return false, j
		}
		i = j
		valid = valid && b
	}
	return valid, i
}

type Or struct {
	left, right Ordered
}

func (or Or) apply(msg string, i int, rules RuleMap) (bool, int) {
	/*
		if i >= len(msg) {
			return false, i
		}
	*/

	if b, j := or.left.apply(msg, i, rules); b {
		return b, j
	}

	return or.right.apply(msg, i, rules)
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
	for _, msg := range msgs {
		if ok, i := ruleMap[0].apply(msg, 0, ruleMap); ok {
			if i == len(msg) {
				fmt.Println(msg)
				count++
			}
		}
	}

	return count
}

func part1() {
	lines, _ := util.ReadInput("input.txt")
	rules, msgs := parse(lines)
	fmt.Printf("Part 1: %v\n", match(rules, msgs))
}

var test2Pass []string = []string{
	"bbabbbbaabaabba",
	"babbbbaabbbbbabbbbbbaabaaabaaa",
	"aaabbbbbbaaaabaababaabababbabaaabbababababaaa",
	"bbbbbbbaaaabbbbaaabbabaaa",
	"bbbababbbbaaaaaaaabbababaaababaabab",
	"ababaaaaaabaaab",
	"ababaaaaabbbaba",
	"baabbaaaabbaaaababbaababb",
	"abbbbabbbbaaaababbbbbbaaaababb",
	"aaaaabbaabaaaaababaa",
	"aaaabbaabbaaaaaaabbbabbbaaabbaabaaa",
	"aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba",
}

// 176 too low
// 352 too high
func part2() {
	lines, _ := util.ReadInput("test2.txt")
	rules, _ := parse(lines)

	parseRule("8: 42 | 42 8", rules)
	parseRule("11: 42 31 | 42 11 31", rules)

	fmt.Printf("rule 8: %v\n", rules[8])
	fmt.Printf("rule 11: %v\n", rules[11])
	fmt.Printf("rule 42: %v\n", rules[42])
	fmt.Printf("rule 31: %v\n", rules[31])

	count := 0
	for i, m := range test2Pass {
		b, j := rules[0].apply(m, 0, rules)
		fmt.Printf("%v: %v j=%v len=%v %v\n", i, b, j, len(m), m)
		if b {
			count++
		}
	}
	fmt.Println(count)

	for k, v := range rules {
		fmt.Printf("rule %v: %v\n", k, v)
	}

	//fmt.Printf("Part 2: %v\n", match(rules, msgs))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	//part1()
	part2()
}
