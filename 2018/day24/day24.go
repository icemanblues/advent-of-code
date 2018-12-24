package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ImmuneSystem .
const ImmuneSystem = "Immune System:"

// Infection .
const Infection = "Infection:"

func readInput(filename string, boost int) (immune, infection []*Group) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("something bad with file: %v\n", err)
		return nil, nil
	}
	defer file.Close()

	// var immune []*Group
	// var infection []*Group

	var currUnits *[]*Group
	team := ImmuneSystem

	id := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}
		if line == ImmuneSystem {
			currUnits = &immune
			team = ImmuneSystem
			id = 1
			continue
		}
		if line == Infection {
			currUnits = &infection
			team = Infection
			id = 1
			continue
		}

		// 17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2
		// 989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3

		words := strings.Split(line, " ")
		units, err := strconv.Atoi(words[0])
		if err != nil {
			fmt.Printf("Could not parse units: %v\n", words[0])
		}
		hp, err := strconv.Atoi(words[4])
		if err != nil {
			fmt.Printf("Could not parse HP: %v\n", words[4])
		}
		initiative, err := strconv.Atoi(words[len(words)-1])
		if err != nil {
			fmt.Printf("Could not parse Initiative: %v\n", words[len(words)-1])
		}

		// immune
		// weakness
		weakness := make(map[string]struct{})
		immune := make(map[string]struct{})
		weakImmune := strings.Split(line, "(")
		if len(weakImmune) > 1 {
			weakImmune = strings.Split(weakImmune[1], ")")
			weakImmune = strings.Split(weakImmune[0], "; ")
			for _, wi := range weakImmune {
				w := strings.Split(wi, " ")

				m := &weakness
				if w[0] == "immune" {
					m = &immune
				}

				for i := 2; i < len(w); i++ {
					element := strings.Replace(w[i], ",", "", -1)
					(*m)[element] = struct{}{}
				}
			}
		}

		// attack power
		// dmg type
		ap := 0
		dmg := ""
		for i := len(words) - 1; i >= 0; i-- {
			if words[i] == "does" {
				ap, err = strconv.Atoi(words[i+1])
				if err != nil {
					fmt.Printf("Unable to parse attack power: %v\n", words[i+1])
				}
				dmg = words[i+2]
				break
			}
		}
		//apply the boost
		if team == ImmuneSystem {
			ap += boost
		}

		group := &Group{
			ID:     id,
			Team:   team,
			Units:  units,
			HP:     hp,
			MaxHP:  hp,
			Immune: immune,
			Weak:   weakness,
			AP:     ap,
			DMG:    dmg,
			Init:   initiative,
		}
		*currUnits = append(*currUnits, group)
		id++
	}

	return
}

// Group .
type Group struct {
	ID     int
	Team   string
	Units  int
	HP     int
	MaxHP  int
	Immune map[string]struct{}
	Weak   map[string]struct{}
	AP     int
	DMG    string
	Init   int
}

// EP Effective Power of the group. The total amount of damage it would deal
func (g Group) EP() int {
	return g.AP * g.Units
}

func lessGroup(a, b *Group) bool {
	if a.EP() > b.EP() {
		return true
	}
	if b.EP() > a.EP() {
		return false
	}

	// must be the same EP
	if a.Init > b.Init {
		return true
	}
	if b.Init > a.Init {
		return false
	}

	return true
}

func lessAtkGroup(a, b *Group) bool {
	if a.Init >= b.Init {
		return true
	}
	return false
}

// None if you do not have a target, return None
var None = &Group{}

func damage(a, e *Group) int {
	// fmt.Printf("Group %v is attacking %v\n", a, e)
	mult := 1
	if _, ok := e.Immune[a.DMG]; ok {
		mult = 0
	}
	if _, ok := e.Weak[a.DMG]; ok {
		mult = 2
	}

	return mult * a.EP()
}

func chooseTarget(g *Group, enemies []*Group, taken map[*Group]struct{}) *Group {
	var targets []*Group
	td := -1
	for _, e := range enemies {
		if _, ok := taken[e]; ok {
			continue
		}

		d := damage(g, e)
		if d == 0 {
			continue
		}
		if td == d {
			targets = append(targets, e)
		}
		if td < d {
			td = d
			targets = []*Group{e}
		}
	}

	if len(targets) == 0 {
		return None
	}

	sort.Slice(targets, func(i, j int) bool {
		return lessGroup(targets[i], targets[j])
	})
	return targets[0]
}

