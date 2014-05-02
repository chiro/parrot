package main

import "math/rand"

type RandomPlayer struct {
	State GameState
}

func (p *RandomPlayer) SetState(s GameState) {
	p.State = s
}

func (p * RandomPlayer) NextHand() Hand {
	p.State.showState()
	var h int = rand.Int() % 4;
	return intToHand(h)
}
