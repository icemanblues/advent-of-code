package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "25"
	dayTitle = "Snowverload"
)

func contains(wires [][2]string, wire [2]string) bool {
	for _, w := range wires {
		if wire == w {
			return true
		}
	}
	return false
}

func mapClone(m map[string][]string, cut [][2]string) map[string][]string {
	clone := make(map[string][]string)
	for k, v := range m {
		for _, vv := range v {
			c := [2]string{k, vv}
			if !contains(cut, c) {
				clone[k] = append(clone[k], vv)
			}
		}
	}
	return clone
}

func groupCount(m map[string][]string) []map[string]struct{} {
	sets := make([]map[string]struct{}, 0)

	for k, v := range m {
		// get the existing set (if none, create it)
		set := make(map[string]struct{})
		found := false
		for _, s := range sets {
			if _, ok := s[k]; ok {
				set = s
				found = true
			}
		}
		if !found {
			sets = append(sets, set)
		}
		set[k] = struct{}{}
		for _, vv := range v {
			set[vv] = struct{}{}
		}
	}
	return sets
}

func makeSets(m map[string][]string) []map[string]struct{} {
	sets := make([]map[string]struct{}, 0)
outer:
	for k := range m {
		if len(sets) > 2 { // specific to our problem space
			return sets
		}

		// if k exists in a set. If it does, continue
		for _, set := range sets {
			if _, ok := set[k]; ok {
				continue outer
			}
		}

		// not in a set, so create one and populate it
		set := make(map[string]struct{})
		queue := []string{k}
		for len(queue) != 0 {
			curr := queue[0]
			queue = queue[1:]

			if _, ok := set[curr]; ok { // we've already visited
				continue
			}

			set[curr] = struct{}{}
			for _, next := range m[curr] {
				queue = append(queue, next)
			}
		}
		sets = append(sets, set)
	}
	return sets
}

func temp(adjList map[string][]string) {
	try := mapClone(adjList, [][2]string{
		{"hfx", "pzl"}, {"pzl", "hfx"},
		{"bvb", "cmg"}, {"cmg", "bvb"},
		{"nvd", "jqt"}, {"jqt", "nvd"},
	})
	//gg := groupCount(try)
	gg := makeSets(try)
	prod := 1
	for i, e := range gg {
		fmt.Printf("Set %v: %v\n", i, e)
		prod *= len(e)
	}
	fmt.Printf("Part 1: %v\n", prod)
	return

}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	input, _ := util.ReadInput("input.txt")
	adjList := make(map[string][]string)
	for _, line := range input {
		split := strings.Split(line, ": ")
		comps := strings.Fields(split[1])
		for _, c := range comps {
			adjList[split[0]] = append(adjList[split[0]], c)
			adjList[c] = append(adjList[c], split[0])
		}
	}

	// generate all possible wires that we can cut
	wires := make([][2]string, 0)
	for k, v := range adjList {
		for _, w := range v {
			wires = append(wires, [2]string{k, w})
		}
	}
	fmt.Printf("number of wires: %v\n", len(wires))
	sort.Slice(wires, func(i, j int) bool {
		iLen := len(adjList[wires[i][0]]) + len(adjList[wires[i][1]])
		jLen := len(adjList[wires[j][0]]) + len(adjList[wires[j][1]])
		return iLen > jLen
	})

	// search for the 3 to cut
	var g []map[string]struct{}
search:
	for i := 0; i < len(wires)-2; i++ {
		for j := i + 1; j < len(wires)-1; j++ {
			for k := j + 1; k < len(wires); k++ {
				cut := [][2]string{
					wires[i], {wires[i][1], wires[i][0]},
					wires[j], {wires[j][1], wires[j][0]},
					wires[k], {wires[k][1], wires[k][0]},
				}

				// check that the wires to be cut are valid (no dupes)
				checker := make(map[[2]string]struct{})
				for _, c := range cut {
					checker[c] = struct{}{}
				}
				if len(checker) != 6 {
					continue
				}

				// copy the map with removed wires
				clone := mapClone(adjList, cut)

				// count the number of connected groups
				//g = groupCount(clone)
				g = makeSets(clone)
				if len(g) == 2 {
					break search
				}
			}
		}
	}

	// count it up
	fmt.Println(g[0])
	fmt.Println(g[1])
	sum := len(g[0]) * len(g[1])
	fmt.Printf("Part 1: %v\n", sum)
}
