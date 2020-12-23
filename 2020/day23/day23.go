package main

import (
	"fmt"
	"math"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "23"
	dayTitle = "Title"
)

const input string = "974618352"
const test1 string = "389125467"

func ParseCups(s string) []int {
	cups := make([]int, 0, len(s))
	for _, r := range s {
		cups = append(cups, util.MustAtoi(string(r)))
	}
	return cups
}

func PlayGame(cups []int, numMoves int) []int {
	l := len(cups)
	currIdx := 0
	curr := cups[currIdx]
	min, max := cups[0], cups[0]
	for _, c := range cups {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}

	for m := 1; m <= numMoves; m++ {
		// find pickups
		pickups := make([]int, 0, 3)
		for i := 1; i <= 3; i++ {
			idx := (currIdx + i) % l
			pickups = append(pickups, cups[idx])
		}

		// find the label
		label := curr - 1
		if label < min {
			label = max
		}
		for label == pickups[0] || label == pickups[1] || label == pickups[2] {
			label--
			if label < min {
				label = max
			}
		}

		// remove the pick ups
		for _, p := range pickups {
			for i, c := range cups {
				if c == p {
					cups = append(cups[:i], cups[i+1:]...)
					break
				}
			}
		}

		// find the index of the label
		labelIdx := -1
		for i, c := range cups {
			if c == label {
				labelIdx = i
				break
			}
		}

		// add them back in by the label
		newCups := make([]int, 0, l)
		insertPoint := (labelIdx + 1) % l
		newCups = append(newCups, cups[:insertPoint]...)
		newCups = append(newCups, pickups...)
		newCups = append(newCups, cups[insertPoint:]...)
		cups = newCups

		// increment currIdx and curr
		for i, c := range cups {
			if c == curr {
				currIdx = i
			}
		}
		currIdx = (currIdx + 1) % l
		curr = cups[currIdx]
	}

	return cups
}

func score(cups []int) int {
	var labels []int
	foundOne := false
	it := 0
	score := 0
	exp := len(cups) - 2
	for len(labels) != len(cups)-1 {
		c := cups[it]
		if foundOne {
			labels = append(labels, c)
			score += c * int(math.Pow10(exp))
			exp--
		}

		if c == 1 {
			foundOne = true
		}
		it = (it + 1) % len(cups)
	}
	return score
}

type BigCup struct {
	Label int
	Next  *BigCup
}

type BigCupGame struct {
	Head               *BigCup
	Lookup             map[int]*BigCup
	MinLabel, MaxLabel int
}

func ParseBigCups(s string, size int) BigCupGame {
	cupIndexMap := make(map[int]*BigCup)
	max := -1
	var head *BigCup = nil
	var curr *BigCup = nil
	for i, r := range s {
		n := util.MustAtoi(string(r))
		if n > max {
			max = n
		}
		bigCup := &BigCup{n, nil}
		cupIndexMap[n] = bigCup

		if i == 0 {
			head = bigCup
		} else {
			curr.Next = bigCup
		}
		curr = bigCup
	}

	for i := len(s); i < size; i++ {
		max++
		bigCup := &BigCup{max, nil}
		cupIndexMap[max] = bigCup
		curr.Next = bigCup
		curr = bigCup
	}
	curr.Next = head
	return BigCupGame{head, cupIndexMap, 1, max}
}

func PlayBigGame(game BigCupGame, numMoves int) BigCupGame {
	current := game.Head

	for m := 1; m <= numMoves; m++ {
		// find the pick ups
		cup1 := current.Next
		cup2 := cup1.Next
		cup3 := cup2.Next

		// remove them from the linked list
		current.Next = cup3.Next

		// find the destination
		destLabel := current.Label - 1
		if destLabel < game.MinLabel {
			destLabel = game.MaxLabel
		}
		for destLabel == cup1.Label || destLabel == cup2.Label || destLabel == cup3.Label {
			destLabel--
			if destLabel < game.MinLabel {
				destLabel = game.MaxLabel
			}
		}

		// insert at the destination
		destCup := game.Lookup[destLabel]
		keep := destCup.Next
		destCup.Next = cup1
		cup3.Next = keep

		// determine the next current cup
		current = current.Next
	}
	return game
}

func scoreBig(game BigCupGame) int {
	for label, bigcup := range game.Lookup {
		if label == 1 {
			one := bigcup.Next
			two := one.Next
			return one.Label * two.Label
		}
	}
	panic("There is no 1")
}

func part1() {
	cups := ParseCups(input)
	cups = PlayGame(cups, 100)
	labels := score(cups)
	fmt.Printf("Part 1: %v\n", labels)
}

func part2() {
	numMoves := 10000000
	numCups := 1000000
	cups := ParseBigCups(input, numCups)
	cups = PlayBigGame(cups, numMoves)
	labels := scoreBig(cups)
	fmt.Printf("Part 2: %v\n", labels)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
