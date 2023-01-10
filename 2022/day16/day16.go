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
}

func Search(graph Graph, curr string, state State) State {
	if state.Time <= 0 {
		return state
	}

	best := state
	for i, adj := range state.Available {
		time := state.Time - graph.MinDist[Edge{curr, adj}] - 1
		if time <= 0 {
			continue
		}
		score := state.Score + (time * graph.Nodes[adj].Flow)
		avail := make([]string, 0, len(state.Available)-1)
		avail = append(avail, state.Available[:i]...)
		avail = append(avail, state.Available[i+1:]...)
		newState := State{adj, avail, time, score}
		s := Search(graph, adj, newState)
		if best.Score < s.Score {
			best = s
		}
	}

	return best
}

func part1() {
	graph := parse("input.txt")

	start := "AA"
	allNodes := make([]string, 0, len(graph.Nodes))
	for name, node := range graph.Nodes {
		if node.Flow > 0 {
			allNodes = append(allNodes, name)
		}
	}

	startState := State{start, allNodes, 30, 0}
	state := Search(graph, start, startState)
	fmt.Printf("Part 1: %v\n", state.Score)
}

func part2() {
	fmt.Println("Part 2")
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
