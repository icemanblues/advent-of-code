package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	dayNum   = "10"
	dayTitle = "Elves Look, Elves Say"
)

const input = "3113322113"

func LookAndSay(s string) string {
	b := strings.Builder{}
	copies, digit := 0, []rune(s)[0]
	for _, r := range s {
		if r == digit {
			copies++
		} else {
			b.WriteString(strconv.Itoa(copies))
			b.WriteRune(digit)
			digit = r
			copies = 1
		}
	}
	b.WriteString(strconv.Itoa(copies))
	b.WriteRune(digit)
	return b.String()
}

func process(start string, num int) string {
	curr := start
	for i := 0; i < num; i++ {
		curr = LookAndSay(curr)
	}
	return curr
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	p := process(input, 40)
	fmt.Printf("Part 1: %v\n", len(p))
	q := process(p, 10)
	fmt.Printf("Part 2: %v\n", len(q))
}
