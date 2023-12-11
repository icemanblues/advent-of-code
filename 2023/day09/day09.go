package main

import (
	"fmt"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "09"
	dayTitle = "Mirage Maintenance"
)

func extrapolate(seq []int) int {
	deriv := make([]int, len(seq)-1, len(seq)-1)
	isAllZero := true
	for i := 1; i < len(seq); i++ {
		v := seq[i] - seq[i-1]
		deriv[i-1] = v
		isAllZero = isAllZero && v == 0
	}

	if isAllZero {
		return seq[len(seq)-1]
	}

	d := extrapolate(deriv)
	return seq[len(seq)-1] + d
}

func Reverse(a []int) []int {
	b := make([]int, 0, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		b = append(b, a[i])
	}
	return b
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	input, _ := util.ReadIntGrid("input.txt", " ")
	sum, revSum := 0, 0
	for _, seq := range input {
		sum += extrapolate(seq)
		rev := Reverse(seq)
		revSum += extrapolate(rev)
	}
	fmt.Printf("Part 1: %v\n", sum)
	fmt.Printf("Part 2: %v\n", revSum)
}