func printGroup(groups []*Group) {
	for i, g := range groups {
		fmt.Printf("%2d Group %2d %s %d Init: %2d HP: %d/%d Atk: %d %s Immune: %v Weakness %v\n", i, g.ID, g.Team, g.Units, g.Init, g.HP, g.MaxHP, g.AP, g.DMG, g.Immune, g.Weak)
	}
}

func main() {
	fmt.Println("Day 24: Immune System Simulator 20XX")

	// part1("test.txt")
	// part1("input24.txt")

	part2("test.txt")
	part2("input24.txt")
}

func printState(immune, infection []*Group) {
	fmt.Printf("Immune: %d\n", len(immune))
	printGroup(immune)
	fmt.Printf("Infection: %d\n", len(infection))
	printGroup(infection)
}

func battle(immune, infection *[]*Group) (string, int) {
	// Battle until we have a winner

	for len(*immune) != 0 && len(*infection) != 0 {
		// order the groups for target selection
		var allGroups []*Group
		for _, imm := range *immune {
			allGroups = append(allGroups, imm)
		}
		for _, inf := range *infection {
			allGroups = append(allGroups, inf)
		}
		sort.Slice(allGroups, func(i, j int) bool { return lessGroup(allGroups[i], allGroups[j]) })

		// Target Selection
		atkTargets := make(map[*Group]*Group)
		taken := make(map[*Group]struct{})
		for _, g := range allGroups {
			// pick your target of the available ones
			// TODO: Need to ignore targets that have already been choosen
			// the values of the atkTargets map
			if g.Team == ImmuneSystem {
				t := chooseTarget(g, *infection, taken)
				if t != None {
					atkTargets[g] = t
					taken[t] = struct{}{}
				}
			} else {
				t := chooseTarget(g, *immune, taken)
				if t != None {
					atkTargets[g] = t
					taken[t] = struct{}{}
				}
			}
		}

		// Attack Phase
		sort.Slice(allGroups, func(i, j int) bool { return lessAtkGroup(allGroups[i], allGroups[j]) })
		for _, g := range allGroups {
			// dead units cant fight
			if g.Units <= 0 {
				continue
			}

			target, ok := atkTargets[g]
			if !ok {
				continue
			}

			d := damage(g, target)
			target.Units = target.Units - (d / target.HP)

			// if dead, remove it from the slice
			if target.Units <= 0 {
				target.Units = 0
				immIdx, infIdx := -1, -1
				for i, imm := range *immune {
					if target == imm {
						immIdx = i
					}
				}
				for i, inf := range *infection {
					if target == inf {
						infIdx = i
					}
				}
				if immIdx != -1 {
					*immune = append((*immune)[:immIdx], (*immune)[immIdx+1:]...)
				}
				if infIdx != -1 {
					*infection = append((*infection)[:infIdx], (*infection)[infIdx+1:]...)
				}
			}
		}
	}

	sumUnits := 0
	for _, g := range *immune {
		sumUnits += g.Units
	}
	for _, g := range *infection {
		sumUnits += g.Units
	}

	if len(*immune) == 0 {
		return Infection, sumUnits
	}
	return ImmuneSystem, sumUnits
}

// 10538
func part1(fn string) {
	fmt.Println("Part 1")

	immune, infection := readInput(fn, 0)
	printState(immune, infection)

	winner, sumUnits := battle(&immune, &infection)

	printState(immune, infection)
	fmt.Printf("Winner %v !!! Total units remaining: %d\n", winner, sumUnits)
}

func part2(fn string) {
	fmt.Println("Part 2")

	// there was a stalemate, so I just had to work around it
	// 27 - 35
	boost := 40
	for {
		fmt.Println(boost)
		immune, infection := readInput(fn, boost)
		winner, sumUnits := battle(&immune, &infection)
		fmt.Printf("Winner %v Boost %d !!! Total units remaining: %d\n", winner, boost, sumUnits)
		if winner == Infection {
			break
		}

		boost--
	}
}
