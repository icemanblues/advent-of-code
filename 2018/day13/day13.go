package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func readInput(filename string) (Carts, []string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return nil, nil
	}
	defer file.Close()

	lineNum := 0
	var lines []string
	var carts Carts
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for i, r := range line {
			if r == '>' {
				c := &Cart{
					x:         i,
					y:         lineNum,
					direction: right,
					turn:      0,
				}
				carts = append(carts, c)
			} else if r == '<' {
				c := &Cart{
					x:         i,
					y:         lineNum,
					direction: left,
					turn:      0,
				}
				carts = append(carts, c)
			} else if r == '^' {
				c := &Cart{
					x:         i,
					y:         lineNum,
					direction: up,
					turn:      0,
				}
				carts = append(carts, c)
			} else if r == 'v' {
				c := &Cart{
					x:         i,
					y:         lineNum,
					direction: down,
					turn:      0,
				}
				carts = append(carts, c)
			}
		}

		// replace the carts with the proper road symbol
		line = strings.Replace(line, "<", "-", -1)
		line = strings.Replace(line, ">", "-", -1)
		line = strings.Replace(line, "^", "|", -1)
		line = strings.Replace(line, "v", "|", -1)
		lines = append(lines, line)

		lineNum++
	}
	return carts, lines
}

const (
	up    = 1
	down  = 2
	left  = 3
	right = 4
)

const (
	turnLeft  = 0
	straight  = 1
	turnRight = 2
)

type Cart struct {
	x         int
	y         int
	direction int
	turn      int
}

type Carts []*Cart

func (c Carts) Len() int { return len(c) }
func (c Carts) Less(i, j int) bool {
	if c[i].y < c[j].y {
		return true
	} else if c[i].y > c[j].y {
		return false
	}

	if c[i].x < c[j].x {
		return true
	}
	// must be if c[i].x >= c[j].x
	return false
}
func (c Carts) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func main() {
	fmt.Println("Day 13: Mine Cart Madness")

	// part1()
	part2()
}

func printGrid(g []string) {
	for _, s := range g {
		fmt.Println(s)
	}
}

func printCarts(c Carts) {
	fmt.Println(len(c))
	for _, e := range c {
		fmt.Println(*e)
	}
}

func move(c *Cart) {
	if c.direction == up {
		c.y--
	} else if c.direction == down {
		c.y++
	} else if c.direction == right {
		c.x++
	} else if c.direction == left {
		c.x--
	}
}

func intersection(direction, turn int) int {
	if turn == straight {
		return direction
	}

	if direction == up && turn == turnLeft {
		return left
	}
	if direction == up && turn == turnRight {
		return right
	}

	if direction == down && turn == turnLeft {
		return right
	}
	if direction == down && turn == turnRight {
		return left
	}

	if direction == left && turn == turnLeft {
		return down
	}
	if direction == left && turn == turnRight {
		return up
	}

	if direction == right && turn == turnLeft {
		return up
	}
	if direction == right && turn == turnRight {
		return down
	}

	fmt.Printf("Never should have reach here. direction %v turn %v\n", direction, turn)
	return -1
}

func turn(c *Cart, tile rune) {
	if tile == '+' {
		td := c.turn % 3
		c.direction = intersection(c.direction, td)
		c.turn++
	} else if tile == '\\' && c.direction == right {
		c.direction = down
	} else if tile == '\\' && c.direction == up {
		c.direction = left
	} else if tile == '\\' && c.direction == left {
		c.direction = up
	} else if tile == '\\' && c.direction == down {
		c.direction = right
	} else if tile == '/' && c.direction == right {
		c.direction = up
	} else if tile == '/' && c.direction == up {
		c.direction = right
	} else if tile == '/' && c.direction == left {
		c.direction = down
	} else if tile == '/' && c.direction == down {
		c.direction = left
	}

}

func part1() {
	fmt.Println("Part 1")

	carts, grid := readInput("input13.txt")
	printCarts(carts)
	printGrid(grid)

	tick := 0
	crash := false
	for !crash {
		sort.Sort(carts)
		for i, c := range carts {
			//advance the carts
			tile := rune(grid[c.y][c.x])
			turn(c, tile)
			move(c)

			// check for collosions
			for j, k := range carts {
				// cant collide into yourself
				if i == j {
					continue
				}

				crash = crash || (c.x == k.x && c.y == k.y)
				if crash {
					fmt.Println("C R A S H E D !!!", c.x, c.y)
					return
				}
			}
		}

		tick++
	}
}

func part2() {
	fmt.Println("Part 2")

	carts, grid := readInput("input13.txt")
	printCarts(carts)
	printGrid(grid)

	tick := 0
	for len(carts) > 1 {
		sort.Sort(carts)
		var kaput []int
		for i, c := range carts {
			//advance the carts
			tile := rune(grid[c.y][c.x])
			turn(c, tile)
			move(c)

			// check for collosions
			for j, k := range carts {
				// cant collide into yourself
				if i == j {
					continue
				}

				crash := c.x == k.x && c.y == k.y
				if crash {
					fmt.Println("C R A S H E D !!!", c.x, c.y)
					kaput = append(kaput, i)
					kaput = append(kaput, j)
				}
			}
		}

		// remove crashed carts
		sort.Sort(sort.Reverse(sort.IntSlice(kaput)))
		// fmt.Printf("deleting these carts %v\n", kaput)
		for _, i := range kaput {
			// fmt.Printf("removing idx %v from carts %v\n", i, len(carts))
			carts = append(carts[:i], carts[i+1:]...)
		}
		tick++
	}

	printCarts(carts)
}
