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

func score(polymer runecount) int64 {
	var min, max int64 = -1, -1
	for _, v := range polymer {
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
	counts := make(runecount)
	for _, r := range start {
		counts[r]++
	}

	fmt.Printf("Part 1: %v\n", score(counts))
}

type runecount map[rune]int64

func elements(polymer string, counts runecount) runecount {
	for _, r := range polymer {
		counts[r]++
	}
	return counts
}

type Reaction struct {
	in     string
	out    []string
	insert rune
}

func rulesToReaction(rule Rule) Reaction {
	return Reaction{
		rule.reagents,
		[]string{
			rule.reagents[:1] + rule.results,
			rule.results + rule.reagents[1:],
		},
		[]rune(rule.results)[0],
	}
}

func part2() {
	start, rules := Parse("input.txt")

	elementCounts := make(runecount)
	elements(start, elementCounts)

	reactions := make([]Reaction, 0, len(rules))
	for _, rule := range rules {
		reactions = append(reactions, rulesToReaction(rule))
	}

	pairCounts := make(map[string]int64)
	for i := 0; i < len(start)-1; i++ {
		pairCounts[start[i:i+2]]++
	}

	// run the number of steps
	for i := 0; i < 40; i++ {
		newPairs := make(map[string]int64)
		for pair, count := range pairCounts {
			// find the corresponding reaction // this should be a map lookup. ugh!
			var reaction Reaction
			for _, re := range reactions {
				if re.in == pair {
					reaction = re
				}
			}

			// add the new pairs and rune
			for _, p := range reaction.out {
				newPairs[p] += count
			}
			elementCounts[reaction.insert] += count
		}
		//iterate
		pairCounts = newPairs
	}

	fmt.Printf("Part 2: %v\n", score(elementCounts))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
