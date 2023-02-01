package main

import (
	"fmt"
	"math"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "25"
	dayTitle = "Full of Hot Air"
)

var snafus map[rune]int = map[rune]int{
	'2': 2, '1': 1, '0': 0, '-': -1, '=': -2,
}

func FivePow(b int) int {
	bf := float64(b)
	cf := math.Pow(5, bf)
	return int(cf)
}

func FiveLog(b int) int {
	c := math.Log(float64(b)) / math.Log(5)
	return int(math.Ceil(c))
}

func FromSnafu(snafu string) int {
	num := 0
	for i, r := range []rune(snafu) {
		expo := len(snafu) - 1 - i
		coeff := snafus[r]
		num += coeff * FivePow(expo)
	}
	return num
}

func ToSnafu(num int) string {
	l := FiveLog(num)
	snafu := make([]rune, 0, l)
	dec := 0

	for i := 0; i <= l; i++ {
		// find the value that is the closest, min distance
		minDist := num
		minRune := '0'
		minDec := 0
		for r, v := range snafus {
			n := v*FivePow(l-i) + dec
			dist := util.Abs(num - n)
			if dist < minDist {
				minDist = dist
				minRune = r
				minDec = n
			}
		}
		snafu = append(snafu, minRune)
		dec = minDec
	}

	if snafu[0] == '0' {
		snafu = snafu[1:]
	}
	return string(snafu)
}

func part1() {
	sum := 0
	input, _ := util.ReadInput("input.txt")
	for _, line := range input {
		sum += FromSnafu(line)
	}
	fmt.Printf("sum: %v\n", sum)
	fmt.Printf("Part 1: %v\n", ToSnafu(sum))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
}
