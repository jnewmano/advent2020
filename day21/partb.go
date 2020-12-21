package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/jnewmano/advent2020/input"
)

func main() {

	// input.SetRaw(raw)
	var things = input.LoadSliceString("")

	// Each allergen is found in exactly one ingredient
	// Each ingredient contains zero or one allergen.
	// Allergens aren't always marked

	allergens := make(map[string]Allergen)

	var list = make([]Food, 0)
	for i, v := range things {
		f := process(i, v)
		list = append(list, f)

		for _, v := range f.Allergens {
			allergens[v] = Allergen{
				Allergen: v,
			}
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
					a := Allergen{
						Allergen:   commonAllergens[0],
						Ingredient: commonIngredients[0],
					}
					allergens[commonAllergens[0]] = a

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

	allergenList := []Allergen{}
	for _, v := range allergens {
		allergenList = append(allergenList, v)
	}

	// Arrange the ingredients alphabetically by their allergen
	sort.Sort(ByAllergen(allergenList))

	all := []string{}
	for _, v := range allergenList {
		all = append(all, v.Ingredient)
	}
	output := strings.Join(all, ",")
	fmt.Printf("%s\n", output)
}
