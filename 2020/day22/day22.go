package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "22"
	dayTitle = "Title"
)

func parseCards(lines []string) ([]int, []int) {
	var one, two []int
	isOne := true
	for _, line := range lines {
		if strings.HasPrefix(line, "Player") {
			continue
		}
		if len(line) == 0 {
			isOne = false
			continue
		}
		n, _ := strconv.Atoi(line)
		if isOne {
			one = append(one, n)
		} else {
			two = append(two, n)
		}
	}

	return one, two
}

func round(one, two []int) int {
	round := 0
	for len(one) != 0 && len(two) != 0 {
		p := one[0]
		one = one[1:]

		q := two[0]
		two = two[1:]

		if p > q {
			one = append(one, p, q)
		} else if q > p {
			two = append(two, q, p)
		} else {
			fmt.Println("draw!!")
		}
		round++
	}

	winner := two
	if len(one) > 0 {
		winner = one
	}
	return score(winner)
}

func score(deck []int) int {
	score := 0
	for i, c := range deck {
		multiplier := len(deck) - i
		score += multiplier * c
	}
	return score
}

func deck2string(deck []int) string {
	s := make([]string, 0, len(deck))
	for _, d := range deck {
		s = append(s, strconv.Itoa(d))
	}
	return strings.Join(s, ",")
}

func gameKey(p1, p2 []int) string {
	return deck2string(p1) + "+" + deck2string(p2)
}

var cnt int = 1

func recursive(one, two []int) (int, int) {
	previous := make(map[string]struct{})
	//game := cnt
	cnt++
	round := 1
	for len(one) != 0 && len(two) != 0 {
		//fmt.Printf("Game %v Round %v P1 %v \n", game, round, one)
		//fmt.Printf("Game %v Round %v P2 %v \n", game, round, two)
		// check the hands for infinite recursion
		// player 1 wins?
		state := gameKey(one, two)
		if _, ok := previous[state]; ok {
			// infinite recursion check
			// player 1 wins the game
			//fmt.Printf("infinite recursion. Game %v Round %v P1 wins!\n", game, round)
			return 1, score(one)
		}

		// add the game state
		previous[state] = struct{}{}

		p := one[0]
		one = one[1:]

		q := two[0]
		two = two[1:]

		if len(one) >= p && len(two) >= q {
			//  play a subgame to determine the winner of the round
			copyOne := make([]int, p)
			copyTwo := make([]int, q)
			copy(copyOne, one[:p])
			copy(copyTwo, two[:q])
			winner, _ := recursive(copyOne, copyTwo)
			if winner == 1 {
				one = append(one, p, q)
			} else {
				two = append(two, q, p)
			}
			round++
			continue
		}

		if p > q {
			one = append(one, p, q)
		} else if q > p {
			two = append(two, q, p)
		}
		round++
	}

	// game has ended.
	if len(one) == 0 {
		//fmt.Printf("Game %v Round %v won by P2 %v\n", game, round, score(two))
		return 2, score(two)
	}
	//fmt.Printf("Game %v Round %v won by P1 %v\n", game, round, score(one))
	return 1, score(one)
}

func part1() {
	lines, _ := util.ReadInput("input.txt")
	one, two := parseCards(lines)
	fmt.Printf("Part 1: %v\n", round(one, two))
}

func test1() {
	lines, _ := util.ReadInput("test1.txt")
	one, two := parseCards(lines)
	fmt.Println(one)
	fmt.Println(two)
	_, s := recursive(one, two)
	fmt.Printf("Test 1: %v\n", s)
}

func test2() {
	lines, _ := util.ReadInput("test2.txt")
	one, two := parseCards(lines)
	fmt.Println(one)
	fmt.Println(two)
	_, s := recursive(one, two)
	fmt.Printf("Test 2: %v\n", s)
}

func part2() {
	lines, _ := util.ReadInput("input.txt")
	one, two := parseCards(lines)
	fmt.Println(one)
	fmt.Println(two)
	_, s := recursive(one, two)
	fmt.Printf("Part 2: %v\n", s)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	test1()
	test2()
	part2()
}
