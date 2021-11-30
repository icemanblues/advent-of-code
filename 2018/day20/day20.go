package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

// part1 -> 251 too low
func main() {
	fmt.Println("Day 20: A Regular Map")

	// test1 := "^WNE$"                                                             // 3
	// test2 := "^ENWWW(NEEE|SSE(EE|N))$"                                           // 10
	// test3 := "^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$"                         // 18, 30 nodes
	// test4 := "^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$"               // 23
	// test5 := "^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$" // 31
	input := readInput("input20.txt")

	// part1(test1)
	// part1(test2)
	// part1(test3)
	// part1(test4)
	// part1(test5)
	part1(input)

	part2(input)
}

type Point struct {
	X int
	Y int
}
type Node struct {
	ID    int
	loc   Point
	North *Node
	East  *Node
	South *Node
	West  *Node
}

var empty Node = Node{}

func NewNode(p Point, i int) *Node {
	return &Node{
		ID:    i,
		loc:   p,
		North: &empty,
		South: &empty,
		West:  &empty,
		East:  &empty,
	}
}
func furthest(regex string, i int, steps int) int {
	if i >= len(regex) {
		return steps
	}

	switch r := rune(regex[i]); r {
	case '^':
		return furthest(regex, i+1, steps)
	case '$':
		return steps
	case 'N':
		return furthest(regex, i+1, steps+1)
	case 'W':
		return furthest(regex, i+1, steps+1)
	case 'E':
		return furthest(regex, i+1, steps+1)
	case 'S':
		return furthest(regex, i+1, steps+1)

	// handle branches
	case '(':
		return furthest(regex, i+1, steps)
	case ')':
		return furthest(regex, i+1, steps)
	case '|':
		return furthest(regex, i+1, steps)
	}

	fmt.Println("!!! ERROR !!! never should have reached here ")
	return steps
}

func buildGraph(regex string) (*Node, map[Point]*Node) {
	start := NewNode(Point{0, 0}, 0)
	curr := start

	qi := 0
	stack := []*Node{start}

	nodes := make(map[Point]*Node)

	// build the graph
	for i, r := range regex {
		switch r {
		case '^':
		case '$':
		case 'N':
			p := Point{X: curr.loc.X, Y: curr.loc.Y + 1}
			n, ok := nodes[p]
			if !ok {
				n = NewNode(p, i)
				nodes[p] = n
			}
			curr.North = n
			n.South = curr
			curr = curr.North
		case 'S':
			p := Point{X: curr.loc.X, Y: curr.loc.Y - 1}
			n, ok := nodes[p]
			if !ok {
				n = NewNode(p, i)
				nodes[p] = n
			}
			curr.South = n
			n.North = curr
			curr = curr.South
		case 'E':
			p := Point{X: curr.loc.X + 1, Y: curr.loc.Y}
			n, ok := nodes[p]
			if !ok {
				n = NewNode(p, i)
				nodes[p] = n
			}
			curr.East = n
			n.West = curr
			curr = curr.East
		case 'W':
			p := Point{X: curr.loc.X - 1, Y: curr.loc.Y}
			n, ok := nodes[p]
			if !ok {
				n = NewNode(p, i)
				nodes[p] = n
			}
			curr.West = n
			n.East = curr
			curr = curr.West

		// handle branches
		// make recursives calls to this, or use a stack?
		case '(':
			stack = append(stack, curr)
			qi++
		case '|':
			curr = stack[qi]
		case ')':
			curr = stack[qi]
			stack = stack[:qi]
			qi--
		}
	}

	//fmt.Printf("stack: %d with qi %d\n", len(stack), qi)
	return start, nodes
}

func (n Node) isTerminal() bool {
	return len(n.adj()) == 1
}

func (n Node) adj() []Node {
	var a []Node

	if *n.North != empty {
		a = append(a, *n.North)
	}
	if *n.South != empty {
		a = append(a, *n.South)
	}
	if *n.East != empty {
		a = append(a, *n.East)
	}
	if *n.West != empty {
		a = append(a, *n.West)
	}
	return a
}

type NodePath struct {
	node Node
	dist int
}

func path(start, end Node) int {
	queue := []NodePath{NodePath{start, 0}}
	visited := make(map[Node]struct{})

	// iter := 0
	for len(queue) != 0 {
		// fmt.Printf("pathing iter %d queue:%d\n", iter, len(queue))
		c := queue[0]
		queue = queue[1:]

		if _, ok := visited[c.node]; ok {
			// fmt.Printf("continuing.... %d\n", iter)
			continue
		}
		visited[c.node] = struct{}{}

		if c.node == end {
			// fmt.Printf("path to end discovered in %d iterations. distance: %d\n", iter, c.dist)
			return c.dist
		}

		for _, n := range c.node.adj() {
			queue = append(queue, NodePath{n, c.dist + 1})
		}

		// iter++
	}

	fmt.Println("!!! ERROR !!! This should be solvable")
	return -1
}

func part1(regex string) {
	fmt.Println("Part 1")

	start, nodes := buildGraph(regex)
	//fmt.Printf("graph built with %d nodes\n", len(nodes))
	//fmt.Println(start)

	var terminals []*Node
	for _, n := range nodes {
		if n.isTerminal() && n != start {
			terminals = append(terminals, n)
		}
	}
	// fmt.Printf("Terminals detected: %d\n", len(terminals))

	// figure out the shortest path from the terminal to start
	dist := -1
	for _, t := range terminals {
		// fmt.Printf("pathing %dth terminal id: %d\n", i, t.ID)
		d := path(*start, *t)
		if dist < d {
			dist = d
		}
	}

	fmt.Printf("The shortest path to the furthest room: %d\n", dist)
}

func part2(regex string) {
	fmt.Println("Part 2")

	start, nodes := buildGraph(regex)

	// figure out the shortest path from the all nodes to start
	count := 0
	iter := 0
	for _, n := range nodes {
		//fmt.Printf("iter %d pathing to %d out of %d\n", iter, n.ID, len(nodes))

		d := path(*start, *n)
		if d >= 1000 {
			count++
		}

		iter++
	}

	fmt.Printf("Part 2: %d\n", count)
}
