package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return 0, 0
	}
	defer file.Close()

	//10 players; last marble is worth 1618 points
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	words := strings.Split(line, " ")
	numPlayers, _ := strconv.Atoi(words[0])
	lastMarble, _ := strconv.Atoi(words[6])

	return numPlayers, lastMarble
}

func main() {
	fmt.Println("Day 9: Marble Mania")

	// part1()
	part2()
}

func part1() {
	fmt.Println("Part 1")

	numPlayers, lastMarble := readInput("input09.txt")

	score := make([]int, numPlayers, numPlayers)
	curr := 0
	marbles := []int{0}
	turn := 1

	for turn <= lastMarble {
		player := (turn % numPlayers)

		if is23(turn) {
			score[player] += turn
			remove := modulus((curr - 7), len(marbles))
			value := marbles[remove]
			score[player] += value

			marbles = marbles[:remove+copy(marbles[remove:], marbles[remove+1:])]

			curr = remove
		} else {
			marbles, curr = insert(marbles, curr, turn)
		}

		if turn%1000 == 0 {
			fmt.Printf("turn %v end %v\n", turn, lastMarble)
		}

		// fmt.Printf("turn %v current %v currmarb %v marbles %v\n", turn, curr, marbles[curr], marbles)
		turn++
	}

	max := -1
	p := -1
	for i, e := range score {
		if e > max {
			max = e
			p = i
		}
	}
	fmt.Printf("Player %v wins! high score of %v\n", p, max)
}

func part2() {
	fmt.Println("Part 2")

	numPlayers, lastMarble := readInput("input09.txt")
	lastMarble = lastMarble * 100
	// lastMarble = 25

	score := make([]int, numPlayers, numPlayers)
	marbles := list.New()
	marbles.PushFront(0)
	curr := marbles.Front()
	turn := 1

	for turn <= lastMarble {
		player := (turn % numPlayers)

		if is23(turn) {
			score[player] += turn

			// move the pointer back 7 elements
			for i := 0; i < 7; i++ {
				if curr == marbles.Front() {
					curr = marbles.Back()
				} else {
					curr = curr.Prev()
				}
			}
			value := curr.Value.(int)
			score[player] += value

			e := curr.Next()
			marbles.Remove(curr)
			curr = e
		} else {
			if curr == marbles.Back() {
				curr = marbles.Front()
			} else {
				curr = curr.Next()
			}
			curr = marbles.InsertAfter(turn, curr)
		}

		if turn%1000 == 0 {
			fmt.Printf("turn %v end %v\n", turn, lastMarble)
		}

		// fmt.Printf("turn %v current %v currmarb %v marbles %v\n", turn, curr, marbles[curr], marbles)
		turn++
	}

	max := -1
	p := -1
	for i, e := range score {
		if e > max {
			max = e
			p = i
		}
	}
	fmt.Printf("Player %v wins! high score of %v\n", p, max)
}

func is23(i int) bool {
	return i%23 == 0
}

func insert(s []int, i int, x int) ([]int, int) {
	next := (i + 1) % len(s)

	if next == len(s)-1 {
		s = append(s, x)
		return s, len(s) - 1
	}

	at := modulus((i + 2), len(s))
	// var r []int
	// r = append(r, s[:at]...)
	// r = append(r, x)
	// r = append(r, s[at:]...)
	// return r, at

	s = append(s, 0)
	copy(s[at+1:], s[at:])
	s[at] = x
	return s, at
}

func modulus(a, b int) int {
	r := a % b
	if r < 0 {
		r = r + b
	}
	return r
}
