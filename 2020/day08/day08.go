package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "08"
	dayTitle = "Handheld Halting"
)

func runProg(lines []string) (int, bool) {
	acc := 0
	i := 0

	visited := make(map[int]struct{})

	for i >= 0 && i < len(lines) {
		if _, ok := visited[i]; ok {
			return acc, false
		}
		visited[i] = struct{}{}

		line := lines[i]
		words := strings.Split(line, " ")
		arg, _ := strconv.Atoi(words[1])
		switch words[0] {
		case "nop":
			i++
		case "acc":
			acc += arg
			i++
		case "jmp":
			i += arg
		default:
			fmt.Printf("unknown command %v\n", words[0])
		}
	}

	return acc, true
}

func fixProg(lines []string) int {
	for i, line := range lines {
		words := strings.Split(line, " ")
		switch words[0] {
		case "nop":
			fix := "jmp " + words[1]
			lines[i] = fix
			if acc, ok := runProg(lines); ok {
				return acc
			}
		case "jmp":
			fix := "nop " + words[1]
			lines[i] = fix
			if acc, ok := runProg(lines); ok {
				return acc
			}
		default:
			continue
		}
		lines[i] = line
	}
	panic("it finished, yet none of the fixes finished")
}

func part1() {
	lines, _ := util.ReadInput("input.txt")
	acc, _ := runProg(lines)
	fmt.Printf("Part 1: %v\n", acc)
}

func part2() {
	lines, _ := util.ReadInput("input.txt")
	fmt.Printf("Part 2: %v\n", fixProg(lines))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
