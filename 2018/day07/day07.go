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

func readInput(filename string) map[string]*Node {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return nil
	}
	defer file.Close()

	// Step C must be finished before step A can begin.
	nameNode := make(map[string]*Node)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		fmt.Println(words)
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

		fmt.Println(n1)
		fmt.Println(n2)
		n1.deps = append(n1.deps, *n2)
	}

	return nameNode
}

func main() {
	fmt.Println("Day 07: The Sum of Its Parts")

	part1()
	part2()
}

func part1() {
	fmt.Println("Part 1")

	nameNode := readInput("test.txt")
	fmt.Println(nameNode)
}

func part2() {
	fmt.Println("Part 2")
}
