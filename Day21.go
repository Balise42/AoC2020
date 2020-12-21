package main

import (
	"fmt"
	"sort"
	"strings"
)

func (day *Day) Day21a() {
	res, _ := ComputeDay21a(day.input)
	fmt.Printf("Part 1: %d\n", res)
}

func (day *Day) Day21b() {
	fmt.Printf("Part 2: %s\n", ComputeDay21b(day.input))
}

type Day21 struct {
	ingredients []string
	allergens []string
}

func (d Day21) hasAllergen(allergen string) bool {
	for _, cand := range d.allergens {
		if cand == allergen {
			return true
		}
	}
	return false
}

func (d Day21) hasIngredient(ingredient string) bool {
	for _, cand := range d.ingredients {
		if cand == ingredient {
			return true
		}
	}
	return false
}

var ingredients map[string]bool
var allergens map[string]bool
var lines []Day21

func ComputeDay21a(input string) (int, []string) {
	ingredients = make(map[string]bool)
	allergens = make(map[string]bool)
	lines = make([]Day21, 0)

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " (contains ")
		parsedLine := Day21{ make([]string, 0), make([]string, 0)}
		for _, part := range strings.Split(parts[0], " ") {
			ingredients[part] = true
			parsedLine.ingredients = append(parsedLine.ingredients, part)
		}
		for _, part := range strings.Split(strings.TrimSuffix(parts[1], ")"), ", ") {
			allergens[part] = true
			parsedLine.allergens = append(parsedLine.allergens, part)
		}
		lines = append(lines, parsedLine)
	}

	canBeContainedIn := computeCanBeContainedIn()

	excluded := make([]string, 0)

	for ingredient := range ingredients {
		found := false
		for _, ings := range canBeContainedIn {
			for _, ing := range ings {
				if ing == ingredient {
					found = true
					break
				}
			}
		}
		if !found {
			excluded = append(excluded, ingredient)
		}
	}

	count := 0
	for _, line := range lines {
		for _, excl := range excluded {
			if line.hasIngredient(excl) {
				count++
			}
		}
	}

	return count, excluded
}

func computeCanBeContainedIn() map[string][]string {
	canBeContainedIn := make(map[string][]string)
	// first we make a full list
	for allergen := range allergens {
		canBeContainedIn[allergen] = make([]string, 0)
		for ingredient := range ingredients {
			canBeContainedIn[allergen] = append(canBeContainedIn[allergen], ingredient)
		}
	}
	//then we cull
	for allergen, ings := range canBeContainedIn {
		for _, line := range lines {
			if !line.hasAllergen(allergen){
				continue
			}
			for i := len(ings) - 1; i>= 0; i-- {
				if !line.hasIngredient(ings[i]) {
					ings = append(ings[:i], ings[i+1:]...)
				}
			}
			canBeContainedIn[allergen] = ings
		}
	}
	return canBeContainedIn
}

func ComputeDay21b(input string) string {
	_, excluded := ComputeDay21a(input)
	for k, line := range lines {
		for i := len(line.ingredients) - 1; i >= 0; i-- {
			for _, excl := range excluded {
				if line.ingredients[i] == excl {
					line.ingredients = append(line.ingredients[:i], line.ingredients[i+1:]...)
					break
				}
			}
		}
		lines[k] = line
	}

	canBeContainedIn := computeCanBeContainedIn()
	matched := make(map[string]string)

	for len(canBeContainedIn) > 0 {
		for all, ings := range canBeContainedIn {
			if len(ings) == 1 {
				delete(canBeContainedIn, all)
				matched[all] = ings[0]

				for all2, ings2 := range canBeContainedIn {
					for i, ing := range ings2 {
						if ing == ings[0] {
							canBeContainedIn[all2] = append(ings2[:i], ings2[i+1:]...)
							break
						}
					}
				}
			}
		}
	}

	allergenList := make([]string, 0, len(matched))
	for k := range matched {
		allergenList = append(allergenList, k)
	}
	sort.Strings(allergenList)

	res := matched[allergenList[0]]
	for _, all := range allergenList[1:] {
		res = res + "," + matched[all]
	}
	return res
}