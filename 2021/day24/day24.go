package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "24"
	dayTitle = "Arithmetic Logic Unit"
)

type Alu struct {
	registers map[string]int
	model     Model
	index     int
}

func (alu *Alu) String() string {
	return fmt.Sprintf("model: %v index: %v regs: %v", alu.model, alu.index, alu.registers)
}

func (alu *Alu) Lookup(arg string) int {
	i, err := strconv.Atoi(arg)
	if err != nil { // not an int, so return the register value
		return alu.registers[arg]
	}
	return i
}

func (alu *Alu) Apply(inst Inst) error {
	switch inst.Op {
	case "inp":
		inp := alu.model[alu.index]
		if inp == 0 {
			return errors.New("Model number cannot contain zero!")
		}
		alu.registers[inst.Arg1] = inp
		alu.index++
	case "add":
		alu.registers[inst.Arg1] = alu.Lookup(inst.Arg1) + alu.Lookup(inst.Arg2)
	case "mul":
		alu.registers[inst.Arg1] = alu.Lookup(inst.Arg1) * alu.Lookup(inst.Arg2)
	case "div":
		arg2 := alu.Lookup(inst.Arg2)
		if arg2 == 0 {
			return errors.New("divide by zero")
		}
		alu.registers[inst.Arg1] = alu.Lookup(inst.Arg1) / arg2
	case "mod":
		arg2 := alu.Lookup(inst.Arg2)
		if arg2 == 0 {
			return errors.New("modulus by zero")
		}
		alu.registers[inst.Arg1] = alu.Lookup(inst.Arg1) % arg2
	case "eql":
		if alu.Lookup(inst.Arg1) == alu.Lookup(inst.Arg2) {
			alu.registers[inst.Arg1] = 1
		} else {
			alu.registers[inst.Arg1] = 0
		}
	default:
		return fmt.Errorf("Unknown operator: %v\n", inst.Op)
	}

	return nil
}

func NewAlu(model Model) Alu {
	return Alu{make(map[string]int), model, 0}
}

type Inst struct {
	Op, Arg1, Arg2 string
}

func InstParse(s string) Inst {
	fields := strings.Fields(s)
	if len(fields) == 2 {
		return Inst{fields[0], fields[1], ""}
	}
	return Inst{fields[0], fields[1], fields[2]}
}

func Instructions(input []string) []Inst {
	insts := make([]Inst, 0, len(input))
	for _, inp := range input {
		insts = append(insts, InstParse(inp))
	}
	return insts
}

type Model [14]int

const (
	minModel int = 11111111111111
	maxModel int = 99999999999999
)

func IntToModel(a int) (Model, error) {
	var m Model
	size := len(m)
	for i := 0; i < size; i++ {
		ones := a % 10
		if ones == 0 {
			return m, errors.New("Contains a zero")
		}
		m[size-i-1] = ones
		a = a / 10
	}
	return m, nil
}

func ModelToInt(model Model) int {
	size := len(model)
	sum := 0
	for i := 0; i < size; i++ {
		sum += model[i] * int(math.Pow10(size-i-1))
	}
	return sum
}

func Run(insts []Inst, model Model) (Alu, error) {
	alu := NewAlu(model)
	for _, inst := range insts {
		err := alu.Apply(inst)
		if err != nil {
			return alu, err
		}
	}
	return alu, nil
}

// RunMonad returns true if the model number is valid
func RunMonad(insts []Inst, model Model) (bool, error) {
	alu, err := Run(insts, model)
	if err != nil {
		return false, err
	}

	return alu.registers["z"] == 0, nil
}

func GridSearch(insts []Inst, model Model, mi int, max int) int {
	if mi == len(model) {
		fmt.Printf("Grid Searching... %v\n", model)
		valid, _ := RunMonad(insts, model)
		if valid {
			if v := ModelToInt(model); v > max {
				return v
			}
			return max
		}
		return max
	}

	for n := 1; n <= 9; n++ {
		model[mi] = n
		b := GridSearch(insts, model, mi+1, max)
		if b > max {
			max = b
		}
	}

	return max
}

func part1() {
	fmt.Println("Part 1")

	// Parse input into slice of instructions
	input, _ := util.ReadInput("input.txt")
	insts := Instructions(input)
	model, _ := IntToModel(minModel)
	ans := GridSearch(insts, model, 0, minModel)
	fmt.Printf("best: %v\n", ans)
}

func part2() {
	fmt.Println("Part 2")
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
