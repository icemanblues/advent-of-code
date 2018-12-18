package main

import (
	"bufio"
	"fmt"
	"os"
)

const sq = 50

func readInput(filename string) [sq][sq]rune {
	var grid [sq][sq]rune

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return grid
	}
	defer file.Close()

	lineNum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for i, s := range line {
			grid[i][lineNum] = rune(s)
		}

		lineNum++
	}
	return grid
}

func main() {
	fmt.Println("Day 18: Settlers of The North Pole")

	// part1("test.txt", 10)
	part1("input18.txt", 10)
	part1("input18.txt", 1000000000)
}

func printGrid(grid [sq][sq]rune) {
	for y := 0; y < sq; y++ {
		for x := 0; x < sq; x++ {
			fmt.Printf("%c", grid[x][y])
		}
		fmt.Println()
	}
}

func transition(grid [sq][sq]rune) [sq][sq]rune {
	var update [sq][sq]rune

	for y := 0; y < sq; y++ {
		for x := 0; x < sq; x++ {

			landCount := make(map[rune]int)
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if i == 0 && j == 0 {
						continue
					}
					if x+i < 0 || x+i >= sq {
						continue
					}
					if y+j < 0 || y+j >= sq {
						continue
					}

					r := grid[x+i][y+j]
					c, ok := landCount[r]
					if !ok {
						c = 0
					}
					landCount[r] = c + 1
				}
			}

			//open ground (.),
			//trees (|), or a
			//lumberyard (#)
			if grid[x][y] == '.' {
				if landCount['|'] >= 3 {
					update[x][y] = '|'
				} else {
					update[x][y] = '.'
				}
			}
			if grid[x][y] == '|' {
				if landCount['#'] >= 3 {
					update[x][y] = '#'
				} else {
					update[x][y] = '|'
				}
			}
			if grid[x][y] == '#' {
				if landCount['#'] >= 1 && landCount['|'] >= 1 {
					update[x][y] = '#'
				} else {
					update[x][y] = '.'
				}
			}
		}
	}

	return update
}

func part1(fn string, limit int) {
	fmt.Println("Part 1")

	grid := readInput(fn)
	// printGrid(grid)

	cache := make(map[[sq][sq]rune][sq][sq]rune)
	var list [][sq][sq]rune
	for i := 0; i < limit; i++ {

		if g, ok := cache[grid]; !ok {
			g = transition(grid)
			cache[grid] = g
			list = append(list, grid)
		} else {
			// cache hit, we are in a loop
			fmt.Printf("cache hit! loop detected on iteration %d\n", i)

			curr := -1
			for j, l := range list {
				if grid == l {
					curr = j
				}
			}
			fmt.Printf("curr: %d on iter %d with list size %d\n", curr, i, len(list))

			idx := (limit - curr) % (len(list) - curr)
			grid = list[curr+idx]
			break

		}
		grid = cache[grid]
	}

	// printGrid(grid)

	landCount := make(map[rune]int)
	for y := 0; y < sq; y++ {
		for x := 0; x < sq; x++ {
			c, ok := landCount[grid[x][y]]
			if !ok {
				c = 0
			}
			landCount[grid[x][y]] = c + 1
		}
	}
	numTrees := landCount['|']
	numLumberyard := landCount['#']
	solution := numLumberyard * numTrees
	fmt.Printf("trees: %d lumbeyards: %d solution: %d\n", numTrees, numLumberyard, solution)
}
