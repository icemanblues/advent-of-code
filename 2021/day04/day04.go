package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/icemanblues/advent-of-code/pkg/util"
)

const (
	dayNum   = "04"
	dayTitle = "Giant Squid"
)

type Bingo struct {
	Board [][]int
}

func parse(filename string) ([]int, []Bingo) {
	file, err := os.Open(filename)
	if err != nil {
		panic("cannot find file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var numbers []int
	scanner.Scan()
	first := scanner.Text()
	nums := strings.Split(first, ",")
	for _, n := range nums {
		numbers = append(numbers, util.MustAtoi(n))
	}

	scanner.Scan() // empty line

	var bingos []Bingo
	board := make([][]int, 0, 5)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			bingo := Bingo{board}
			bingos = append(bingos, bingo)
			board = make([][]int, 0, 5)
			continue
		}

		items := strings.Split(line, " ")
		var row []int
		for _, e := range items {
			if len(e) == 0 {
				continue
			}
			row = append(row, util.MustAtoi(e))
		}
		board = append(board, row)
	}
	bingos = append(bingos, Bingo{board})

	return numbers, bingos
}

func play(bingo Bingo, numbers []int) (turns int, score int) {
	var marked [][]bool
	for i := 0; i < 5; i++ {
		row := make([]bool, 5, 5)
		marked = append(marked, row)
	}

	// play the numbers until we win
	num := numbers[0]
	for _, n := range numbers {
		turns++
		num = n
		// search for the number n and mark it
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				if bingo.Board[x][y] == n {
					marked[x][y] = true
				}
			}
		}

		// check to see if we have a bingo
		win := false
		for i := 0; i < 5; i++ {
			hWin := marked[i][0] && marked[i][1] && marked[i][2] && marked[i][3] && marked[i][4]
			vWin := marked[0][i] && marked[1][i] && marked[2][i] && marked[3][i] && marked[4][i]
			win = win || hWin || vWin
		}
		if win {
			break
		}
	}

	// compute the score
	sum := 0
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if !marked[x][y] {
				sum += bingo.Board[x][y]
			}
		}
	}
	score = sum * num

	return turns, score
}

func part1() {
	numbers, bingos := parse("input.txt")

	minT := -1
	minTScore := -1
	for _, b := range bingos {
		t, s := play(b, numbers)
		if minT == -1 {
			minT, minTScore = t, s
		} else if t < minT {
			minT, minTScore = t, s
		}
	}
	fmt.Printf("Part 1: %v\n", minTScore)
}

func part2() {
	numbers, bingos := parse("input.txt")

	maxT := -1
	maxTScore := -1
	for _, b := range bingos {
		t, s := play(b, numbers)
		if maxT == -1 {
			maxT, maxTScore = t, s
		} else if t > maxT {
			maxT, maxTScore = t, s
		}
	}
	fmt.Printf("Part 2: %v\n", maxTScore)
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
