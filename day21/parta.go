package main

import (
	"fmt"

	"github.com/jnewmano/advent2020/input"
)

func main() {

	//	input.SetRaw(raw)
	var things = input.LoadSliceString("")

	// Each allergen is found in exactly one ingredient
	// Each ingredient contains zero or one allergen.
	// Allergens aren't always marked

	// list of ingredients and facts about them
	allergens := make(map[string]Allergen)

	var list = make([]Food, 0)
	for i, v := range things {
		f := process(i, v)
		list = append(list, f)

		for _, v := range f.Allergens {
			allergens[v] = Allergen{}
		}
	}

	known := []string{}
	for {
		for _, v := range list {
			for _, w := range list {
				// look at intersection of ingredients and allergens
				commonIngredients := remove(intersection(v.Ingredients, w.Ingredients), known)
				commonAllergens := remove(intersection(v.Allergens, w.Allergens), known)

				// if we have an empty set, continue on
				if len(commonIngredients) == 0 || len(commonAllergens) == 0 {
					continue
				}
				// if we can identify an allergen, log it and continue
				if len(commonIngredients) == 1 && len(commonAllergens) == 1 {
					allergens[commonAllergens[0]] = Allergen{Ingredient: commonIngredients[0]}
					known = append(known, commonIngredients[0], commonAllergens[0])
					continue
				}
				// if we can't identify an allergen, create a new food item combining the two
				newFood := Food{
					ID:          -1, // made up food
					Ingredients: commonIngredients,
					Allergens:   commonAllergens,
				}
				list = append(list, newFood)
			}
		}

		// see if there are any unknown allergens
		knownThemAll := true
		for _, v := range allergens {
			if v.Ingredient == "" {
				knownThemAll = false
				break
			}
		}
		if knownThemAll {
			break
		}
	}

	var items []string
	for _, v := range list {
		if v.ID == -1 {
			continue
		}

		item := remove(v.Ingredients, known)
		items = append(items, item...)
	}

	fmt.Printf("Allergen free ingredients show up %d times\n", len(items))
}
