package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

func seatCount(seats [][]rune) int {
	count := 0
	for _, line := range seats {
		for _, s := range line {
			if s == '#' {
				count++
			}
		}
	}
	return count
}

func seatAdj(seatMap [][]rune) [][]rune {
	nextMap := make([][]rune, len(seatMap), len(seatMap))
	for i := range nextMap {
		nextMap[i] = make([]rune, len(seatMap[i]), len(seatMap[i]))
	}

	for i, line := range seatMap {
		for j, s := range line {
			if s == '.' {
				nextMap[i][j] = '.'
				continue
			}
			// count adjacent taken seats
			count := 0
			for ii := -1; ii <= 1; ii++ {
				for jj := -1; jj <= 1; jj++ {
					if ii == 0 && jj == 0 {
						continue
					}
					if i+ii >= 0 && i+ii < len(seatMap) && j+jj >= 0 && j+jj < len(line) {
						if seatMap[i+ii][j+jj] == '#' {
							count++
						}
					}
				}
			}
			// apply seat change logic
			if s == '#' && count >= 4 {
				nextMap[i][j] = 'L'
			} else if s == 'L' && count == 0 {
				nextMap[i][j] = '#'
			} else {
				nextMap[i][j] = s
			}
		}
	}
	return nextMap
}

func isOccupied(seatMap [][]rune, x, y int, dx, dy int) int {
	for x+dx >= 0 && x+dx < len(seatMap) && y+dy >= 0 && y+dy < len(seatMap[x]) {
		if seatMap[x+dx][y+dy] == '#' {
			return 1
		}
		if seatMap[x+dx][y+dy] == 'L' {
			return 0
		}

		x += dx
		y += dy
	}
	return 0
}

func seatLos(seatMap [][]rune) [][]rune {
	nextMap := make([][]rune, len(seatMap), len(seatMap))
	for i := range nextMap {
		nextMap[i] = make([]rune, len(seatMap[i]), len(seatMap[i]))
	}

	for i, line := range seatMap {
		for j, s := range line {
			if s == '.' {
				nextMap[i][j] = '.'
				continue
			}
			// count seats in line of sight
			count := 0
			count += isOccupied(seatMap, i, j, 1, 0)
			count += isOccupied(seatMap, i, j, -1, 0)
			count += isOccupied(seatMap, i, j, 0, 1)
			count += isOccupied(seatMap, i, j, 0, -1)
			count += isOccupied(seatMap, i, j, 1, 1)
			count += isOccupied(seatMap, i, j, 1, -1)
			count += isOccupied(seatMap, i, j, -1, 1)
			count += isOccupied(seatMap, i, j, -1, -1)
			// apply seat change logic
			if s == '#' && count >= 5 {
				nextMap[i][j] = 'L'
			} else if s == 'L' && count == 0 {
				nextMap[i][j] = '#'
			} else {
				nextMap[i][j] = s
			}
		}
	}
	return nextMap
}

func seatOfLife(filename string, tick func([][]rune) [][]rune) int {
	seatMap, _ := util.ReadRuneput(filename)
	prev := -1
	count := seatCount(seatMap)
	for count != prev {
		prev = count
		seatMap = tick(seatMap)
		count = seatCount(seatMap)
	}

	return count
}

func part1() {
	fmt.Printf("Part 1: %v\n", seatOfLife("input.txt", seatAdj))
}

func part2() {
	fmt.Printf("Part 2: %v\n", seatOfLife("input.txt", seatLos))
}

func main() {
	fmt.Println("Day 11: title")
	part1()
	part2()
}
