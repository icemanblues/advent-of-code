package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "12"
	dayTitle = "Passage Pathing"
)

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func Parse(filename string) (map[string][]string, map[string]struct{}) {
	input, _ := util.ReadInput(filename)
	adj := make(map[string][]string)
	small := make(map[string]struct{})
	for _, line := range input {
		parts := strings.Split(line, "-")
		if !IsUpper(parts[0]) {
			small[parts[0]] = struct{}{}
		}
		if !IsUpper(parts[1]) {
			small[parts[1]] = struct{}{}
		}

		adj[parts[0]] = append(adj[parts[0]], parts[1])
		adj[parts[1]] = append(adj[parts[1]], parts[0])
	}

	delete(small, "start")
	delete(small, "end")
	return adj, small
}

func search(n string, path map[string]int, adj map[string][]string, small map[string]struct{}) int {
	if n == "end" {
		return 1
	}
	if _, isSmall := small[n]; isSmall && path[n] >= 2 {
		return 0
	}

	path[n]++
	sum := 0
	for _, child := range adj[n] {
		if child == "start" {
			continue
		}

		path[child]++
		sum += search(child, path, adj, small)
		path[child]--
	}
	path[n]--

	return sum
}

func searchTwo(n string, path map[string]int, visitedSmall string, adj map[string][]string, small map[string]struct{}) int {
	if n == "end" {
		return 1
	}

	_, isSmall := small[n]
	if visitedSmall == "" && isSmall && path[n] == 1 {
		visitedSmall = n
	}
	if visitedSmall != n && isSmall && path[n] >= 1 {
		return 0
	}
	if visitedSmall == n && isSmall && path[n] >= 2 {
		return 0
	}

	path[n]++
	sum := 0
	for _, child := range adj[n] {
		if child == "start" {
			continue
		}

		sum += searchTwo(child, path, visitedSmall, adj, small)
	}
	path[n]--

	return sum
}

func part1() {
	adj, small := Parse("input.txt")
	path := make(map[string]int)
	fmt.Printf("Part 1: %v\n", search("start", path, adj, small))
}

func part2() {
	adj, small := Parse("input.txt")
	path := make(map[string]int)
	fmt.Printf("Part 2: %v\n", searchTwo("start", path, "", adj, small))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
