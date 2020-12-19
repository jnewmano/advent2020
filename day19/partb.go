package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jnewmano/advent2020/input"
)

func main() {

	// input.SetRaw(raw2)
	var things = input.LoadSliceString("")

	for i, v := range things {
		if v == "" {
			things = things[i+1:]
			break
		}
		parseRule(v)
	}
	// and new rules
	parseRule(`8: 42 | 42 8`)
	parseRule(`11: 42 31 | 42 11 31`)

	// find how many entries match rule 0
	count := 0
	r0 := rules[0]
	for _, v := range things {
		matched, matches := r0.Match(v)
		if matched {
			for _, v := range matches {
				// check for left over unmatched input
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

var raw1 = `0: 4 1 5
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

var raw2 = `42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`
