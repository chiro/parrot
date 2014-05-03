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
	p.tryCount = 30
}

func (p *MonteCarloPlayer) Playout(firstHand Hand) float64 {
	avg := 0.0
	for cnt := 0; cnt < p.tryCount; cnt++ {
		var sim Simulator = Simulator{p.State.Grid, 0, p.State.Over}
		sim.move(firstHand)

		pcnt := 0
		for !sim.Finish && pcnt < 1000 {
			sim.move(intToHand(rand.Int() % 4))
			pcnt++
		}
		avg += float64(sim.Score + sim.getAvailableCells()*100)
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
