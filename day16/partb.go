package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jnewmano/advent2020/input"
)

type Rule struct {
	start int
	stop  int
}

var ticketRules = make(map[string][]Rule)
var ticketMap = make(map[int][]string)

func main() {

	//input.SetRaw(raw)

	var things = input.LoadSliceString("")

	// var list = make([]int)
	for i, v := range things {
		if v == "" {
			things = things[i:]
			break
		}
		processRule(v)
	}

	validTickets := []string{things[2]}
	// sort out valid tickets
	for _, v := range things[5:] {
		invalid := getInvalidNumbers(v)
		if len(invalid) > 0 {
			continue
		}
		validTickets = append(validTickets, v)
	}

	fmt.Println("valid tickets:", len(validTickets))
	fmt.Println(len(ticketRules), "ticket rules")
	for k, rule := range ticketRules {

		// find which ticket column satisfies the rule
		for i := 0; i < len(ticketRules); i++ {

			// assume all tickets satisfy rule `rule` for ticket index i
			valid := true
			for _, t := range validTickets {
				if checkTicketField(t, i, rule) == false {
					valid = false
					break
				}
			}
			if valid == true {
				ticketMap[i] = append(ticketMap[i], k)
			}
		}
	}

	fields := []int{}
	// reduce the ticket rules
	for i := 0; i < len(ticketRules); i++ {
		for idx, v := range ticketMap {
			if len(v) == 1 {
				removeRuleFromMap(ticketMap, v[0])
				if strings.HasPrefix(v[0], "departure") {
					fields = append(fields, idx)
				}
			}
		}
	}

	parts := strings.Split(things[2], ",")

	product := 1
	for _, v := range fields {
		n, _ := strconv.Atoi(parts[v])
		product *= n
	}
	fmt.Println(product)

}

func removeRuleFromMap(maps map[int][]string, s string) {
	for k, v := range maps {
		left := []string{}
		for _, v := range v {
			if v != s {
				left = append(left, v)
			}
		}
		maps[k] = left
	}
}

func getInvalidNumbers(t string) []int {
	parts := strings.Split(t, ",")

	invalid := []int{}
	for idx, n := range parts {
		valid := false
		n, _ := strconv.Atoi(n)
		for _, rule := range ticketRules {
			if checkTicketField(t, idx, rule) {
				valid = true
			}

		}
		if valid == false {
			invalid = append(invalid, n)
		}

	}
	return invalid
}

func checkTicketField(t string, idx int, rule []Rule) bool {
	parts := strings.Split(t, ",")
	n, err := strconv.Atoi(parts[idx])
	if err != nil {
		panic(err)
	}

	valid := false
	for _, v := range rule {
		if n >= v.start && n <= v.stop {
			valid = true
			break
		}
	}

	return valid
}

func processRule(s string) {
	parts := strings.Split(s, ": ")
	name := strings.TrimSpace(parts[0])

	ranges := strings.Split(parts[1], " or ")

	rules := []Rule{}
	for _, v := range ranges {
		p := strings.Split(v, "-")
		start, err := strconv.Atoi(p[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(p[1])
		if err != nil {
			panic(err)
		}

		rules = append(rules, Rule{start: start, stop: end})
	}
	ticketRules[name] = rules
}

var raw = `class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`
