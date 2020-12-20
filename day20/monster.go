package main

import (
	"fmt"
	"strings"
)

func findMonsters(sea [][seaWidth]string) ([][seaWidth]string, int) {
	monster := loadMonster()
	monster.Display()

	marked := copySea(sea)
	found := 0
	for dy := range sea {
		for dx := 0; dx < seaWidth; dx++ {
			f := monster.markMonster(sea, marked, dx, dy, false)
			if f {
				found++
			}
		}
	}
	return marked, found
}

func (m *monster) markMonster(sea [][seaWidth]string, marked [][seaWidth]string, dx int, dy int, clear bool) bool {
	hits := 0
	for y, v := range m.data {
		for x, bodyPart := range v {
			if bodyPart != "#" {
				continue
			}

			// count the monster hits
			imgX := x + dx
			imgY := y + dy
			if imgX >= seaWidth || imgY >= len(sea) {
				continue
			}

			// see how many monster hits we get at this position
			if sea[imgY][imgX] == "#" {
				if clear {
					marked[imgY][imgX] = "O"
				}
				hits++
			}

		}
	}

	if hits == m.points && clear == false {
		m.markMonster(sea, marked, dx, dy, true)
		return true
	}
	return false
}

type monster struct {
	data   [][]string
	points int
}

func loadMonster() monster {
	rows := strings.Split(monsterInput, "\n")
	var data [][]string = make([][]string, 3)
	for y, row := range rows {
		data[y] = make([]string, 20)
		for x, xy := range row {
			data[y][x] = string(xy)
		}
	}

	m := monster{
		data:   data,
		points: 15,
	}
	return m
}

func (m *monster) Display() {

	for _, v := range m.data {
		for _, x := range v {
			fmt.Print(string(x))
		}
		fmt.Print("\n")
	}
	fmt.Print("~~~~~~~~~~~~~~~~~~~~\n")
	fmt.Print("~~~~~~~~~~~~~~~~~~~~\n")
}

var monsterInput = `                  # 
#    ##    ##    ###
 #  #  #  #  #  #   `
