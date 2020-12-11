package input

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var raw string

func SetRaw(a string) error {
	raw = strings.TrimSpace(a)
	return nil
}

func Load() string {
	if raw != "" {
		return raw
	}

	_, err := os.Stat("input.txt")
	if os.IsNotExist(err) {
		loadInputFromAPI()
	}

	all, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	if checkSHA(all) != true {
		panic("sha didn't validate, input.txt file has been modified")
	}

	return strings.TrimSpace(string(all))
}

func LoadSliceString(sep string) []string {
	if sep == "" {
		sep = "\n"
	}
	all := Load()
	things := strings.Split(all, sep)
	return things
}

func LoadSliceInt(sep string) []int {

	things := LoadSliceString(sep)

	var ints []int

	for _, v := range things {
		if v == "" {
			continue
		}

		a, err := strconv.Atoi(v)
		if err != nil {
			panic(fmt.Sprintf("invalid input: [%s] (%s)\n", v, err))
		}
		ints = append(ints, a)
	}
	return ints

}

func LoadSliceSliceString(sepLine string, sepRow string) [][]string {
	if sepRow == "" {
		sepRow = ","
	}

	things := LoadSliceString(sepLine)
	var resp [][]string

	for _, v := range things {
		v = strings.TrimSpace(v)
		parts := strings.Split(v, sepRow)
		resp = append(resp, parts)
	}
	return resp
}

func loadInputFromAPI() {
	path, _ := os.Getwd()
	inputFilePath := filepath.Join(path, "input.txt")

	day := filepath.Base(path)
	if strings.HasPrefix(day, "day") == false {
		panic("invalid path, expected day in direction name")
	}

	day = strings.TrimPrefix(day, "day")

	fmt.Println("loading remote input file")

	cmd := exec.Command("aocdl", "-year", "2020", "-day", day, "-output", inputFilePath)
	fmt.Println("%", cmd.Args)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		panic(err)
	}

	fmt.Println(string(output))
}

func checkSHA(output []byte) bool {

	expectedSHA := sha256.Sum256(output)

	_, err := os.Stat("input.txt.sha")
	if os.IsNotExist(err) {
		writeSHA(expectedSHA[:])
	}

	sha, err := ioutil.ReadFile("input.txt.sha")
	if err != nil {
		panic(err)
	}

	if bytes.Compare(sha, expectedSHA[:]) != 0 {
		fmt.Printf("got:      %x\n", sha)
		fmt.Printf("expected: %x\n", expectedSHA)
		panic("input.txt sha is different than expected")
	}

	return true
}

func writeSHA(sha []byte) {
	err := ioutil.WriteFile("input.txt.sha", sha, 0444)
	if err != nil {
		panic(err)
	}
}
