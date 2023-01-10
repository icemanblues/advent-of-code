package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "16"
	dayTitle = "Proboscidea Volcanium"
)

type Node struct {
	Name string
	Flow int
	Adj  []string
}

type Edge struct {
	S, E string
}

type Graph struct {
	Nodes   map[string]Node
	Edges   []Edge
	MinDist map[Edge]int
}

func parse(filename string) Graph {
	input, _ := util.ReadInput(filename)
	nodes := make(map[string]Node)
	for _, line := range input {
		fields := strings.Fields(line)
		valveName := fields[1]

		as := strings.Split(line, ";")[0]
		bs := strings.Split(as, "=")[1]
		flowRate, _ := strconv.Atoi(bs)

		adjValves := fields[9:]
		adj := make([]string, 0, len(adjValves))
		for _, a := range adjValves {
			b := strings.ReplaceAll(a, ",", "")
			adj = append(adj, b)
		}
		nodes[valveName] = Node{valveName, flowRate, adj}
	}

	// compute node dists (shortest paths)
	minDist := make(map[Edge]int) // here edge means path

	// all edges (directed)
	var edges []Edge
	for _, node := range nodes {
		for _, adj := range node.Adj {
			edges = append(edges, Edge{node.Name, adj})
		}
	}

	// Floyd Warshall Algo
	for _, edge := range edges {
		minDist[edge] = 1
	}
	for _, node := range nodes {
		minDist[Edge{node.Name, node.Name}] = 0
	}
	for _, k := range nodes {
		for _, i := range nodes {
			for _, j := range nodes {
				eij := Edge{i.Name, j.Name}
				eik := Edge{i.Name, k.Name}
				ekj := Edge{k.Name, j.Name}
				dij, ijok := minDist[eij]
				dik, ikok := minDist[eik]
				dkj, kjok := minDist[ekj]
				if !ijok && ikok && kjok {
					minDist[eij] = dik + dkj
				}
				if ikok && kjok && dij > dik+dkj {
					minDist[eij] = dik + dkj
				}
			}
		}
	}

	return Graph{nodes, edges, minDist}
}

type State struct {
	Current   string
	Available []string
	Time      int
	Score     int
	Path      []string
}

var final []State

func Search(graph Graph, state State) State {
	final = append(final, state)
	if state.Time <= 0 {
		return state
	}

	best := state
	for i, adj := range state.Available {
		time := state.Time - graph.MinDist[Edge{state.Current, adj}] - 1
		if time <= 0 {
			continue
		}
		score := state.Score + (time * graph.Nodes[adj].Flow)
		avail := make([]string, 0, len(state.Available)-1)
		avail = append(avail, state.Available[:i]...)
		avail = append(avail, state.Available[i+1:]...)
		path := make([]string, 0, len(state.Path)+1)
		path = append(path, state.Path...)
		path = append(path, adj)
		newState := State{adj, avail, time, score, path}
		s := Search(graph, newState)
		final = append(final, s)
		if best.Score < s.Score {
			best = s
		}
	}

	return best
}

func part1() {
	graph := parse("test.txt")

	start := "AA"
	allNodes := make([]string, 0, len(graph.Nodes))
	for name, node := range graph.Nodes {
		if node.Flow > 0 {
			allNodes = append(allNodes, name)
		}
	}

	startState := State{start, allNodes, 30, 0, nil}
	state := Search(graph, startState)
	fmt.Printf("Part 1: %v\n", state.Score)
}

func part2() {
	graph := parse("input.txt")

	start := "AA"
	allNodes := make([]string, 0, len(graph.Nodes))
	for name, node := range graph.Nodes {
		if node.Flow > 0 {
			allNodes = append(allNodes, name)
		}
	}

	final = make([]State, 0)
	startState := State{start, allNodes, 26, 0, nil}
	Search(graph, startState)

	best := 0
	for i := 0; i < len(final)-1; i++ {
	iter:
		for j := 0; j < len(final); j++ {
			s1, s2 := final[i], final[j]
			set := make(map[string]struct{})
			for _, n := range s1.Path {
				set[n] = struct{}{}
			}
			for _, n := range s2.Path {
				if _, ok := set[n]; ok {
					continue iter
				}
			}
			if s := s1.Score + s2.Score; s > best {
				best = s
			}
		}
	}

	fmt.Printf("Part 2: %v\n", best)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
