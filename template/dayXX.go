package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	dayNum = "XX"
	dayTitle = "Title"
	subTitle1 = "Subtitle One"
	subTitle2 = "Subtitle Two"
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

func part1() {
	fmt.Printf("Part 1 : %v\n", subTitle1)
}

func part2() {
	fmt.Printf("Part 2 : %v\n", subTitle2)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)

	part1()
	part2()
}
