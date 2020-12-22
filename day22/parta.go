package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jnewmano/advent2020/input"
)

func main() {

	//input.SetRaw(raw)
	var things = input.LoadSliceString("")

	game := process(things)

	i := 0
	for {
		i++
		fmt.Printf("-- Round %d --\n", i)
		fmt.Println(game.PlayerA)
		fmt.Println(game.PlayerB)

		if game.GameOver() {
			break
		}

		cardA := game.PlayerA[0]
		game.PlayerA = game.PlayerA[1:]

		cardB := game.PlayerB[0]
		game.PlayerB = game.PlayerB[1:]

		if cardA > cardB {
			game.PlayerA = append(game.PlayerA, cardA, cardB)
		} else if cardA < cardB {
			game.PlayerB = append(game.PlayerB, cardB, cardA)
		} else {
			panic("got a tie")
		}

	}

	fmt.Println(game.Score())
}

func (g *Game) GameOver() bool {
	return len(g.PlayerA) == 0 || len(g.PlayerB) == 0
}

func process(input []string) *Game {
	g := Game{}
	playerCount := 0
	for _, v := range input {
		if strings.HasPrefix(v, "Player") {
			playerCount++
			continue
		}
		if v == "" {
			continue
		}

		p, err := strconv.Atoi(v)
		if err != nil {
			panic(v)
		}

		switch playerCount {
		case 1:
			g.PlayerA = append(g.PlayerA, p)
		case 2:
			g.PlayerB = append(g.PlayerB, p)
		default:
			panic("unknown player")
		}
	}

	return &g

}

type Game struct {
	PlayerA []int
	PlayerB []int
}

func (g *Game) Score() int {
	cards := g.PlayerA
	if len(cards) == 0 {
		cards = g.PlayerB
	}
	score := 0
	l := len(cards)
	for i, v := range cards {

		score += v * (l - i)
	}
	return score
}

var raw = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`
