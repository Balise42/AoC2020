package main

import (
	"fmt"
	"strings"
)

func (day *Day) Day22a() {
	fmt.Printf("Part 1: %d\n", ComputeDay22a(day.input))
}

func (day *Day) Day22b() {
	fmt.Printf("Part 2: %d\n", ComputeDay22b(day.input))
}


func ComputeDay22a(input string) int {
	deckstr := strings.Split(input, "\n\n")
	deck1 := convertStringListToInts(strings.Split(deckstr[0], "\n")[1:])
	deck2 := convertStringListToInts(strings.Split(deckstr[1], "\n")[1:])
	for len(deck1) != 0 && len(deck2) != 0 {
		deck1, deck2 = playDay22aRound(deck1, deck2)
	}
	return computeDay22DeckValue(deck1, deck2)
}

func computeDay22DeckValue(deck1 []int, deck2 []int) int {
	var deck []int
	if len(deck1) == 0 {
		deck = deck2
	} else {
		deck = deck1
	}
	res := 0
	for i, card := range deck {
		res += (len(deck) - i) * card
	}
	return res
}

func playDay22aRound(deck1 []int, deck2 []int) ([]int, []int) {
	c1, d1 := deck1[0], deck1[1:]
	c2, d2 := deck2[0], deck2[1:]
	if c1 < c2 {
		d2 = append(d2, c2, c1)
	} else {
		d1 = append(d1, c1, c2)
	}
	return d1, d2
}

func ComputeDay22b(input string) int {
	deckstr := strings.Split(input, "\n\n")
	deck1 := convertStringListToInts(strings.Split(deckstr[0], "\n")[1:])
	deck2 := convertStringListToInts(strings.Split(deckstr[1], "\n")[1:])

	deck1, deck2, _ = playDay22bGame(deck1, deck2)
	return computeDay22DeckValue(deck1, deck2)
}

type Day22State struct {
	deck1 []int
	deck2 []int
}

func (d Day22State) equals(other Day22State ) bool {
	if len(d.deck1) != len(other.deck1) {
		return false
	}
	if len(d.deck2) != len(other.deck2) {
		return false
	}
	for i := range d.deck1 {
		if d.deck1[i] != other.deck1[i] {
			return false
		}
	}
	for i := range d.deck2 {
		if d.deck2[i] != other.deck2[i] {
			return false
		}
	}
	return true
}


func playDay22bGame(deck1 []int, deck2 []int) ([]int, []int, bool) {
	states := make([]Day22State, 0)
	states = append(states, Day22State{deck1, deck2})

	d1 := make([]int, len(deck1))
	d2 := make([]int, len(deck2))
	copy(d1, deck1)
	copy(d2, deck2)
	var c1, c2 int

	for len(d1) != 0 && len(d2) != 0 {
		c1, d1 = d1[0], d1[1:]
		c2, d2 = d2[0], d2[1:]

		if len(d1) >= c1 && len(d2) >= c2 {
			_, _, winner := playDay22bGame(d1[:c1], d2[:c2])
			if !winner {
				d1 = append(d1, c1, c2)
			} else {
				d2 = append(d2, c2, c1)
			}
		} else {
			if c1 < c2 {
				d2 = append(d2, c2, c1)
			} else {
				d1 = append(d1, c1, c2)
			}
		}

		currState := Day22State{d1, d2}
		for _, state := range states {
			if currState.equals(state) {
				return d1, d2, false
			}
		}
		states = append(states, currState)
	}
	if len(d1) == 0 {
		return d1, d2, true
	}
	return d1, d2, false
}