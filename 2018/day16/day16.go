package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InstTest struct {
	Before [4]int
	Inst   [4]int
	After  [4]int
}

const splitter = 3045

func readInput(filename string) ([]*InstTest, [][4]int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return nil, nil
	}
	defer file.Close()

	lineNum := 0

	var program [][4]int
	var insts []*InstTest
	it := &InstTest{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNum++

		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		words := strings.Split(line, ":")

		if len(words) == 1 {
			if lineNum > splitter {
				program = append(program, parseInst(line))
			} else {
				it.Inst = parseInst(line)
				continue
			}
		}

		if words[0] == "Before" {
			it.Before = parseArray(words[1])
		} else if words[0] == "After" {
			it.After = parseArray(words[1])
			insts = append(insts, it)

			it = &InstTest{}
		}

	}
	return insts, program
}

func parseInst(s string) [4]int {
	var z [4]int
	c := strings.Split(s, " ")
	for i, d := range c {
		e, err := strconv.Atoi(d)
		if err != nil {
			fmt.Printf("!!! ERROR !!! Unable to parse number %v\n", d)
		}
		z[i] = e
	}
	return z
}

// [3, 2, 1, 1]
func parseArray(s string) [4]int {
	a := strings.Split(s, "[")
	b := strings.Split(a[1], "]")
	c := strings.Split(b[0], ", ")

	var z [4]int
	for i, d := range c {
		e, err := strconv.Atoi(d)
		if err != nil {
			fmt.Printf("!!! ERROR !!! Unable to parse number %v\n", d)
		}
		z[i] = e
	}

	return z
}

func main() {
	fmt.Println("Day 16: Chronal Classification")

	part1("test.txt")
	part1("input16.txt")
	// part2()
}

func makeOpName() map[int]string {
	opName := make(map[int]string)
	opName[4] = "addi"
	opName[11] = "borr"
	opName[8] = "muli"
	opName[9] = "bori"
	opName[15] = "mulr"
	opName[1] = "seti"
	opName[14] = "addr"
	opName[7] = "gtri"
	opName[3] = "eqrr"
	opName[2] = "eqri"
	opName[0] = "eqir"
	opName[6] = "gtrr"
	opName[12] = "gtir"
	opName[5] = "setr"
	opName[13] = "banr"
	opName[10] = "bani"
	return opName
}

func makeInst() map[string]func(reg *[4]int, a, b, c int) {
	inst := make(map[string]func(reg *[4]int, a, b, c int))

	inst["addr"] = func(reg *[4]int, a, b, c int) {
		reg[c] = reg[a] + reg[b]
	}

	inst["addi"] = func(reg *[4]int, a, b, c int) {
		reg[c] = reg[a] + b
	}

	inst["mulr"] = func(reg *[4]int, a, b, c int) {
		reg[c] = reg[a] * reg[b]
	}

	inst["muli"] = func(reg *[4]int, a, b, c int) {
		reg[c] = reg[a] * b
	}

	inst["banr"] = func(reg *[4]int, a, b, c int) {
		reg[c] = reg[a] & reg[b]
	}

	inst["bani"] = func(reg *[4]int, a, b, c int) {
		reg[c] = reg[a] & b
	}

	inst["borr"] = func(reg *[4]int, a, b, c int) {
		reg[c] = reg[a] | reg[b]
	}

	inst["bori"] = func(reg *[4]int, a, b, c int) {
		reg[c] = reg[a] | b
	}

	inst["setr"] = func(reg *[4]int, a, b, c int) {
		reg[c] = reg[a]
	}

	inst["seti"] = func(reg *[4]int, a, b, c int) {
		reg[c] = a
	}

	// gtir (greater-than immediate/register) sets register C to 1 if value A is greater than register B. Otherwise, register C is set to 0.
	inst["gtir"] = func(reg *[4]int, a, b, c int) {
		if a > reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	}
	// gtri (greater-than register/immediate) sets register C to 1 if register A is greater than value B. Otherwise, register C is set to 0.
	inst["gtri"] = func(reg *[4]int, a, b, c int) {
		if reg[a] > b {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	}
	// gtrr (greater-than register/register) sets register C to 1 if register A is greater than register B. Otherwise, register C is set to 0.
	inst["gtrr"] = func(reg *[4]int, a, b, c int) {
		if reg[a] > reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	}

	// eqir (equal immediate/register) sets register C to 1 if value A is equal to register B. Otherwise, register C is set to 0.
	inst["eqir"] = func(reg *[4]int, a, b, c int) {
		if a == reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	}
	// eqri (equal register/immediate) sets register C to 1 if register A is equal to value B. Otherwise, register C is set to 0.
	inst["eqri"] = func(reg *[4]int, a, b, c int) {
		if reg[a] == b {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	}
	// eqrr (equal register/register) sets register C to 1 if register A is equal to register B. Otherwise, register C is set to 0.	}
	inst["eqrr"] = func(reg *[4]int, a, b, c int) {
		if reg[a] == reg[b] {
			reg[c] = 1
		} else {
			reg[c] = 0
		}
	}

	return inst
}

func intersect(a, b map[string]struct{}) map[string]struct{} {
	c := make(map[string]struct{})
	for k := range a {
		if _, ok := b[k]; ok {
			c[k] = struct{}{}
		}
	}
	return c
}

func part1(fn string) {
	fmt.Println("Part 1")

	functors := makeInst()
	reg := &[4]int{}
	insts, program := readInput(fn)

	opMap := make(map[int]map[string]struct{})
	opName := makeOpName()
	count := 0

	for _, it := range insts {
		seta := make(map[string]struct{})
		for name, f := range functors {
			// copy values into reg
			for i, v := range it.Before {
				reg[i] = v
			}

			// apply
			f(reg, it.Inst[1], it.Inst[2], it.Inst[3])

			// confirm the result
			if *reg == it.After {
				seta[name] = struct{}{}
			}
		}

		if len(seta) >= 3 {
			count++
		}

		// do an intersection of seta with opmap
		op := it.Inst[0]
		set, ok := opMap[op]
		if !ok {
			set = seta
			opMap[op] = seta
		}

		opMap[op] = intersect(set, seta)
	}

	fmt.Printf("Part 1: %d out of %d\n", count, len(insts))

	// do the sieve to figure out the unique func names to op codes
	// used := make(map[int]struct{})
	// for op, set := range opMap {
	// 	if _, ok := used[op]; len(set) == 1 && !ok {
	// 		used[op] = struct{}{}

	// 		// extract the one name
	// 		var name string
	// 		for k := range set {
	// 			name = k
	// 		}

	// 		for o, s := range opMap {
	// 			if o == op {
	// 				continue
	// 			}
	// 			delete(s, name)
	// 		}
	// 	}
	// }

	cpu := &[4]int{}
	for _, p := range program {
		op := p[0]
		name := opName[op]
		f := functors[name]
		f(cpu, p[1], p[2], p[3])
	}

	fmt.Printf("register 0: %d\n", cpu[0])
}
