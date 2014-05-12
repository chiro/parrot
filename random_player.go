package main

type RandomPlayer struct {
	State GameState
}

func (p *RandomPlayer) SetState(s GameState) {
	p.State = s
}

func (p *RandomPlayer) NextHand(gen func() uint32) Hand {
	return intToHand(int(gen() % 4))
}
