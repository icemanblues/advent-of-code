package main

import (
	"fmt"
)

const (
	dayNum   = "21"
	dayTitle = "Dirac Dice"
)

type Dice interface {
	Roll() int
	Multi(int) int
}

type DeterministicDice struct {
	seed int
}

func (d *DeterministicDice) Roll() int {
	d.seed++
	return d.seed
}

func (d *DeterministicDice) Multi(a int) int {
	sum := 0
	for i := 0; i < a; i++ {
		sum += d.Roll()
	}
	return sum
}

func move(p, n int) int {
	for i := 0; i < n; i++ {
		p++
		if p > 10 {
			p = 1
		}
	}
	return p
}

func play(p1, p2 int, dice Dice) int {
	s1, s2 := 0, 0
	t := 0
	ptr := true
	for s1 < 1000 && s2 < 1000 {
		t++
		r := dice.Multi(3)
		if ptr {
			p1 = move(p1, r)
			s1 += p1
		} else {
			p2 = move(p2, r)
			s2 += p2
		}
		ptr = !ptr
	}

	dieRolls := t * 3
	losingScore := s1
	if s2 < s1 {
		losingScore = s2
	}
	return dieRolls * losingScore
}

func part1() {
	p1, p2 := 9, 4 // input.txt
	d := DeterministicDice{}
	s := play(p1, p2, &d)
	fmt.Printf("Part 1: %v\n", s)
}

type Game struct {
	p1, p2 int
	s1, s2 int
	ptr    bool
}

// dice roll -> num of universes
var quantum = map[int]int64{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

// this is do all of the iteration and build the cache
func dynamic(game Game, memo map[Game]int64) int64 {
	if winner, ok := memo[game]; ok {
		return winner
	}

	// check if game is complete
	if game.s1 >= 21 {
		memo[game] = 1
		return 1
	}
	if game.s2 >= 21 {
		memo[game] = 0
		return 0
	}

	var sum int64 = 0
	for roll, universes := range quantum {
		if game.ptr {
			newP1 := move(game.p1, roll)
			newGame := Game{newP1, game.p2, game.s1 + newP1, game.s2, !game.ptr}
			value := universes * dynamic(newGame, memo)
			sum += value
		} else {
			newP2 := move(game.p2, roll)
			newGame := Game{game.p1, newP2, game.s1, game.s2 + newP2, !game.ptr}
			value := universes * dynamic(newGame, memo)
			sum += value
		}
	}

	memo[game] = sum
	return sum
}

func part2() {
	in1, in2 := 9, 4 // input.txt
	game := Game{in1, in2, 0, 0, true}
	memo := make(map[Game]int64)
	universes := dynamic(game, memo)
	fmt.Printf("Part 2: %v\n", universes)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
