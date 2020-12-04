package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jnewmano/advent2020/input"
)

func main() {

	sum := parta()
	fmt.Println(sum)
}

/*
detecting which passports have all required fields. The expected fields are as follows:

byr (Birth Year)
iyr (Issue Year)
eyr (Expiration Year)
hgt (Height)
hcl (Hair Color)
ecl (Eye Color)
pid (Passport ID)
cid (Country ID)
*/
func parta() interface{} {
	// input.SetRaw(raw)
	things := input.LoadSliceString("\n\n")

	expected := map[string]func(s string) bool{
		"byr": byr,
		"iyr": iyr,
		"eyr": eyr,
		"hgt": hgt,
		"hcl": hcl,
		"ecl": ecl,
		"pid": pid,
		//	"cid": nil,
	}

	countValid := 0
	for _, v := range things {
		parts := parse(v)
		valid := true

		for key, f := range expected {
			data, ok := parts[key]
			if !ok && key != "cid" {
				valid = false
				break
			}
			if f != nil && f(data) == false {
				valid = false
				break
			}
		}
		if valid {
			countValid++
		}
	}
	return countValid
}

// four digits; at least 1920 and at most 2002
func byr(s string) bool {
	if len(s) != 4 {
		return false
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return v >= 1920 && v <= 2002
}

// four digits; at least 1920 and at most 2002
func iyr(s string) bool {
	if len(s) != 4 {
		return false
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return v >= 2010 && v <= 2020
}

// four digits; at least 2020 and at most 2030
func eyr(s string) bool {
	if len(s) != 4 {
		return false
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return v >= 2020 && v <= 2030
}

/*
(Hair Color) - a # followed by exactly six characters 0-9 or a-f
*/
func hcl(s string) bool {
	if strings.HasPrefix(s, "#") == false {
		return false
	}
	s = s[1:]
	for _, v := range s {
		if v < '0' || v > 'f' {
			return false
		}
	}
	return true
}

/*
	a number followed by either cm or in:
	If cm, the number must be at least 150 and at most 193.
	If in, the number must be at least 59 and at most 76.
*/
func hgt(s string) bool {
	max := 76
	min := 59
	switch {
	case strings.HasSuffix(s, "in"):
	case strings.HasSuffix(s, "cm"):
		max = 193
		min = 150
	}

	s = strings.TrimSuffix(s, "cm")
	s = strings.TrimSuffix(s, "in")

	v, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	return v >= min && v <= max
}

// exactly one of: amb blu brn gry grn hzl oth
func ecl(s string) bool {
	return s == "amb" ||
		s == "blu" ||
		s == "brn" ||
		s == "gry" ||
		s == "grn" ||
		s == "hzl" ||
		s == "oth"
}
func pid(s string) bool {
	if len(s) != 9 {
		return false
	}
	for _, v := range s {
		_, err := strconv.Atoi(string(v))
		if err != nil {
			return false
		}
	}
	return true
}

func parse(s string) map[string]string {
	parts := strings.Split(s, "\n")
	all := strings.Join(parts, " ")

	data := make(map[string]string)
	kv := strings.Split(all, " ")
	for _, v := range kv {
		p := strings.Split(v, ":")
		data[p[0]] = p[1]
	}

	return data
}

var raw = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`
