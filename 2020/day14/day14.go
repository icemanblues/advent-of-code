package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "14"
	dayTitle = "Docking Data"
)

func runMask(lines []string) int {
	var mask string
	mem := make(map[int]int)
	for _, line := range lines {
		leftRight := strings.Split(line, " = ")
		if leftRight[0] == "mask" {
			mask = leftRight[1]
			continue
		}

		// must be a memory instruction
		mapIndex := leftRight[0]
		index, _ := strconv.Atoi(mapIndex[4 : len(mapIndex)-1])
		value, _ := strconv.Atoi(leftRight[1])

		// apply the mask
		number := 0
		for i, r := range mask {
			bit := 35 - i
			v := int(math.Pow(2, float64(bit)))
			switch r {
			case '1':
				// force it to 1
				number += v
			case '0':
				// force it to 0, skip
			case 'X':
				// copy the bit from the number
				n := value & v
				number += n
			default:
				panic("mask is messed up")
			}
		}

		// update memory
		mem[index] = number
	}

	sum := 0
	for _, v := range mem {
		sum += v
	}
	return sum
}

func runMemMask(lines []string) int {
	var mask string
	mem := make(map[int]int)
	for _, line := range lines {
		leftRight := strings.Split(line, " = ")
		if leftRight[0] == "mask" {
			mask = leftRight[1]
			continue
		}

		// must be a memory instruction
		mapIndex := leftRight[0]
		index, _ := strconv.Atoi(mapIndex[4 : len(mapIndex)-1])
		value, _ := strconv.Atoi(leftRight[1])

		// apply the mask
		number := value
		var floating []int
		iterMem := index
		for i, r := range mask {
			bit := 35 - i
			v := int(math.Pow(2, float64(bit)))
			memOk := iterMem & v
			switch r {
			case '1':
				// force it to 1
				if memOk == 0 {
					iterMem += v
				}
			case '0':
				// skip
			case 'X':
				floating = append(floating, 35-i)
			default:
				panic("mask is messed up")
			}
		}

		// update memory
		limit := int(math.Pow(2, float64(len(floating))))
		for i := 0; i < limit; i++ {
			addr := iterMem
			for j, floater := range floating {

				v := int(math.Pow(2, float64(j)))
				valueSet := i & v // do I want to put a 1 or 0

				membit := int(math.Pow(2, float64(floater)))
				memOk := addr & membit // the bit in the value, 0 if not present

				if valueSet == 0 && memOk == 0 { // both are zero, nothing to do
				}
				if valueSet != 0 && memOk == 0 { // we want to set 1, but its current 0
					addr = addr + membit
				}
				if valueSet == 0 && memOk != 0 { // we want to set 0, but its current 1
					addr = addr - membit
				}
				if valueSet != 0 && memOk != 0 { // we want to set 1, but its current 1
				}
			}
			mem[addr] = number
		}
	}

	sum := 0
	for _, v := range mem {
		sum += v
	}
	return sum
}

func part1() {
	lines, _ := util.ReadInput("input.txt")
	fmt.Printf("Part 1: %v\n", runMask(lines))
}

func part2() {
	lines, _ := util.ReadInput("input.txt")
	fmt.Printf("Part 2: %v\n", runMemMask(lines))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
