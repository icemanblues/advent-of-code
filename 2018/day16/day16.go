package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Println("Day 16: Chronal Classification")

	part1()
	part2()
}

func part1() {
	fmt.Println("Part 1")
}

func part2() {
	fmt.Println("Part 2")
}
