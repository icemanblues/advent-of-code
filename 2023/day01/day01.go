package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "01"
	dayTitle = "Trebuchet?!"
)

type Pair struct{ index, value int }

func calibration(rs []rune) (Pair, Pair) {
	var first, last Pair
	for i := 0; i < len(rs); i++ {
		a, err := strconv.Atoi(string(rs[i]))
		if err != nil {
			continue
		}
		first = Pair{i, a}
		break
	}

	for i := len(rs) - 1; i >= 0; i-- {
		a, err := strconv.Atoi(string(rs[i]))
		if err != nil {
			continue
		}
		last = Pair{i, a}
		break
	}
	return first, last
}

func part1() {
	input, _ := util.ReadRuneput("input.txt")
	sum := 0
	for _, rs := range input {
		first, last := calibration(rs)
		sum += first.value*10 + last.value
	}
	fmt.Printf("Part 1: %v\n", sum)
}

var spelledOut = map[string]int{
	"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
}

func spellCalibration(line string) (Pair, Pair) {
	min, max := calibration([]rune(line))

	for s, v := range spelledOut {
		i := strings.Index(line, s)
		if i != -1 && i < min.index {
			min.index, min.value = i, v
		}
		j := strings.LastIndex(line, s)
		if j != -1 && j > max.index {
			max.index, max.value = j, v
		}
	}

	return min, max
}

func part2() {
	input, _ := util.ReadInput("input.txt")
	sum := 0
	for _, s := range input {
		first, last := spellCalibration(s)
		sum += first.value*10 + last.value
	}
	fmt.Printf("Part 2: %v\n", sum)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
