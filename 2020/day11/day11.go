package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/2020/util"
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

func seatOfLife(seatMap [][]rune) [][]rune {
	seatMapper := make([][]rune, len(seatMap), len(seatMap))
	for i := range seatMapper {
		seatMapper[i] = make([]rune, len(seatMap[i]), len(seatMap[i]))
	}

	for i, line := range seatMap {
		for j, s := range line {
			if s == '.' {
				seatMapper[i][j] = '.'
				continue
			}

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
			if s == '#' && count >= 4 {
				seatMapper[i][j] = 'L'
			} else if s == 'L' && count == 0 {
				seatMapper[i][j] = '#'
			} else {
				seatMapper[i][j] = s
			}
		}
	}
	return seatMapper
}

func seatOfVisible(seatMap [][]rune) [][]rune {
	seatMapper := make([][]rune, len(seatMap), len(seatMap))
	for i := range seatMapper {
		seatMapper[i] = make([]rune, len(seatMap[i]), len(seatMap[i]))
	}

	for i, line := range seatMap {
		for j, s := range line {
			if s == '.' {
				seatMapper[i][j] = '.'
				continue
			}

			count := 0
			// up
			for k := 1; j-k >= 0; k++ {
				if seatMap[i][j-k] == '.' {
					continue
				}
				if seatMap[i][j-k] == '#' {
					count++
					break
				}
				if seatMap[i][j-k] == 'L' {
					break
				}
			}

			// down
			for k := 1; j+k < len(seatMap[i]); k++ {
				if seatMap[i][j+k] == '.' {
					continue
				}
				if seatMap[i][j+k] == '#' {
					count++
					break
				}
				if seatMap[i][j+k] == 'L' {
					break
				}
			}

			// right
			for k := 1; i+k < len(seatMap); k++ {
				if seatMap[i+k][j] == '.' {
					continue
				}
				if seatMap[i+k][j] == '#' {
					count++
					break
				}
				if seatMap[i+k][j] == 'L' {
					break
				}
			}

			// left
			for k := 1; i-k >= 0; k++ {
				if seatMap[i-k][j] == '.' {
					continue
				}
				if seatMap[i-k][j] == '#' {
					count++
					break
				}
				if seatMap[i-k][j] == 'L' {
					break
				}
			}

			// ur
			for k := 1; i+k < len(seatMap) && j-k >= 0; k++ {
				if seatMap[i+k][j-k] == '.' {
					continue
				}
				if seatMap[i+k][j-k] == '#' {
					count++
					break
				}
				if seatMap[i+k][j-k] == 'L' {
					break
				}
			}

			// ul
			for k := 1; i-k >= 0 && j-k >= 0; k++ {
				if seatMap[i-k][j-k] == '.' {
					continue
				}
				if seatMap[i-k][j-k] == '#' {
					count++
					break
				}
				if seatMap[i-k][j-k] == 'L' {
					break
				}
			}

			// dr
			for k := 1; i+k < len(seatMap) && j+k < len(seatMap[i]); k++ {
				if seatMap[i+k][j+k] == '.' {
					continue
				}
				if seatMap[i+k][j+k] == '#' {
					count++
					break
				}
				if seatMap[i+k][j+k] == 'L' {
					break
				}
			}

			// dl
			for k := 1; i-k >= 0 && j+k < len(seatMap[i]); k++ {
				if seatMap[i-k][j+k] == '.' {
					continue
				}
				if seatMap[i-k][j+k] == '#' {
					count++
					break
				}
				if seatMap[i-k][j+k] == 'L' {
					break
				}
			}

			if s == '#' && count >= 5 {
				seatMapper[i][j] = 'L'
			} else if s == 'L' && count == 0 {
				seatMapper[i][j] = '#'
			} else {
				seatMapper[i][j] = s
			}
		}
	}
	return seatMapper
}

func seatIter(filename string) int {
	seatMap, _ := util.ReadRuneput(filename)
	prev := -1
	count := seatCount(seatMap)
	for count != prev {
		prev = count
		seatMap = seatOfLife(seatMap)
		count = seatCount(seatMap)
	}

	return count
}

func seatIterVis(filename string) int {
	seatMap, _ := util.ReadRuneput(filename)
	prev := -1
	count := seatCount(seatMap)
	for count != prev {
		prev = count
		seatMap = seatOfVisible(seatMap)
		count = seatCount(seatMap)
	}

	return count
}

func part1() {
	fmt.Printf("Part 1: %v\n", seatIter("input.txt"))
}

func part2() {
	fmt.Printf("Part 2: %v\n", seatIterVis("input.txt"))
}

func main() {
	fmt.Println("Day 11: title")
	part1()
	part2()
}
