package main

import (
	"fmt"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "19"
	dayTitle = "Aplenty"
)

type Tool struct {
	x, m, a, s int
}

func (t Tool) rating() int {
	return t.x + t.m + t.a + t.s
}

type Inst struct {
	field string
	op    string
	value int
	next  string
}

func (inst Inst) Test(t Tool) bool {
	tv := t.x
	switch inst.field {
	case "x":
		tv = t.x
	case "m":
		tv = t.m
	case "a":
		tv = t.a
	case "s":
		tv = t.s
	}

	if inst.op == ">" {
		return tv > inst.value
	}
	return tv < inst.value

}

type Workflow struct {
	name  string
	insts []Inst
}

func (w Workflow) Apply(t Tool) string {
	for _, inst := range w.insts {
		if inst.op == "go" || inst.Test(t) {
			return inst.next
		}
	}
	return w.insts[len(w.insts)-1].next
}

func parse(filename string) (map[string]Workflow, []Tool) {
	workflows := make(map[string]Workflow)
	tools := make([]Tool, 0)
	input, _ := util.ReadInput(filename)
	isWorkflow := true
	for _, line := range input {
		if line == "" {
			isWorkflow = false
		} else if isWorkflow {
			s := strings.Split(line, "{")
			name := s[0]
			flowsline := s[1][:len(s[1])-1]
			flows := strings.Split(flowsline, ",")
			insts := make([]Inst, 0, len(flows))
			for _, flow := range flows {
				fline := strings.Split(flow, ":")
				if len(fline) == 1 { // just go to this location
					insts = append(insts, Inst{"go", "go", 0, fline[0]})
					continue
				}
				field := fline[0][:1]
				op := fline[0][1:2]
				value := util.MustAtoi(fline[0][2:])
				nextName := fline[1]
				insts = append(insts, Inst{field, op, value, nextName})
			}
			workflows[name] = Workflow{name, insts}
		} else {
			toolLine := line[1 : len(line)-1]
			parts := strings.Split(toolLine, ",")
			values := make([]int, 0, 4)
			for _, p := range parts {
				values = append(values, util.MustAtoi(strings.Split(p, "=")[1]))
			}
			tools = append(tools, Tool{values[0], values[1], values[2], values[3]})
		}
	}
	return workflows, tools
}

func isAccepted(workflows map[string]Workflow, t Tool) bool {
	name := "in"
	for name != "A" && name != "R" {
		name = workflows[name].Apply(t)
	}
	return name == "A"
}

func part1() {
	workflows, tools := parse("input.txt")
	var accepted []Tool
	for _, t := range tools {
		if isAccepted(workflows, t) {
			accepted = append(accepted, t)
		}
	}
	sum := 0
	for _, t := range accepted {
		sum += t.rating()
	}
	fmt.Printf("Part 1: %v\n", sum)
}

func part2() {
	workflows, _ := parse("input.txt")
	count := 0
	for x := 1; x <= 4000; x++ {
		for m := 1; m <= 4000; m++ {
			for a := 1; a <= 4000; a++ {
				for s := 1; s <= 4000; s++ {
					t := Tool{x, m, a, s}
					if isAccepted(workflows, t) {
						count++
					}
				}
			}
		}
	}
	fmt.Printf("Part 2: %v\n", count)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
