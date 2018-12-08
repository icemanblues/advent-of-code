package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	fmt.Println("Day 08: Memory Maneuver")

	// part1()
	part2()
}

type Node struct {
	numChild    int
	numMetaData int
	metadata    []int
	children    []*Node
}

func part1() {
	fmt.Println("Part 1")

	line := readInput("input08.txt")[0]
	tokens := strings.Split(line, " ")
	_, _, sum := buildNode(tokens, 0)
	fmt.Println(sum)

}

func buildNode(l []string, i int) (*Node, int, int) {
	sum := 0

	numChild, _ := strconv.Atoi(l[i])
	numMetaData, _ := strconv.Atoi(l[i+1])
	n := &Node{numChild, numMetaData, nil, nil}

	k := i + 2
	for kid := 0; kid < numChild; kid++ {
		child, j, jmd := buildNode(l, k)
		n.children = append(n.children, child)
		k = j
		sum += jmd
	}

	for m := 0; m < numMetaData; m++ {
		md, _ := strconv.Atoi(l[k])
		sum += md
		n.metadata = append(n.metadata, md)
		k++
	}

	// fmt.Println(n)
	return n, k, sum
}

func part2() {
	fmt.Println("Part 2")

	line := readInput("input08.txt")[0]
	tokens := strings.Split(line, " ")
	root, _, _ := buildNode(tokens, 0)
	fmt.Println(value(*root))
}

func value(n Node) int {
	if n.numChild == 0 {
		sum := 0
		for _, md := range n.metadata {
			sum += md
		}
		return sum
	}

	v := 0
	for _, md := range n.metadata {
		if md == 0 {
			continue
		}
		if md >= len(n.children)+1 {
			continue
		}
		v += value(*n.children[md-1])
	}
	return v
}
