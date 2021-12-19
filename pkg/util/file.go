package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func MustRead(filename string) []string {
	l, err := ReadInput(filename)
	if err != nil {
		panic(err)
	}
	return l
}

func ReadIntput(filename string) ([]int, error) {
	lines, err := ReadInput(filename)
	if err != nil {
		return nil, err
	}

	nums := make([]int, 0, len(lines))
	for _, e := range lines {
		n, err := strconv.Atoi(e)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func ReadRuneput(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var runes [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		runer := []rune(scanner.Text())
		runes = append(runes, runer)
	}
	return runes, nil
}

func ReadIntLine(filename string, delim string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ints []int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	words := strings.Split(line, delim)
	for _, e := range words {
		i, err := strconv.Atoi(e)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}

func ReadIntGrid(filename string, delim string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var ints []int
		line := scanner.Text()
		words := strings.Split(line, delim)
		for _, e := range words {
			i, err := strconv.Atoi(e)
			if err != nil {
				return nil, err
			}
			ints = append(ints, i)
		}
		grid = append(grid, ints)
	}
	return grid, nil
}

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
