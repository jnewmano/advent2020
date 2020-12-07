package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jnewmano/advent2020/input"
	"github.com/jnewmano/advent2020/output"
)

func main() {

	sum := parta()
	fmt.Println(sum)
}

func parta() interface{} {
	// input.SetRaw(raw2)
	// var things = input.Load()
	// var things = input.LoadSliceSliceString("")
	var things = input.LoadSliceString("")

	rules := parseBags(things)
	for _, v := range rules {
		fmt.Printf("%s - %#v\n", v.bag, v.contains)
	}

	// How many individual bags are required inside your single shiny gold bag?
	sum := howManyBagsInBag(rules, "shiny gold")

	return sum
}

func howManyBagsInBag(rules []Rule, color string) int {

	fmt.Println("finding rule for", color)
	r := findRuleForColor(rules, color)

	sum := 0
	for _, v := range r.contains {
		n := howManyBagsInBag(rules, v.color)
		sum += v.count + v.count*n
	}

	return sum
}

func findRuleForColor(rules []Rule, color string) Rule {
	for _, v := range rules {
		if v.bag == color {
			return v
		}
	}
	panic("couldn't find rule for: " + color)
}

func canContain(rule Rule, color string) bool {
	for _, v := range rule.contains {
		if v.color == color {
			return true
		}
	}
	return false
}

func parseBags(things []string) []Rule {
	all := []Rule{}
	for _, v := range things {
		r := parseRule(v)
		all = append(all, r)
	}
	return all
}

type Rule struct {
	bag      string
	contains []Contains
}

type Contains struct {
	count int
	color string
}

func parseRule(r string) Rule {
	fmt.Println(r)
	parts := strings.Split(r, " bags contain ")
	bag := strings.TrimSpace(parts[0])

	items := strings.TrimSpace(parts[1])
	bags := strings.Split(items, ", ")

	all := make([]Contains, 0)
	for _, v := range bags {
		v = strings.Trim(v, ". s")

		if v == "no other bag" {
			continue
		}

		var count int

		parts := strings.SplitN(v, " ", 2)
		count, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		bag := strings.TrimSuffix(parts[1], "bag")

		all = append(all, Contains{
			count: count,
			color: strings.TrimSpace(bag),
		})

	}

	rule := Rule{
		bag:      bag,
		contains: all,
	}

	return rule
}

var _ = output.High(nil)

var raw = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

var raw2 = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`
