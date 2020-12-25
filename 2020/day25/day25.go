package main

import (
	"fmt"
)

const (
	dayNum   = "25"
	dayTitle = "Combo Breaker"
)

const (
	card = 18356117
	door = 5909654
)

func transform(subject, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value *= subject
		value = value % 20201227
	}
	return value
}

func findLoopSize(target int) int {
	value := 1
	i := 0
	for target != value {
		value *= 7
		value = value % 20201227
		i++
	}
	return i
}

func part1() {
	cardLoop := findLoopSize(card)
	doorLoop := findLoopSize(door)

	var answer int
	if doorLoop < cardLoop {
		answer = transform(card, doorLoop)
	} else {
		answer = transform(door, cardLoop)
	}

	fmt.Printf("Part 1: %v\n", answer)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
}
