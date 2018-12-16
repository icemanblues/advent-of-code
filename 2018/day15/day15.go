package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readInput(filename string) ([][]rune, Units, Units) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return nil, nil, nil
	}
	defer file.Close()

	var elves Units
	var goblins Units

	cave := make(map[Point]bool)
	var grid [][]rune

	lineNum := 0
	uniqueID := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var l []rune
		for i, r := range line {
			p := Point{i, lineNum}
			if r == 'G' {
				g := NewGoblin(uniqueID, i, lineNum)
				goblins = append(goblins, g)
				l = append(l, intToRune(uniqueID))
				uniqueID++
				continue
			}
			if r == 'E' {
				e := NewElf(uniqueID, i, lineNum)
				elves = append(elves, e)
				l = append(l, intToRune(uniqueID))
				uniqueID++
				continue
			}

			if r == '#' {
				cave[p] = false
				l = append(l, '#')
			} else {
				cave[p] = true
				l = append(l, '.')
			}
		}

		grid = append(grid, l)
		lineNum++
	}
	return grid, elves, goblins
}

func intToRune(i int) rune {
	return rune(strconv.Itoa(i)[0])
}

type Unit struct {
	ID   int
	Race rune
	loc  Point
	Ap   int
	Hp   int
}

func NewBaseUnit() *Unit {
	return &Unit{
		Ap: 3,
		Hp: 200,
	}
}

func NewElf(id, x, y int) *Unit {
	elf := NewBaseUnit()
	elf.Race = 'E'

	elf.ID = id
	elf.loc = Point{x, y}
	return elf
}

func NewGoblin(id, x, y int) *Unit {
	gob := NewBaseUnit()
	gob.Race = 'G'

	gob.ID = id
	gob.loc = Point{x, y}
	return gob
}

type Units []*Unit

func (u Units) Len() int           { return len(u) }
func (u Units) Less(i, j int) bool { return ComparePoint(u[i].loc, u[j].loc) == -1 }
func (u Units) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

type Point struct {
	X int
	Y int
}

// reading order comparison
func ComparePoint(p1, p2 Point) int {
	if p1.Y < p2.Y {
		return -1
	}
	if p1.Y > p2.Y {
		return 1
	}
	// Y's are equal
	if p1.X < p2.X {
		return -1
	}
	if p1.X >= p2.X {
		return 1
	}

	return 0
}

type Points []Point

func (ps Points) Len() int           { return len(ps) }
func (ps Points) Less(i, j int) bool { return ComparePoint(ps[i], ps[j]) == -1 }
func (ps Points) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }

func main() {
	fmt.Println("Day 15: Beverage Bandits")

	// tests := []string{"test1.txt", "test2.txt", "test3.txt", "test4.txt", "test5.txt", "test6.txt"}
	tests := []string{"test1.txt"}
	for _, t := range tests {
		part1(t)
	}

	// part2()
}

func printState(cave [][]rune, elves []*Unit, goblins []*Unit, verbose bool) {
	fmt.Printf("Elves: %v\n", len(elves))
	if verbose {
		printUnits(elves)
	}

	fmt.Printf("Goblins: %v\n", len(goblins))
	if verbose {
		printUnits(goblins)
	}

	printCave(cave)
}

