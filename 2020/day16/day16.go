package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "16"
	dayTitle = "Ticket Translation"
)

type mode int

const (
	field mode = iota
	your
	nearby
)

type Rule struct {
	field                string
	firstMin, firstMax   int
	secondMin, secondMax int
}

func (r Rule) apply(n int) bool {
	return (n >= r.firstMin && n <= r.firstMax) || (n >= r.secondMin && n <= r.secondMax)
}

func parseRules(lines []string) ([]Rule, [][]int, []int) {
	var rules []Rule
	var tickets [][]int
	var ur []int

	mode := field
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if line == "your ticket:" {
			mode = your
			continue
		}
		if line == "nearby tickets:" {
			mode = nearby
			continue
		}

		switch mode {
		case field:
			// build a struct and add it to a list
			rule := Rule{}
			fieldRanges := strings.Split(line, ": ")
			fieldName := fieldRanges[0]
			rule.field = fieldName
			for i, ranges := range strings.Split(fieldRanges[1], " or ") {
				loHi := strings.Split(ranges, "-")
				lo, _ := strconv.Atoi(loHi[0])
				hi, _ := strconv.Atoi(loHi[1])
				if i == 0 {
					rule.firstMin = lo
					rule.firstMax = hi
				} else {
					rule.secondMin = lo
					rule.secondMax = hi
				}
			}
			rules = append(rules, rule)
		case your:
			words := strings.Split(line, ",")
			urTicket := make([]int, 0, len(words))
			for _, word := range words {
				num, _ := strconv.Atoi(word)
				urTicket = append(urTicket, num)
			}
			ur = urTicket
		case nearby:
			words := strings.Split(line, ",")
			nearbyTicket := make([]int, 0, len(words))
			for _, word := range words {
				num, _ := strconv.Atoi(word)
				nearbyTicket = append(nearbyTicket, num)
			}
			tickets = append(tickets, nearbyTicket)
		}
	}

	return rules, tickets, ur
}

func part1() {
	lines, _ := util.ReadInput("input.txt")
	rules, nearby, _ := parseRules(lines)

	sum := 0
	for _, t := range nearby {
		for _, n := range t {
			invalidCount := 0
			for _, r := range rules {
				if !r.apply(n) {
					invalidCount++
				}
			}
			if invalidCount == len(rules) {
				sum += n
			}
		}
	}

	fmt.Printf("Part 1: %v\n", sum)
}

func backTracker(rules []Rule, tickets [][]int, sol []Rule) []Rule {
	if len(rules) == 0 {
		return sol
	}
	if len(sol) == len(tickets[0]) {
		return sol
	}

	//	for i := len(sol); i < len(tickets[0]); i++ {
	i := len(sol)
ruleSearch:
	for r, rule := range rules {
		//fmt.Printf("trying.. %v %v %v\n", i, rule, rules)
		for _, t := range tickets {
			if !rule.apply(t[i]) {
				//fmt.Printf("bad rule %v not good for ticket %v\n", rule, i)
				continue ruleSearch
			}
		}

		// valid. add it to the solution list and continue
		//fmt.Println(i)
		//fmt.Printf("POSSIBILITY! i %v rule: %v sol %v\n", i, rule, sol)
		//fmt.Printf("pos rules: %v\n", rules)
		//fmt.Printf("POSSIBILITY! i %v sol: %v\n", i, sol)

		//sol := append(sol, rule)
		itersol := make([]Rule, 0, len(sol)+1)
		itersol = append(itersol, sol...)
		itersol = append(itersol, rule)

		iterRules := make([]Rule, 0, len(rules)-1)
		iterRules = append(iterRules, rules[:r]...)
		iterRules = append(iterRules, rules[r+1:]...)
		//fmt.Printf("add r %v and remove r %v\n", rule, iterRules)
		//fmt.Printf("rules: %v\n", rules)
		answer := backTracker(iterRules, tickets, itersol)
		if answer != nil {
			return answer
		}
		// otherwise its bad so remove what we choose
		//sol = sol[:len(sol)-1]
		//fmt.Printf("nope! i %v sol: %v\n", i, sol)
		//fmt.Printf("nope! no good. try the next rule. answer: %v sol: %v r: %v rule %v lenRule: %v\n",
		//answer, sol, r, rule, len(rules))
	}
	//return nil
	// println(sol)
	// panic("none of the rules work")
	//}

	return nil
}

func setReduction(rules []Rule, tickets [][]int) []Rule {
	possibilities := make(map[int]map[Rule]struct{})

	for i := 0; i < len(tickets[0]); i++ {
	ruleLoop:
		for _, rule := range rules {
			for _, t := range tickets {
				if !rule.apply(t[i]) {
					continue ruleLoop
				}
			}
			// rule is good for i
			m, ok := possibilities[i]
			if !ok {
				m = make(map[Rule]struct{})
				possibilities[i] = m
			}
			possibilities[i][rule] = struct{}{}
		}
	}
	//fmt.Println(possibilities)

	solo := []int{}
	for k, v := range possibilities {
		if len(v) == 1 {
			solo = append(solo, k)
			//delete(possibilities, k)
		}
	}

	sol := make([]Rule, len(rules), len(rules))
	//	for len(possibilities) != 0 {
	for len(solo) != 0 {
		/*
			fmt.Printf("sol: %v \n", sol)
			fmt.Printf("solo: %v %v\n", len(solo), solo)
			fmt.Printf("poss: %v\n", len(possibilities))
			fmt.Println(possibilities)
		*/

		for _, n := range solo {
			//fmt.Println("thinking of deleting something")
			m := possibilities[n]
			//fmt.Printf("anything to delete? %v %v\n", n, m)
			for k := range m { // should only be one
				// add k to solution
				sol[n] = k
				// and remove k from all possibilies
				for x := range possibilities {
					//fmt.Printf("removing %v %v\n", x, k)
					delete(possibilities[x], k)
				}
			}
		}

		// find the next ones to process
		solo = nil
		for k, v := range possibilities {
			if len(v) == 1 {
				solo = append(solo, k)
				//delete(possibilities, k)
			}
		}
	}

	return sol
}

func part2() {
	//lines, _ := util.ReadInput("test.txt")
	//lines, _ := util.ReadInput("test2.txt")
	lines, _ := util.ReadInput("input.txt")
	rules, nearby, ur := parseRules(lines)

	sum := 0
	var goodTickets [][]int
ticketSearch:
	for _, t := range nearby {
		for _, n := range t {
			validCount := 0
			for _, r := range rules {
				if r.apply(n) {
					validCount++
				}
			}
			if validCount == 0 {
				sum += n
				continue ticketSearch
			}
		}
		goodTickets = append(goodTickets, t)
	}
	fmt.Printf("removed value sum: %v\n", sum)
	fmt.Printf("lenGood: %v lenTotal: %v\n", len(goodTickets), len(nearby))
	//fmt.Println(goodTickets)

	// backtracking... ugh!
	// solution := make([]Rule, 0)
	// fieldMapper := backTracker(rules, goodTickets, solution)
	fieldMapper := setReduction(rules, goodTickets)
	//fmt.Println(fieldMapper)

	prod := 1
	for i, rule := range fieldMapper {
		if strings.HasPrefix(rule.field, "departure") {
			fmt.Println(rule.field)
			prod *= ur[i]
		}
	}

	fmt.Printf("Part 2: %v\n", prod)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
