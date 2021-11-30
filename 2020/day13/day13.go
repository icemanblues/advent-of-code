package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

func part1() {
	lines, _ := util.ReadInput("input.txt")
	timestamp, _ := strconv.Atoi(lines[0])
	var busIDs []int
	for _, s := range strings.Split(lines[1], ",") {
		if s != "x" {
			n, _ := strconv.Atoi(s)
			busIDs = append(busIDs, n)
		}
	}

	t := timestamp
	answer := 0
busport:
	for true {
		for _, bus := range busIDs {
			if t%bus == 0 { // found it
				answer = (t - timestamp) * bus
				break busport
			}
		}
		t++
	}

	fmt.Printf("Part 1: %v\n", answer)
}

func part2() {
	lines, _ := util.ReadInput("input.txt")
	busOrdinal := make(map[int]int)
	for i, s := range strings.Split(lines[1], ",") {
		if s != "x" {
			n, _ := strconv.Atoi(s)
			busOrdinal[n] = i
		}
	}

	n := make([]*big.Int, 0, len(busOrdinal)) // buses
	a := make([]*big.Int, 0, len(busOrdinal)) // ord (negative)
	for bus, ord := range busOrdinal {
		n = append(n, big.NewInt(int64(bus)))
		a = append(a, big.NewInt(int64(-ord)))
	}
	answer, _ := crt(a, n)

	fmt.Printf("Part 2: %v\n", answer)
}

func main() {
	fmt.Println("Day 13: Shuttle Search")
	part1()
	part2()
}

// The Chinese Remainder Theorem code is lifted from
// https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
var one = big.NewInt(1)

func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}
