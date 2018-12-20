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
	fmt.Println("Day 19: Go With The Flow")

	// part1("test.txt")
	// part1("input19.txt", &[6]int{})
	// part1("input19.txt", &[6]int{1, 0, 0, 0, 0, 0})
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

func part1(fn string, reg *[6]int) {
	fmt.Println("Part 1")

	ipReg, insts := readInput(fn)
	fmt.Printf("ipReg: %v with %v instructions in program\n", ipReg, len(insts))

	functors := makeInst()

	// iter := 0
	ip := 0
	for ip >= 0 && ip < len(insts) {
		fmt.Printf("ip=%2d %v \n", ip, reg)
		// if iter%1000 == 0 {
		// 	fmt.Printf("%d: ip=%d %v \n", iter, ip, reg)
		// }
		// iter++

		in := insts[ip]
		reg[ipReg] = ip
		// fmt.Printf("ip=%v %v %v ", ip, reg, in)

		functors[in.name](reg, in.values[0], in.values[1], in.values[2])

		// fmt.Printf("%v\n", reg)

		ip = reg[ipReg]
		ip++
	}

	fmt.Printf("value of 0 register: %v\n", reg[0])
}

func part2() {
	fmt.Println("Part 2")

	const limit = 10551348
	const l = 3249
	sum := 0
	for i := 1; i <= l; i++ {
		if limit%i == 0 {
			num := limit / i
			sum = sum + i + num
		}
	}

	fmt.Println(sum)

}
