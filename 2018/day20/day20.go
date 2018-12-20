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

func main() {
	fmt.Println("Day 20: A Regular Map")

	test1 := "^WNE$"                                     // 3
	test2 := "^ENWWW(NEEE|SSE(EE|N))$"                   // 10
	test3 := "^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$" // 18
	// test4 := "^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$"               // 23
	// test5 := "^WSSEESWWWNW(S|NENNEEEENN(ESSSSW(NWSW|SSEN)|WSWWN(E|WWS(E|SS))))$" // 31
	// input := readInput("input20.txt")

	part1(test1)
	part1(test2)
	part1(test3)
	// part2()
}

type Point struct {
	X int
	Y int
}
type Node struct {
	loc   Point
	North *Node
	East  *Node
	South *Node
	West  *Node
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

func buildGraph(regex string) {
	s := regex[1 : len(regex)-1]
	start := &Node{}
	curr := start

	nodes := make(map[Point]*Node)

	// build the graph
	for _, r := range s {
		switch r {
		case 'N':
			p := Point{X: curr.loc.X, Y: curr.loc.Y + 1}
			n, ok := nodes[p]
			if !ok {
				n := &Node{loc: p}
				nodes[p] = n
			}
			curr.North = n
			n.South = curr
			curr = curr.North
		case 'S':
			curr.South = &Node{}
			curr = curr.South
		case 'E':
			curr.East = &Node{}
			curr = curr.East
		case 'W':
			curr.West = &Node{}
			curr = curr.West
		}
	}
}

func rec() {

}

func part1(regex string) {
	fmt.Println("Part 1")

}

func part2() {
	fmt.Println("Part 2")
}
