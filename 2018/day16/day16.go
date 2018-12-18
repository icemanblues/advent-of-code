package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InstTest struct {
	Before [4]int
	Inst   [4]int
	After  [4]int
}

func readInput(filename string) []*InstTest {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return nil
	}
	defer file.Close()

	var insts []*InstTest
	var it *InstTest
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Before: [3, 2, 1, 1]
		// 9 2 1 2
		// After:  [3, 2, 2, 1]
		line := scanner.Text()
		words := strings.Split(line, ":")
		if words[0] == "Before" {
			it.Before = parseArray(words[1])
		} else if words[0] == "After" {
			it.After = parseArray(words[1])
			insts = append(insts, it)

			it = &InstTest{}
		} else {
			it.Inst = parseArray("[" + line + "]")
		}

	}
	return insts
}

// [3, 2, 1, 1]
func parseArray(s string) [4]int {
	a := strings.Split(s, "[")
	b := strings.Split(a[1], "]")
	c := strings.Split(b[0], ", ")

	var z [4]int
	for i, d := range c {
		e, err := strconv.Atoi(d)
		if err != nil {
			fmt.Printf("!!! ERROR !!! Unable to parse number %v\n", d)
		}
		z[i] = e
	}

	return z
}

func main() {
	fmt.Println("Day 16: Chronal Classification")

	part1("test.txt")
	part2()
}

func makeInst() map[string]func(reg [4]int, a, b, c int) {
	inst := make(map[string]func(reg [4]int, a, b, c int))

	inst["addr"] = func(reg [4]int, a, b, c int) {
		reg[c] = reg[a] + reg[b]
	}

	return inst
}

func part1(fn string) {
	fmt.Println("Part 1")

	insts := readInput(fn)
	for _, it := range insts {
		fmt.Println(it)
	}

	// 4 registers
	// reg := [4]int{}

	// op code input input output
}

func part2() {
	fmt.Println("Part 2")
}
