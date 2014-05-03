package main

import (
	"math/rand"
	//	"fmt"
)

type MonteCarloPlayer struct {
	State    GameState
	tryCount int
}

func (p *MonteCarloPlayer) SetState(s GameState) {
	p.State = s
	p.tryCount = 100
}

func (p *MonteCarloPlayer) Playout(firstHand Hand) float64 {
	avg := 0.0
	for cnt := 0; cnt < p.tryCount; cnt++ {
		var sim Simulator = &Kanna{p.State.Grid, 0, p.State.Over}
		if !sim.Move(firstHand) {
			break
		}
		avg += 1.0

		for sim.Move(intToHand(rand.Int() % 4)) {
		}
		avg += float64(sim.Score()*10 + sim.GetAvailableCells()*100)
	}
	return avg / float64(p.tryCount)
}

func (p *MonteCarloPlayer) NextHand() Hand {
	us := p.Playout(Up)
	rs := p.Playout(Right)
	ds := p.Playout(Down)
	ls := p.Playout(Left)

	ret := 0
	if rs > us {
		ret = 1
		us = rs
	}
	if ds > us {
		ret = 2
		us = ds
	}
	if ls > us {
		ret = 3
		us = ls
	}

	return intToHand(ret)
}
