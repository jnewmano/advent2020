package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jnewmano/advent2020/input"
)

var games = 1

func main() {

	//	input.SetRaw(raw)
	var things = input.LoadSliceString("")

	game := process(things)

	winner := game.Play()
	fmt.Printf("Winner is Player %d - Score %d\n", winner, game.Score())
}

func (g *Game) Play() int {

	for {
		if g.GameOver() {
			break
		}

		instantEnd := g.PlayHand()
		if instantEnd == 0 {
			return 1
		}
	}
	winner := 0
	if len(g.PlayerA) == 0 {
		winner = 2
	} else if len(g.PlayerB) == 0 {
		winner = 1
	} else {
		panic("don't know who won")
	}

	return winner
}

func (g *Game) PlayHand() int {
	// check that we haven't played this hand before
	s := fmt.Sprintf("%v-%v", g.PlayerA, g.PlayerB)
	_, ok := g.States[s]
	if ok {
		// we're stuck, return 0
		return 0
	}
	g.States[s] = struct{}{}

	cardA := g.PlayerA[0]
	g.PlayerA = g.PlayerA[1:]

	cardB := g.PlayerB[0]
	g.PlayerB = g.PlayerB[1:]

	// recursive rules!!!!!
	// if each player has at least the number of cards in hand of the value they just
	// played, then the winner is the winner of the recursive game!
	winner := 0
	if len(g.PlayerA) >= cardA && len(g.PlayerB) >= cardB {
		winner = g.SpawnGame(cardA, cardB)
	} else if cardA > cardB {
		winner = 1
	} else if cardA < cardB {
		winner = 2
	} else {
		panic("what?")
	}

	switch winner {
	case 1:
		g.PlayerA = append(g.PlayerA, cardA, cardB)
	case 2:
		g.PlayerB = append(g.PlayerB, cardB, cardA)
	default:
		panic("unknown winner")
	}

	return -1

}

func (g *Game) SpawnGame(a int, b int) int {
	games++
	newGame := Game{
		ID:     games,
		States: make(map[string]struct{}),
	}

	for i := 0; i < a; i++ {
		newGame.PlayerA = append(newGame.PlayerA, g.PlayerA[i])
	}
	for i := 0; i < b; i++ {
		newGame.PlayerB = append(newGame.PlayerB, g.PlayerB[i])
	}

	return newGame.Play()
}

func (g *Game) GameOver() bool {
	return len(g.PlayerA) == 0 || len(g.PlayerB) == 0
}

func process(input []string) *Game {
	g := Game{
		ID:     1,
		States: make(map[string]struct{}),
	}
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
	ID      int
	PlayerA []int
	PlayerB []int

	States map[string]struct{}
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

var raw2 = `Player 1:
43
19

Player 2:
2
29
14`
