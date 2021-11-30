package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "07"
	dayTitle = "Some Assembly Required"
)

type BitOp struct {
	Op, In1, In2, Out string
}

func lookup(register map[string]uint16, s string) (uint16, bool) {
	i, err := strconv.Atoi(s)
	if err != nil {
		v, ok := register[s]
		return v, ok
	}
	return uint16(i), true
}

func readBitops(filename string) []BitOp {
	lines := util.MustRead("input.txt")
	insts := make([]BitOp, 0, len(lines))
	for _, line := range lines {
		inOut := strings.Split(line, " -> ")
		out := inOut[1]
		in := strings.Fields(inOut[0])
		if len(in) == 1 {
			insts = append(insts, BitOp{"SET", in[0], "0", out})
		} else if len(in) == 2 {
			insts = append(insts, BitOp{in[0], in[1], "1", out})
		} else {
			insts = append(insts, BitOp{in[1], in[0], in[2], out})
		}
	}
	return insts
}

func runInsts(insts []BitOp, registers map[string]uint16) {
	for _, inst := range insts {
		if _, ok := registers[inst.Out]; ok {
			continue
		}

		arg1, ok1 := lookup(registers, inst.In1)
		if !ok1 {
			continue
		}

		arg2, ok2 := lookup(registers, inst.In2)
		if !ok2 {
			continue
		}

		switch inst.Op {
		case "SET":
			registers[inst.Out] = arg1
		case "NOT":
			registers[inst.Out] = ^arg1
		case "AND":
			registers[inst.Out] = arg1 & arg2
		case "OR":
			registers[inst.Out] = arg1 | arg2
		case "LSHIFT":
			registers[inst.Out] = arg1 << arg2
		case "RSHIFT":
			registers[inst.Out] = arg1 >> arg2
		}
	}
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	insts := readBitops("input.txt")
	registers := make(map[string]uint16)
	aok := false
	for !aok {
		runInsts(insts, registers)
		_, aok = registers["a"]
	}
	fmt.Printf("Part 1: %v\n", registers["a"])

	reg2 := make(map[string]uint16)
	reg2["b"] = registers["a"]
	aok = false
	for !aok {
		runInsts(insts, reg2)
		_, aok = reg2["a"]
	}
	fmt.Printf("Part 2: %v\n", reg2["a"])
}
