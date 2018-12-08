package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	name string
	deps []Node
}

type Edge struct {
	to   string
	from string
}

func readInput(filename string) (map[string]*Node, []Edge) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return nil, nil
	}
	defer file.Close()

	// Step C must be finished before step A can begin.
	nameNode := make(map[string]*Node)
	var edges []Edge
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")

		s1 := words[1]
		n1, ok := nameNode[s1]
		if !ok {
			n1 = &Node{name: s1}
			nameNode[s1] = n1
		}

		s2 := words[7]
		n2, ok2 := nameNode[s2]
		if !ok2 {
			n2 = &Node{name: s2}
			nameNode[s2] = n2
		}

		n1.deps = append(n1.deps, *n2)

		e := Edge{s1, s2}
		edges = append(edges, e)

	}

	return nameNode, edges
}

func main() {
	fmt.Println("Day 07: The Sum of Its Parts")

	// part1()
	part2()
}

func part1() {
	fmt.Println("Part 1")

	nameNode, edges := readInput("input07.txt")

	// make map of deps to node
	depsMap := make(map[string][]string)
	for _, e := range edges {
		d, ok := depsMap[e.from]

		if !ok {
			d = []string{}
		}
		d = append(d, e.to)
		depsMap[e.from] = d
	}
	// fmt.Println(depsMap)

	// find the starting node
	deps := make(map[string]bool)
	for _, n := range nameNode {
		// fmt.Printf("name: %v node: %v\n", k, v)
		for _, d := range n.deps {
			deps[d.name] = true
		}
	}
	var startName []string
	for k := range nameNode {
		if _, ok := deps[k]; !ok {
			startName = append(startName, k)
		}
	}
	fmt.Printf("start name: %v\n", startName)

	queued := make(map[string]bool)
	queue := []Node{}
	for _, s := range startName {
		n := nameNode[s]
		queue = append(queue, *n)
		queued[s] = true
	}
	visited := make(map[string]bool)
	var solution []string
	for len(queue) != 0 {
		// find lowest in queue
		minName := "ZZZ"
		var minIdx int
		for i, n := range queue {

			prereqs := depsMap[n.name]
			ready := true
			for _, r := range prereqs {
				ready = ready && visited[r]
			}
			if !ready {
				continue
			}

			if n.name < minName {
				minName = n.name
				minIdx = i
			}
		}
		visited[minName] = true

		queue = append(queue[:minIdx], queue[minIdx+1:]...)
		for _, e := range nameNode[minName].deps {
			_, vok := visited[e.name]
			_, qok := queued[e.name]
			if !qok && !vok {
				queue = append(queue, e)
				queued[e.name] = true
			}
		}

		solution = append(solution, minName)
	}

	for _, s := range solution {
		fmt.Print(s)
	}
	fmt.Println()
}

const mod = 60
const numWorkers = 5

func part2() {
	fmt.Println("Part 2")

	nameNode, edges := readInput("input07.txt")

	// make map of deps to node
	depsMap := make(map[string][]string)
	for _, e := range edges {
		d, ok := depsMap[e.from]

		if !ok {
			d = []string{}
		}
		d = append(d, e.to)
		depsMap[e.from] = d
	}
	// fmt.Println(depsMap)

	// find the starting node
	deps := make(map[string]bool)
	for _, n := range nameNode {
		// fmt.Printf("name: %v node: %v\n", k, v)
		for _, d := range n.deps {
			deps[d.name] = true
		}
	}
	var startName []string
	for k := range nameNode {
		if _, ok := deps[k]; !ok {
			startName = append(startName, k)
		}
	}
	fmt.Printf("start name: %v\n", startName)

	// initialize starting state
	queued := make(map[string]bool)
	queue := []Node{}
	for _, s := range startName {
		n := nameNode[s]
		queue = append(queue, *n)
		queued[s] = true
	}
	visited := make(map[string]bool)

	workStep := [numWorkers]string{}
	workers := [numWorkers]int{}
	t := 0

	var solution []string
	for len(queue) != 0 || !idle(workers) {
		fmt.Println("---------------------------------------------------------------------------")
		fmt.Printf("time: %v\n", t)
		fmt.Printf("queue %v\n", qName(queue))
		fmt.Printf("visited: %v\n", visited)
		fmt.Printf("workStep %v workers %v\n", workStep, workers)
		fmt.Printf("solution: %v\n", solution)

		// pop from queue and push onto a free worker
		for len(queue) != 0 && atLeastIdle(workers) {

			// find next item to dequeue (lowest in queue with all prereqs complete)
			minName := "ZZZ"
			minIdx := -1
			for i, n := range queue {
				prereqs := depsMap[n.name]
				ready := true
				for _, r := range prereqs {
					ready = ready && visited[r]
				}
				if !ready {
					continue
				}

				if n.name < minName {
					minName = n.name
					minIdx = i
				}
			}

			// didn't find anything on the queue to work so advance to the next step
			if minName == "ZZZ" {
				break
			}

			// find a free worker to pick it up
			w := -1
			for i, v := range workers {
				if v == 0 {
					w = i
				}
			}
			// dequeue it and assign it to the free worker
			queue = append(queue[:minIdx], queue[minIdx+1:]...)
			workers[w] = steptime(minName, mod)
			workStep[w] = minName
		}

		if atLeastBusy(workers) {
			// find a worker who is working and work it (advance the clock)
			w := -1
			minTime := 60 * 60
			for i, v := range workers {
				if v < minTime && v != 0 {
					minTime = v
					w = i
				}
			}
			// sit idle until next step is ready
			if minTime != 0 {
				for i := range workers {
					if workers[i] != 0 {
						workers[i] -= minTime
					}
				}
				t += minTime
			}

			// if the worker has completed work, action it
			if workStep[w] != "" {
				solution = append(solution, workStep[w])
				visited[workStep[w]] = true

				for _, e := range nameNode[workStep[w]].deps {
					_, vok := visited[e.name]
					_, qok := queued[e.name]
					if !qok && !vok {
						queue = append(queue, e)
						queued[e.name] = true
					}
				}

				workStep[w] = ""
			}
		}
	}

	for _, s := range solution {
		fmt.Print(s)
	}
	fmt.Println()
	fmt.Println(t)
}

const alpha string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func qName(q []Node) []string {
	s := make([]string, len(q), len(q))
	for i, e := range q {
		s[i] = e.name
	}
	return s
}

func steptime(s string, mod int) int {
	for i := range alpha {
		if alpha[i] == s[0] {
			return i + mod + 1
		}
	}

	return -1
}

func atLeastBusy(w [numWorkers]int) bool {
	b := false
	for _, e := range w {
		b = b || e != 0
	}
	return b
}

func atLeastIdle(w [numWorkers]int) bool {
	b := false
	for _, e := range w {
		b = b || e == 0
	}
	return b
}

func idle(w [numWorkers]int) bool {
	sum := 0
	for _, e := range w {
		sum += e
	}
	return sum == 0
}
