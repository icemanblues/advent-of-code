package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "07"
	dayTitle = "Camel Cards"
)

type CamelCards struct {
	hand string
	bid  int
}

func parseCardsBids(filename string) []CamelCards {
	input, _ := util.ReadInput(filename)
	cards := make([]CamelCards, 0, len(input))
	for _, line := range input {
		cardsbids := strings.Split(line, " ")
		cards = append(cards, CamelCards{cardsbids[0], util.MustAtoi(cardsbids[1])})
	}
	return cards
}

type HandType int

const (
	FiveKind  HandType = 6
	FourKind           = 5
	FullHouse          = 4
	ThreeKind          = 3
	TwoPair            = 2
	OnePair            = 1
	HighCard           = 0
)

type CountFunc func(string) map[rune]int

func handCounts(hand string) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range hand {
		counts[r]++
	}
	return counts
}

func handType(counts map[rune]int) HandType {
	switch len(counts) {
	case 1:
		return FiveKind
	case 2: // 4 of a kind or a full house
		for _, v := range counts {
			if v == 4 {
				return FourKind
			}
		}
		return FullHouse
	case 3: // 2 pair or 3 of a kind
		for _, v := range counts {
			if v == 3 {
				return ThreeKind
			}
		}
		return TwoPair
	case 4:
		return OnePair
	case 5:
		return HighCard
	default:
		fmt.Printf("Error: Unknown hand type %v\n", counts)
		return HighCard
	}
}

var strength = map[rune]int{
	'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10, '9': 9, '8': 8, '7': 7, '6': 6,
	'5': 5, '4': 4, '3': 3, '2': 2,
}

func CardCompare(a, b CamelCards, handFunc CountFunc, str map[rune]int) bool {
	ta, tb := handType(handFunc(a.hand)), handType(handFunc(b.hand))
	if ta < tb {
		return true
	} else if ta > tb {
		return false
	}

	brunes := []rune(b.hand)
	for i, ar := range a.hand {
		br := brunes[i]
		as, bs := str[ar], str[br]
		if as < bs {
			return true
		} else if as > bs {
			return false
		}
	}
	return false
}

func play(filename string, count CountFunc, str map[rune]int) int {
	cards := parseCardsBids(filename)
	sort.Slice(cards, func(i, j int) bool {
		return CardCompare(cards[i], cards[j], count, str)
	})

	sum := 0
	for i, card := range cards {
		sum += (i + 1) * card.bid
	}
	return sum
}

func part1() {
	fmt.Printf("Part 1: %v\n", play("input.txt", handCounts, strength))
}

var strJoker = map[rune]int{
	'A': 14, 'K': 13, 'Q': 12, 'T': 10, '9': 9, '8': 8, '7': 7, '6': 6,
	'5': 5, '4': 4, '3': 3, '2': 2, 'J': 1,
}

func jokerCounts(hand string) map[rune]int {
	counts := handCounts(hand)
	joker := counts['J']
	delete(counts, 'J')
	mk, mv := 'J', 0
	for k, v := range counts {
		if v > mv {
			mk = k
			mv = v
		}
	}
	counts[mk] += joker
	return counts
}

func part2() {
	fmt.Printf("Part 2: %v\n", play("input.txt", jokerCounts, strJoker))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
