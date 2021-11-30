package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "09"
	dayTitle = "All in a Single Night"
)

type Node struct {
	src, dst string
}

func loadNodes(filename string) (map[Node]int, map[string]struct{}) {
	nodeMap := make(map[Node]int)
	nodes := make(map[string]struct{})
	for _, line := range util.MustRead(filename) {
		fields := strings.Fields(line)
		s, d, w := fields[0], fields[2], util.MustAtoi(fields[4])
		nodeMap[Node{s, d}] = w
		nodeMap[Node{d, s}] = w
		nodes[s] = struct{}{}
		nodes[d] = struct{}{}
	}
	return nodeMap, nodes
}

func search(edges map[Node]int, nodes []string,
	visited map[string]struct{}, prev string, acc int) (int, int) {
	if len(visited) == len(nodes) {
		return acc, acc
	}

	min, max := -1, -1
	for _, n := range nodes {
		if _, ok := visited[n]; ok {
			continue
		}

		visited[n] = struct{}{}
		lo, hi := search(edges, nodes, visited, n, acc+edges[Node{n, prev}])
		if min == -1 || lo < min {
			min = lo
		}
		if max == -1 || max < hi {
			max = hi
		}
		delete(visited, n)

	}
	return min, max
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	nodeMap, nodeSet := loadNodes("input.txt")
	nodes := make([]string, 0, len(nodeSet))
	for s := range nodeSet {
		nodes = append(nodes, s)
	}
	visited := make(map[string]struct{})
	min, max := search(nodeMap, nodes, visited, "", 0)
	fmt.Printf("Part 1: %v\n", min)
	fmt.Printf("Part 2: %v\n", max)
}
