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

type graph map[string][]Edge

func buildRules(lines []string) ([]Edge, []string) {
	bags := make(map[string]struct{})
	var roots []string
	var edges []Edge

	for _, line := range lines {
		line = strings.ReplaceAll(line, ".", "")
		outerInner := strings.Split(line, " contain ")
		outerBag := strings.TrimSpace(outerInner[0])

		// add it
		bags[outerBag] = struct{}{}

		if outerInner[1] == "no other bags" {
			roots = append(roots, outerBag)
			continue
		}

		innerBags := strings.Split(outerInner[1], ", ")
		for _, b := range innerBags {
			parts := strings.Split(b, " ")
			num, _ := strconv.Atoi(parts[0])
			name := parts[1] + " " + parts[2] + " bags"

			bags[name] = struct{}{}
			edges = append(edges, Edge{outerBag, name, num})
		}
	}
	return edges, roots
}

func buildInnerToOuterMap(edges []Edge) map[string][]Edge {
	innerToOuter := make(map[string][]Edge)
	for _, edge := range edges {
		inEdges := innerToOuter[edge.inner]
		inEdges = append(inEdges, edge)
		innerToOuter[edge.inner] = inEdges
	}
	return innerToOuter
}

func buildOuterToInnerMap(edges []Edge) map[string][]Edge {
	innerToOuter := make(map[string][]Edge)
	for _, edge := range edges {
		inEdges := innerToOuter[edge.outer]
		inEdges = append(inEdges, edge)
		innerToOuter[edge.outer] = inEdges
	}
	return innerToOuter
}

func part1() {
	lines, _ := util.ReadInput("input.txt")
	edges, _ := buildRules(lines)
	inner2outer := buildInnerToOuterMap(edges)

	target := "shiny gold bags"
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
	fmt.Printf("Part 1: %v\n", len(set))
}

//12206 too high
func part2() {
	lines, _ := util.ReadInput("input.txt")
	edges, roots := buildRules(lines)
	bagCounts := make(map[string]int)
	for _, root := range roots {
		bagCounts[root] = 0
	}

	target := "shiny gold bags"
	outer2inner := buildOuterToInnerMap(edges)

	for _, ok := bagCounts[target]; !ok; _, ok = bagCounts[target] {
	label:
		for _, edge := range edges {
			if _, ok := bagCounts[edge.outer]; ok {
				continue
			}
			inners := outer2inner[edge.outer]
			sum := 0
			for _, inEdge := range inners {
				count, ok := bagCounts[inEdge.inner]
				if !ok {
					continue label
				}
				w := inEdge.weight + (inEdge.weight * count)
				sum = sum + w
			}
			bagCounts[edge.outer] = sum
		}
	}

	fmt.Printf("Part 2: %v\n", bagCounts[target])
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
