package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jnewmano/advent2020/input"
	"github.com/jnewmano/advent2020/output"
)

func main() {

	// input.SetRaw(raw)
	var things = input.LoadSliceString("")

	var result = make([]int, len(things))
	for i, v := range things {
		tokens := tokenize(v)
		result[i] = math(tokens)
	}

	fmt.Println(output.Sum(result))
}

func math(s []string) int {
	// find expressions that we can evaluate ()
	if len(s) == 0 {
		fmt.Println("no math to do?")
		os.Exit(1)
	}

	if len(s) == 1 {
		result, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}
		return result
	}

	hasParans := false
	idxStart := 0
	idxEnd := 0
	for i, v := range s {
		if v == "(" {
			hasParans = true
			idxStart = i
		} else if v == ")" {
			idxEnd = i
			break
		}
	}

	if hasParans {
		result := math(s[idxStart+1 : idxEnd])
		// remove from list

		end := s[idxEnd+1:]
		s = append(s[0:idxStart], strconv.Itoa(result))
		s = append(s, end...)

	} else {
		// no parans, now see if we have additions to do, if so, do those all first
		hasAdditions := false
		idx := 0
		for i, v := range s {
			if v == "+" {
				hasAdditions = true
				idx = i
				break
			}
		}

		// do additions first
		if hasAdditions {
			result := perform(s[idx-1], s[idx+1], s[idx])

			// remove the operation from the list
			end := s[idx+2:]
			s = append(s[0:idx-1], strconv.Itoa(result))
			s = append(s, end...)

		} else {
			// lastly do all of our multiplications
			result := perform(s[0], s[2], s[1])
			s = s[2+1:]
			// push the result back onto list at the beginning
			s = append([]string{strconv.Itoa(result)}, s...)

		}
	}

	return math(s)
}

func perform(x, y string, op string) int {
	a, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(y)
	if err != nil {
		panic(err)
	}

	switch op {
	case "*":
		return a * b
	case "+":
		return a + b
	default:
		panic("unknown operator")
	}
}

func tokenize(s string) []string {

	var num string
	var tokens []string
	for _, v := range s {
		if v >= '0' && v <= '9' {
			num += string(v)
		} else if v == ')' {
			if num != "" {
				tokens = append(tokens, num)
				num = ""
			}
			tokens = append(tokens, string(v))
		} else if v == '(' {
			tokens = append(tokens, string(v))
		} else if v == '+' || v == '*' {
			tokens = append(tokens, string(v))
		} else if v == ' ' {
			if num != "" {
				tokens = append(tokens, num)
				num = ""
			}
		}
	}
	if num != "" {
		tokens = append(tokens, num)
	}

	return tokens
}

var raw = `3 * 1 + 2`

// var raw = `2 * 3 + (4 * 5)`
