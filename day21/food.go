package main

import (
	"strings"

	"github.com/jnewmano/advent2020/output"
)

type ByAllergen []Allergen

func (a ByAllergen) Len() int           { return len(a) }
func (a ByAllergen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAllergen) Less(i, j int) bool { return a[i].Allergen < a[j].Allergen }

type Allergen struct {
	Allergen   string
	Ingredient string
}

type Ingredient struct {
	PossibleAllergens map[string]int
	Foods             []int
	Allergen          string
}

type Food struct {
	ID          int
	Ingredients []string
	Allergens   []string
}

func process(id int, s string) Food {
	parts := strings.Split(s, " (contains ")
	ingredients := strings.Split(parts[0], " ")
	ags := strings.Trim(parts[1], ")")
	allergens := strings.Split(ags, ", ")

	f := Food{
		ID:          id,
		Ingredients: ingredients,
		Allergens:   allergens,
	}

	return f
}

var _ = output.High(nil)

var raw = `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`
