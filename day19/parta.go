package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jnewmano/advent2020/input"
	"github.com/jnewmano/advent2020/output"
)

func main() {

	// input.SetRaw(raw)
	var things = input.LoadSliceString("")

	for i, v := range things {
		if v == "" {
			things = things[i+1:]
			break
		}
		parseRule(v)
	}

	// find how many entries match rule 0
	count := 0
	r0 := rules[0]
	for _, v := range things {
		matched, matches := r0.Match(v)
		if matched {
			// check for left over unmatched input
			for _, v := range matches {
				if len(v) == 0 {
					count++
					break
				}
			}
		}
	}

	fmt.Println(count)

}

var rules = make(map[int]Rule)

type Rule interface {
	Match(string) (bool, []string)
}

type OR struct {
	Rules []Rule
}

func (o OR) Match(input string) (bool, []string) {

	matches := []string{}
	for _, v := range o.Rules {

		_, results := v.Match(input)

		matches = append(matches, results...)
	}

	return len(matches) > 0, matches
}

type AND struct {
	Rules []Rule
}

func (o AND) Match(input string) (bool, []string) {

	// get all the possible result strings from the first rule
	// in the set, and run it through the remaining ones
	matched, results := o.Rules[0].Match(input)

	if len(o.Rules) == 1 {
		if len(results) == 0 {
			return matched, nil
		} else {
			return true, results
		}
	}

	newAnd := AND{
		Rules: o.Rules[1:],
	}

	globalMatch := false
	agg := []string{}
	for _, v := range results {
		matched, result := newAnd.Match(v)
		if matched {
			globalMatch = true
		}
		agg = append(agg, result...)
	}

	return globalMatch, agg
}

type String struct {
	s string
}

func (s String) Match(input string) (bool, []string) {

	if len(input) == 0 {
		return false, nil
	}

	if s.s == string(input[0]) {
		return true, []string{string(input[1:])}
	}
	return false, nil
}

type Ref struct {
	Ref int
}

func (r Ref) Match(input string) (bool, []string) {
	n, ok := rules[r.Ref]
	if !ok {
		panic("rules not found: " + strconv.Itoa(r.Ref))
	}
	return n.Match(input)
}

func parseRule(s string) {
	parts := strings.Split(s, ": ")
	id, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	parts = strings.Split(strings.TrimSpace(parts[1]), " ")
	and := AND{}
	outer := OR{}
	for _, v := range parts {
		switch {
		case v == "|":
			// OR!!!
			outer.Rules = append(outer.Rules, and)
			and = AND{}
		case strings.HasPrefix(v, `"`):
			and.Rules = append(and.Rules, String{s: strings.Trim(v, `"`)})

		default:
			if num, err := strconv.Atoi(v); err == nil {
				and.Rules = append(and.Rules, Ref{Ref: num})
			} else {
				panic("unhandled rule: " + v)
			}
		}
	}

	if len(outer.Rules) > 0 {
		if len(and.Rules) > 0 {
			outer.Rules = append(outer.Rules, and)
		}

		rules[id] = outer

	} else {
		if len(and.Rules) == 1 {
			rules[id] = and.Rules[0]
		} else {
			rules[id] = and
		}
	}
}

var _ = output.High(nil)

var raw = `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`
