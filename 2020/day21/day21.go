package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/icemanblues/advent-of-code/2020/util"
)

const (
	dayNum   = "21"
	dayTitle = "Allergen Assessment"
)

type Dish struct {
	ingredients map[string]struct{}
	allergens   map[string]struct{}
}

type StringSet map[string]struct{}

func copy(s StringSet) StringSet {
	myIngredients := make(StringSet)
	for ingedient := range s {
		myIngredients[ingedient] = struct{}{}
	}
	return myIngredients
}

func intersection(s1, s2 StringSet) StringSet {
	inter := make(StringSet)
	for key := range s1 {
		if _, ok := s2[key]; ok {
			inter[key] = struct{}{}
		}
	}
	return inter
}

func union(s1, s2 StringSet) StringSet {
	u := make(StringSet)
	for key := range s1 {
		u[key] = struct{}{}
	}
	for key := range s2 {
		u[key] = struct{}{}
	}
	return u
}

func difference(s1, s2 StringSet) StringSet {
	d := copy(s1)
	for key := range s1 {
		if _, ok := s2[key]; ok {
			delete(d, key)
		}
	}
	return d
}

// ParseIngredients reads input and returns all the dishes, all ingredients, all allergens
func ParseIngredients(lines []string) ([]Dish, StringSet, StringSet) {
	dishes := make([]Dish, 0, len(lines))
	allAllergens := make(StringSet)
	allIngredients := make(StringSet)

	for _, line := range lines {
		ingredientsSet := make(StringSet)
		allergensSet := make(StringSet)

		parts := strings.Split(line, " (contains ")

		ingredients := strings.Split(parts[0], " ")
		for _, e := range ingredients {
			ingredientsSet[e] = struct{}{}
			allIngredients[e] = struct{}{}
		}

		allergens := strings.ReplaceAll(parts[1], ")", "")
		allergensSlice := strings.Split(allergens, ", ")
		for _, e := range allergensSlice {
			allergensSet[e] = struct{}{}
			allAllergens[e] = struct{}{}
		}
		dishes = append(dishes, Dish{ingredientsSet, allergensSet})
	}

	return dishes, allIngredients, allAllergens
}

func FindSafe(dishes []Dish, allIngredients, allAllergens StringSet) StringSet {
	badIngredients := make(StringSet)
	for allergen := range allAllergens {
		myAllIngredients := copy(allIngredients)
		for _, dish := range dishes {
			if _, ok := dish.allergens[allergen]; !ok {
				continue
			}
			myAllIngredients = intersection(myAllIngredients, dish.ingredients)
		}
		badIngredients = union(badIngredients, myAllIngredients)
	}

	safe := difference(allIngredients, badIngredients)
	return safe
}

func NumSafeInDishes(dishes []Dish, allIngredients, allAllergens StringSet) int {
	safe := FindSafe(dishes, allIngredients, allAllergens)
	sum := 0
	for _, dish := range dishes {
		for safeIngredient := range safe {
			if _, ok := dish.ingredients[safeIngredient]; ok {
				sum++
			}
		}
	}
	return sum
}

func allergenToIngredientMap(dishes []Dish, allIngredients, allAllergens, safe StringSet) map[string]string {
	for _, dish := range dishes {
		dish.ingredients = difference(dish.ingredients, safe)
	}

	allergenToIngredient := make(map[string]string)
	for len(allergenToIngredient) != len(allAllergens) {
		for allergen := range allAllergens {
			// already solved this one
			if _, ok := allergenToIngredient[allergen]; ok {
				continue
			}

			myAllIngredients := copy(allIngredients)
			for _, dish := range dishes {
				if _, ok := dish.allergens[allergen]; ok {
					myAllIngredients = intersection(myAllIngredients, dish.ingredients)
				}
			}
			// not able to solve for this allergen just yet
			if len(myAllIngredients) != 1 {
				continue
			}

			// loop of set, size 1 :(
			for ingred := range myAllIngredients {
				allergenToIngredient[allergen] = ingred
				for _, d := range dishes {
					if _, ok := d.ingredients[ingred]; ok {
						delete(d.ingredients, ingred)
					}
				}
			}
		}
	}
	return allergenToIngredient
}

func FindDangerList(allergenToIngredient map[string]string) string {
	type pair struct {
		allergen, ingredient string
	}
	var danger []pair
	for aller, ingred := range allergenToIngredient {
		p := pair{aller, ingred}
		danger = append(danger, p)
	}

	sort.Slice(danger, func(i, j int) bool {
		return danger[i].allergen < danger[j].allergen
	})

	dangerList := make([]string, 0, len(danger))
	for _, p := range danger {
		dangerList = append(dangerList, p.ingredient)
	}
	dl := strings.Join(dangerList, ",")
	return dl
}

func part1() {
	lines, _ := util.ReadInput("input.txt")
	dishes, allIngredients, allAllergens := ParseIngredients(lines)
	fmt.Printf("Part 1: %v\n", NumSafeInDishes(dishes, allIngredients, allAllergens))
}

func part2() {
	lines, _ := util.ReadInput("input.txt")
	dishes, allIngredients, allAllergens := ParseIngredients(lines)
	safe := FindSafe(dishes, allIngredients, allAllergens)
	a2i := allergenToIngredientMap(dishes, allIngredients, allAllergens, safe)
	fmt.Printf("Part 2: %v\n", FindDangerList(a2i))
}

func main() {
	fmt.Printf("Day %v: %v\n", dayNum, dayTitle)
	part1()
	part2()
}
