package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "09"
	dayTitle = "Smoke Basin"
)

func Parse(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0, len(line))
		for _, r := range line {
			row = append(row, util.MustAtoi(string(r)))
		}
		grid = append(grid, row)
	}

	return grid, nil
}

func part1() {
	fmt.Println("Part 1")
	grid, _ := Parse("input.txt")
	risk := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			top, left, right, bot := true, true, true, true
			if x > 0 {
				left = grid[y][x-1] > grid[y][x]
			}
			if x < len(grid[y])-1 {
				right = grid[y][x+1] > grid[y][x]
			}
			if y > 0 {
				top = grid[y-1][x] > grid[y][x]
			}
			if y < len(grid)-1 {
				bot = grid[y+1][x] > grid[y][x]
			}

			if top && bot && left && right {
				//fmt.Printf("low point: x:%v y:%v height:%v\n", x, y, grid[y][x])
				risk += grid[y][x] + 1
			}
		}
	}

	fmt.Printf("Part 1: %v\n", risk)
}

func part2() {
	fmt.Println("Part 2")
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
