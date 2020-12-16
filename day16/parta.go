package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jnewmano/advent2020/input"
	"github.com/jnewmano/advent2020/output"
)

type Rule struct {
	start int
	stop  int
}

var ticketRules = make(map[string][]Rule)

func main() {

	//input.SetRaw(raw)
	// var things = input.Load()
	// var things = input.LoadSliceSliceString("")
	var things = input.LoadSliceString("")

	// var list = make([]int)
	for i, v := range things {
		if v == "" {
			things = things[i:]
			break
		}
		processRule(v)
	}

	invalid := []int{}
	for _, v := range things[5:] {
		invalid = append(invalid, getInvalidNumbers(v)...)

	}
	fmt.Println("invalids", invalid)
	fmt.Println(output.Sum(invalid))

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
		}
	}
	return valid
}

func processRule(s string) {
	parts := strings.Split(s, ": ")
	name := parts[0]

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

var raw = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`
