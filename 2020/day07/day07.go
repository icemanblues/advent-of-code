package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "07"
	dayTitle = "Handy Haversacks"
)

type Edge struct {
	outer  string
	inner  string
	weight int
}

// buildRules given the input file, converts it to a list of edges and bags that contain nothing
func buildRules(lines []string) ([]Edge, []string) {
	var edges []Edge
	var roots []string

	for _, line := range lines {
		line = strings.ReplaceAll(line, ".", "")
		outerInner := strings.Split(line, " contain ")
		outerBag := strings.TrimSpace(outerInner[0])

		if outerInner[1] == "no other bags" {
			roots = append(roots, outerBag)
			continue
		}

		innerBags := strings.Split(outerInner[1], ", ")
		for _, b := range innerBags {
			parts := strings.Split(b, " ")
			num, _ := strconv.Atoi(parts[0])
			name := parts[1] + " " + parts[2] + " bags"
			edges = append(edges, Edge{outerBag, name, num})
		}
	}
	return edges, roots
}

// buildInnerToOuterMap maps the bags (by their name) to their edges (where they are the inner bag)
func buildInnerToOuterMap(edges []Edge) map[string][]Edge {
	innerToOuter := make(map[string][]Edge)
	for _, edge := range edges {
		inEdges := innerToOuter[edge.inner]
		inEdges = append(inEdges, edge)
		innerToOuter[edge.inner] = inEdges
	}
	return innerToOuter
}

// buildOuterToInnerMap maps the bags (by their name) to their edges (where they are the outer bag)
func buildOuterToInnerMap(edges []Edge) map[string][]Edge {
	outerToInner := make(map[string][]Edge)
	for _, edge := range edges {
		outEdges := outerToInner[edge.outer]
		outEdges = append(outEdges, edge)
		outerToInner[edge.outer] = outEdges
	}
	return outerToInner
}

// bagCountThatContain returns a count of bag types that could contain the target bag
func bagCountThatContain(edges []Edge, target string) int {
	inner2outer := buildInnerToOuterMap(edges)

	set := make(map[string]struct{})
	queue := []string{target}
	for len(queue) != 0 {
		b := queue[0]
		queue = queue[1:]

		edges := inner2outer[b]
		for _, e := range edges {
			set[e.outer] = struct{}{}
			queue = append(queue, e.outer)
		}
	}

	return len(set)
}

// bagCountAllWithin a count of all individual bags within the target bag
func bagCountAllWithin(edges []Edge, roots []string, target string) int {
	outer2inner := buildOuterToInnerMap(edges)
	bagCounts := make(map[string]int)
	for _, root := range roots {
		bagCounts[root] = 0
	}

	for _, ok := bagCounts[target]; !ok; _, ok = bagCounts[target] {
	edgeSearch:
		for _, edge := range edges {
			if _, ok := bagCounts[edge.outer]; ok {
				continue
			}
			inners := outer2inner[edge.outer]
			sum := 0
			for _, inEdge := range inners {
				count, ok := bagCounts[inEdge.inner]
				if !ok {
					continue edgeSearch
				}
				w := inEdge.weight + (inEdge.weight * count)
				sum = sum + w
			}
			bagCounts[edge.outer] = sum
		}
	}

	return bagCounts[target]
}

func part1() {
	lines, _ := util.ReadInput("input.txt")
	edges, _ := buildRules(lines)
	fmt.Printf("Part 1: %v\n", bagCountThatContain(edges, "shiny gold bags"))
}

func part2() {
	lines, _ := util.ReadInput("input.txt")
	edges, roots := buildRules(lines)
	fmt.Printf("Part 2: %v\n", bagCountAllWithin(edges, roots, "shiny gold bags"))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
