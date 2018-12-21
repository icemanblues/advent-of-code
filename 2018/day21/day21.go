package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	name   string
	values [3]int
}

func readInput(filename string) (int, []Instruction) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return 0, nil
	}
	defer file.Close()

	var ip int
	var insts []Instruction
	lineNum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// this is the instruction pointer
		// #ip 0
		if lineNum == 0 {
			i, err := strconv.Atoi(strings.Split(line, " ")[1])
			if err != nil {
				fmt.Printf("!!! ERROR !!! Unable to parse instruction pointer %v\n", line)
				break
			}
			ip = i
		} else {
			// parse the instruction
			// seti 5 0 1
			words := strings.Split(line, " ")
			a, _ := strconv.Atoi(words[1])
			b, _ := strconv.Atoi(words[2])
			c, _ := strconv.Atoi(words[3])
			insts = append(insts, Instruction{
				name:   words[0],
				values: [3]int{a, b, c},
			})
		}

		lineNum++
	}
	return ip, insts
}

func main() {
	fmt.Println("Day 21: Chronal Conversion")

	part1()
	part2()
}

func makeInst() map[string]func(reg *[6]int, a, b, c int) {
	inst := make(map[string]func(reg *[6]int, a, b, c int))

	inst["addr"] = func(reg *[6]int, a, b, c int) {
		reg[c] = reg[a] + reg[b]
	}

	inst["addi"] = func(reg *[6]int, a, b, c int) {
		reg[c] = reg[a] + b
	}

	inst["mulr"] = func(reg *[6]int, a, b, c int) {
		reg[c] = reg[a] * reg[b]
	}

	inst["muli"] = func(reg *[6]int, a, b, c int) {
		reg[c] = reg[a] * b
	}

	inst["banr"] = func(reg *[6]int, a, b, c int) {
		reg[c] = reg[a] & reg[b]
	}

	inst["bani"] = func(reg *[6]int, a, b, c int) {
		reg[c] = reg[a] & b
	}

	inst["borr"] = func(reg *[6]int, a, b, c int) {
		reg[c] = reg[a] | reg[b]
	}

	inst["bori"] = func(reg *[6]int, a, b, c int) {
		reg[c] = reg[a] | b
	}

	inst["setr"] = func(reg *[6]int, a, b, c int) {
		reg[c] = reg[a]
	}

	inst["seti"] = func(reg *[6]int, a, b, c int) {
		reg[c] = a
	}

	// gtir (greater-than immediate/register) sets register C to 1 if value A is greater than register B. Otherwise, register C is set to 0.
	inst["gtir"] = func(reg *[6]int, a, b, c int) {
		if a > reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	}
	// gtri (greater-than register/immediate) sets register C to 1 if register A is greater than value B. Otherwise, register C is set to 0.
	inst["gtri"] = func(reg *[6]int, a, b, c int) {
		if reg[a] > b {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	}
	// gtrr (greater-than register/register) sets register C to 1 if register A is greater than register B. Otherwise, register C is set to 0.
	inst["gtrr"] = func(reg *[6]int, a, b, c int) {
		if reg[a] > reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	}

	// eqir (equal immediate/register) sets register C to 1 if value A is equal to register B. Otherwise, register C is set to 0.
	inst["eqir"] = func(reg *[6]int, a, b, c int) {
		if a == reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	}
	// eqri (equal register/immediate) sets register C to 1 if register A is equal to value B. Otherwise, register C is set to 0.
	inst["eqri"] = func(reg *[6]int, a, b, c int) {
		if reg[a] == b {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	}
	// eqrr (equal register/register) sets register C to 1 if register A is equal to register B. Otherwise, register C is set to 0.	}
	inst["eqrr"] = func(reg *[6]int, a, b, c int) {
		if reg[a] == reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	}

	return inst
}

func part1() {
	fmt.Println("Part 1")

	ipReg, insts := readInput("input21.txt")
	fmt.Printf("ipReg: %v with %v instructions in program\n", ipReg, len(insts))

	functors := makeInst()

	limit := 10000000000
	for reg0 := 0; reg0 < 1; reg0++ {
		// set the value for register 0
		reg := &[6]int{}
		reg[0] = 0 //212115

		var reg3Value int
		visited := make(map[int]struct{})

		ip := 0
		count := 0
		for ip >= 0 && ip < len(insts) {
			// fmt.Printf("ip=%2d %v \n", ip, reg)

			in := insts[ip]
			reg[ipReg] = ip
			if ip == 28 {
				// fmt.Printf("ip=%v %v %v reg3=%d\n", ip, reg, in, reg[3])
				if _, ok := visited[reg[3]]; ok {
					fmt.Printf("Loop detected. last value found: %d\n", reg3Value)
					break
				}
				visited[reg[3]] = struct{}{}
				reg3Value = reg[3]
			}

			functors[in.name](reg, in.values[0], in.values[1], in.values[2])

			ip = reg[ipReg]
			ip++
			count++

			if count == limit {
				break
			}
		}

		if count != limit {
			fmt.Printf("! register 0 value %v completed with %v executions\n", reg0, count)
		}
	}
}

func part2() {
	fmt.Println("Part 2")
}
