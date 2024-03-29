package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadInput returns its line in the file as a string
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

// ReadIntput returns each line in the file as an int
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

// ReadRuneput returns each line in the file as a slice of runes []rune
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

// ReadIntGrid returns a 2D grid of ints from the file
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

// ReadSparseGrid returns sparse grid (set) containing the points of a specific rune
func ReadSparseGrid(filename string, symbol rune) (map[Point]struct{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sparse := make(map[Point]struct{})
	scanner, y := bufio.NewScanner(file), 0
	for scanner.Scan() {
		for x, r := range []rune(scanner.Text()) {
			if r == symbol {
				sparse[Point{x, y}] = struct{}{}
			}
		}
		y++
	}
	return sparse, nil
}

// Read3D returns the lines of the file as Point3D
func Read3D(filename string, delim string) ([]Point3D, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var points []Point3D
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, delim)
		ints := make([]int, 0, 3)
		for _, e := range words {
			i, err := strconv.Atoi(e)
			if err != nil {
				return nil, err
			}
			ints = append(ints, i)
		}
		points = append(points, Point3D{ints[0], ints[1], ints[2]})
	}
	return points, nil
}

func MustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func SliceAtoI(ss []string) []int {
	ints := make([]int, len(ss), len(ss))
	for i, s := range ss {
		ints[i] = MustAtoi(s)
	}
	return ints
}