func printCave(cave [][]rune) {
	for _, s := range cave {
		for _, r := range s {
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
}

func printUnits(units []*Unit) {
	for _, u := range units {
		fmt.Printf("ID: %2d, race: %c, (%v,%v) AP: %d HP: %d\n", u.ID, u.Race, u.loc.X, u.loc.Y, u.Ap, u.Hp)
	}
}

// returns all of the adjacent points that are not walls
func adjNoWall(p Point, grid [][]rune) []Point {
	var points []Point

	// up
	if p.Y-1 >= 0 && grid[p.Y-1][p.X] != '#' {
		points = append(points, Point{p.X, p.Y - 1})
	}

	// left
	if p.X-1 >= 0 && grid[p.Y][p.X-1] != '#' {
		points = append(points, Point{p.X - 1, p.Y})
	}

	// right
	if p.X+1 < len(grid[0]) && grid[p.Y][p.X+1] != '#' {
		points = append(points, Point{p.X + 1, p.Y})
	}

	// down
	if p.Y+1 < len(grid) && grid[p.Y+1][p.X] != '#' {
		points = append(points, Point{p.X, p.Y + 1})
	}
	return points
}

// returns all of the adjacent units/points that are enemies
func adjEnemies(u *Unit, enemies Units, grid [][]rune) Units {
	s := u.loc
	noWalls := adjNoWall(s, grid)

	var points []Point
	for _, p := range noWalls {
		if grid[p.Y][p.X] != '.' {
			points = append(points, p)
		}
	}

	var atkEnemies Units
	for _, p := range points {
		for _, e := range enemies {
			if e.loc.X == p.X && e.loc.Y == p.Y {
				atkEnemies = append(atkEnemies, e)
			}
		}
	}

	return atkEnemies
}

// returns all of the adjacent points that are not walls and not units (aka open space only)
func adj(p Point, grid [][]rune) []Point {
	noWalls := adjNoWall(p, grid)

	var points []Point
	for _, p := range noWalls {
		if grid[p.Y][p.X] == '.' {
			points = append(points, p)
		}
	}

	return points
}

type PointDist struct {
	point Point
	dist  int
	step  Point
	path  []Point
}
type PointDists []PointDist

func (pds PointDists) Len() int { return len(pds) }
func (pds PointDists) Less(i, j int) bool {
	if pds[i].dist < pds[j].dist {
		return true
	}
	if pds[i].dist > pds[j].dist {
		return false
	}

	// must be equal dist, so compare via reading order
	return ComparePoint(pds[i].step, pds[j].step) == -1

}
func (pds PointDists) Swap(i, j int) { pds[i], pds[j] = pds[j], pds[i] }

// does breadth first search to find the point
func isReachable(u Unit, target Point, grid [][]rune) (bool, []PointDist) {
	empty := Point{-1, -1}
	start := u.loc
	startDist := &PointDist{start, 0, empty, nil}
	visited := make(map[Point]struct{}) // set
	queue := []*PointDist{startDist}

	var paths []PointDist

	for len(queue) != 0 {
		// dequeue
		p := queue[0]
		queue = queue[1:]

		visited[p.point] = struct{}{}

		next := adj(p.point, grid)
		for _, n := range next {
			if _, ok := visited[n]; !ok {
				pd := &PointDist{n, p.dist + 1, empty, nil}
				// increment the path
				pd.path = append(p.path, p.point)

				queue = append(queue, pd)
			}
		}

		// is p the target destination
		if p.point.X == target.X && p.point.Y == target.Y {
			p.path = append(p.path, target)
			p.step = p.path[1]

			// return true, *p
			paths = append(paths, *p)
		}
	}

	return len(paths) != 0, paths
}

// MOVE needs to update two places (the unit and the grid)
func move(u *Unit, enemies Units, cave [][]rune) {
	// determine if the unit is adj to an enemy
	if adjEnemies := adjEnemies(u, enemies, cave); len(adjEnemies) != 0 {
		// just go attack
		return
	}

	// determine open spaces adjacent to the targets
	var adjTargets []Point
	for _, e := range enemies {
		adjTargets = append(adjTargets, adj(e.loc, cave)...)
	}

	// determine if the points are reachable, and how many steps away
	var adjDist PointDists
	for _, t := range adjTargets {
		if ok, pd := isReachable(*u, t, cave); ok {
			adjDist = append(adjDist, pd...)
		}
	}

	// find the point with the shortest distance, using reading order as tie breaker
	sort.Sort(adjDist)

	if len(adjDist) != 0 {

		// okay we can move one step in that direction
		direction := adjDist[0]
		next := direction.step

		cave[u.loc.Y][u.loc.X] = '.'
		u.loc = next
		cave[u.loc.Y][u.loc.X] = intToRune(u.ID)
	}
}

// will update HP values if they are attacking
func attack(u *Unit, enemies Units, cave [][]rune) (bool, int) {
	// it should return an updated array of elves and goblins, in case one has died
	// or maybe we mark it as dead, and deal with it later (skip it)

	atkEnemies := adjEnemies(u, enemies, cave)
	if len(atkEnemies) == 0 {
		return false, -1
	}

	sort.Slice(atkEnemies, func(i, j int) bool {
		if atkEnemies[i].Hp < atkEnemies[i].Hp {
			return true
		}
		if atkEnemies[i].Hp > atkEnemies[i].Hp {
			return false
		}

		// must be the same HP totals, so use reading order
		return ComparePoint(atkEnemies[i].loc, atkEnemies[j].loc) == -1
	})

	target := atkEnemies[0]
	target.Hp -= u.Ap

	// the attack killed it, remove it from the board
	if target.Hp <= 0 {
		cave[target.loc.Y][target.loc.X] = '.'
		target.loc = Point{-1, -1}
		return true, target.ID
	}

	return false, -1
}

func isDead(u *Unit) bool {
	return u.Hp <= 0
}

func part1(fn string) {
	fmt.Println("Part 1")

	cave, elves, goblins := readInput(fn)

	round := 0
	// for len(elves) != 0 && len(goblins) != 0 {
	for round < 40 {
		// display cave status
		fmt.Printf("Round %d\n", round)
		printState(cave, elves, goblins, true)

		if len(elves) == 0 || len(goblins) == 0 {
			break
		}
		// order all units in reading order, so we can take turn
		units := append(elves, goblins...)
		sort.Sort(units)

		for _, u := range units {
			if isDead(u) {
				continue
			}

			// determine possible enemy targets
			var enemies Units
			if u.Race == 'E' {
				enemies = goblins
			} else {
				enemies = elves
			}

			if len(enemies) == 0 {
				break
			}
			sort.Sort(enemies)

			// MOVE MOVE MOVE
			move(u, enemies, cave)

			// ATTACK ATTACK ATTACK
			kill, id := attack(u, enemies, cave)
			// attack(u, enemies, cave)
			//if there is a kill, remove it from the slice
			if kill {
				gidx, eidx := -1, -1
				for i, g := range goblins {
					if g.ID == id {
						gidx = i
					}
				}
				for i, e := range elves {
					if e.ID == id {
						eidx = i
					}
				}

				if gidx != -1 {
					goblins = append(goblins[:gidx], goblins[gidx+1:]...)
				}
				if eidx != -1 {
					elves = append(elves[:eidx], elves[eidx+1:]...)
				}
			}
		}

		if len(elves) == 0 || len(goblins) == 0 {
			break
		}
		round++
	}

	sumHP := 0
	for _, g := range goblins {
		sumHP += g.Hp
	}
	for _, e := range elves {
		sumHP += e.Hp
	}

	outcome := round * sumHP
	fmt.Printf("[%v]: rounds %d total HP %d outcome %d\n", fn, round, sumHP, outcome)
}

func part2() {
	fmt.Println("Part 2")
}
