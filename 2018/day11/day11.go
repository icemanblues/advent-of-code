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
	fmt.Println("Day 11: Chronal Charge")

	const gridSerial = 8444
	part1(gridSerial)
	part2(gridSerial)
}

func part1(gridSerial int) {
	fmt.Println("Part 1")

	// testing power level
	// fmt.Printf("power level: 4 %v\n", powerLevel(3, 5, 8))
	// fmt.Printf("power level: -5 %v\n", powerLevel(122, 79, 57))
	// fmt.Printf("power level: 0 %v\n", powerLevel(217, 196, 39))
	// fmt.Printf("power level: 4 %v\n", powerLevel(101, 153, 71))

	// 1 to 300
	var grid [301][301]int
	for i := 1; i <= 300; i++ {
		for j := 1; j <= 300; j++ {
			grid[i][j] = powerLevel(i, j, gridSerial)
		}
	}

	var bestX, bestY, bestSum int
	for i := 1; i <= 298; i++ {
		for j := 1; j <= 298; j++ {
			sum := grid[i][j] + grid[i+1][j] + grid[i+2][j] +
				grid[i][j+1] + grid[i+1][j+1] + grid[i+2][j+1] +
				grid[i][j+2] + grid[i+1][j+2] + grid[i+2][j+2]
			if bestSum < sum {
				bestSum = sum
				bestX = i
				bestY = j
			}
		}
	}

	fmt.Printf("%v,%v %v\n", bestX, bestY, bestSum)
}

func powerLevel(x, y, gridSerial int) int {
	rackID := x + 10
	power := rackID * y
	power += gridSerial
	power *= rackID

	// get hundredth digit
	tenOneDigit := power % 100
	hundredthDigit := power % 1000
	power = (hundredthDigit - tenOneDigit) / 100
	power -= 5

	return power
}

func part2(gridSerial int) {
	fmt.Println("Part 2")

	// 1 to 300
	var grid [301][301]int
	for i := 1; i <= 300; i++ {
		for j := 1; j <= 300; j++ {
			grid[i][j] = powerLevel(i, j, gridSerial)
		}
	}

	var bestX, bestY, bestSize, bestSum int
	for s := 1; s <= 300; s++ {
		limit := 301 - s
		for i := 1; i <= limit; i++ {
			for j := 1; j <= limit; j++ {

				sum := 0
				for k := 0; k < s; k++ {
					for l := 0; l < s; l++ {
						sum += grid[i+k][j+l]
					}
				}

				if bestSum < sum {
					bestSum = sum
					bestX = i
					bestY = j
					bestSize = s
				}
			}
		}
	}

	fmt.Printf("%v,%v,%v %v\n", bestX, bestY, bestSize, bestSum)
}
