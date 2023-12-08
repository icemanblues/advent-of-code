package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "08"
	dayTitle = "Haunted Wasteland"
)

var START, END string = "AAA", "ZZZ"

type Node struct {
	Name, Left, Right string
}

type Directions struct {
	path  []rune
	steps int
}

func (d *Directions) Next() rune {
	r := d.path[d.steps%len(d.path)]
	d.steps++
	return r
}

func ParseMap(filename string) (Directions, map[string]Node) {
	input, _ := util.ReadInput(filename)
	d := Directions{[]rune(input[0]), 0}
	nodeList := make(map[string]Node)
	for _, line := range input[2:] {
		nodeAdjList := strings.Split(line, " = ")
		leftRight := strings.Split(nodeAdjList[1], ", ")
		nodeList[nodeAdjList[0]] = Node{
			Name:  nodeAdjList[0],
			Left:  leftRight[0][1:],
			Right: leftRight[1][:len(leftRight[1])-1],
		}
	}
	return d, nodeList
}

func StepCount(d Directions, nodeMap map[string]Node, start string, ends map[string]struct{}) int {
	for start != END {
		next := d.Next()
		switch next {
		case 'L':
			start = nodeMap[start].Left
		case 'R':
			start = nodeMap[start].Right
		default:
			fmt.Printf("ERROR: Unknown next step: %v\n", next)
			break
		}
		if _, ok := ends[start]; ok {
			break
		}
	}
	return d.steps
}

func part1() {
	d, nodeMap := ParseMap("input.txt")
	fmt.Printf("Part 1: %v\n", StepCount(d, nodeMap, START, map[string]struct{}{END: {}}))
}

func part2() {
	d, nodeMap := ParseMap("input.txt")
	var starts []string
	ends := make(map[string]struct{})
	for name := range nodeMap {
		if strings.HasSuffix(name, "A") {
			starts = append(starts, name)
		}
		if strings.HasSuffix(name, "Z") {
			ends[name] = struct{}{}
		}
	}
	steps := make([]int, len(starts), len(starts))
	for i, curr := range starts {
		steps[i] = StepCount(Directions{d.path, 0}, nodeMap, curr, ends)
	}
	fmt.Printf("Part 2: %v\n", util.LCM(steps[0], steps[1], steps[2:]...))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
