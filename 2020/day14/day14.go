package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

func parseLine(line string) (string, string) {
	leftRight := strings.Split(line, " = ")
	if leftRight[0] == "mask" {
		return leftRight[0], leftRight[1]
	}

	// must be a memory instruction
	mapIndex := leftRight[0]
	index := mapIndex[4 : len(mapIndex)-1]
	return index, leftRight[1]
}

func memSum(mem map[int]int) int {
	sum := 0
	for _, v := range mem {
		sum += v
	}
	return sum
}

func runMask(lines []string) int {
	var mask string
	mem := make(map[int]int)
	for _, line := range lines {
		cmd, arg := parseLine(line)
		if cmd == "mask" {
			mask = arg
			continue
		}

		// apply the mask
		addr, _ := strconv.Atoi(cmd)
		value, _ := strconv.Atoi(arg)

		number := 0
		for i, r := range mask {
			bitIndex := 35 - i
			bit := 1 << bitIndex
			switch r {
			case '1': // force it to 1
				number += bit
			case '0': // force it to 0, skip
			case 'X': // copy the bit from the number
				number += value & bit
			default:
				panic("mask is messed up")
			}
		}

		// update memory
		mem[addr] = number
	}
	return memSum(mem)
}

func applyMaskAddr(mask string, addr int) (int, []int) {
	var floating []int
	memAddr := 0
	for i, r := range mask {
		bitIndex := 35 - i
		bit := 1 << bitIndex

		switch r {
		case '1': // force it to 1
			memAddr += bit
		case '0': // take the value that is there
			memAddr += addr & bit
		case 'X': // floating
			floating = append(floating, bitIndex)
		default:
			panic("mask is messed up")
		}
	}

	return memAddr, floating
}

func runMemAddrDecoder(lines []string) int {
	var mask string
	mem := make(map[int]int)
	for _, line := range lines {
		cmd, arg := parseLine(line)
		if cmd == "mask" {
			mask = arg
			continue
		}

		// must be a memory instruction
		index, _ := strconv.Atoi(cmd)
		value, _ := strconv.Atoi(arg)

		// apply the mask to memory address
		memAddr, floating := applyMaskAddr(mask, index)

		// update memory, apply floating bits (power set)
		limit := 1 << len(floating)
		for i := 0; i < limit; i++ {
			addr := memAddr
			for j, floater := range floating {
				bitIndex := 1 << j
				valueZero := i&bitIndex == 0

				membit := 1 << floater
				memZero := addr&membit == 0

				// both are zero, nothing to do
				// we want to set 1, but its current1

				// we want to set 1, but its current 0
				if !valueZero && memZero {
					addr = addr + membit
				}
				// we want to set 0, but its current 1
				if valueZero && !memZero {
					addr = addr - membit
				}
			}
			mem[addr] = value
		}
	}

	return memSum(mem)
}

func part1() {
	lines, _ := util.ReadInput("input.txt")
	fmt.Printf("Part 1: %v\n", runMask(lines))
}

func part2() {
	lines, _ := util.ReadInput("input.txt")
	fmt.Printf("Part 2: %v\n", runMemAddrDecoder(lines))
}

func main() {
	fmt.Println("Day 14: Docking Data")
	part1()
	part2()
}
