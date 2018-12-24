package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const ImmuneSystem = "Immune System:"
const Infection = "Infection:"

func readInput(filename string) (immune, infection []*Group) {
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

func printGroup(groups []*Group) {
	for i, g := range groups {
		fmt.Printf("%2d Group %2d %s %d Init: %2d HP: %d/%d Atk: %d %s Immune: %v Weakness %v\n", i, g.ID, g.Team, g.Units, g.Init, g.HP, g.MaxHP, g.AP, g.DMG, g.Immune, g.Weak)
	}
}

func main() {
	fmt.Println("Day 24: Immune System Simulator 20XX")

	part1("test.txt")
	part2()
}

func part1(fn string) {
	fmt.Println("Part 1")

	immune, infection := readInput(fn)
	fmt.Println("Immune")
	printGroup(immune)
	fmt.Println("Infection")
	printGroup(infection)
}

func part2() {
	fmt.Println("Part 2")
}
