package main

import (
	"math/rand"
	//	"fmt"
)

type PlayoutResult struct {
	score float64
	hand  Hand
}

type MonteCarloPlayer struct {
	State    GameState
	tryCount int
}

func calcScore(sim Simulator) float64 {
	return float64(sim.Score() + sim.GetMaxTile()*200)
}

func (p *MonteCarloPlayer) SetState(s GameState) {
	p.State = s
	p.tryCount = 100
}

func (p *MonteCarloPlayer) Playout(firstHand Hand, res chan PlayoutResult) {
	avg := 0.0
	for cnt := 0; cnt < p.tryCount; cnt++ {
		var sim Simulator = &Kanna{p.State.Grid, 0, p.State.Over}
		if !sim.Move(firstHand) {
			break
		}

		for sim.Move(intToHand(rand.Intn(4))) {
		}
		avg += calcScore(sim)
	}
	res <- PlayoutResult{avg / float64(p.tryCount), firstHand}
}

func (p *MonteCarloPlayer) NextHand() Hand {
	res := make(chan PlayoutResult)
	go p.Playout(Up, res)
	go p.Playout(Right, res)
	go p.Playout(Down, res)
	go p.Playout(Left, res)

	bestAvg, ret := 0.0, Up
	for i := 0; i < 4; i++ {
		r := <-res
		if bestAvg < r.score {
			bestAvg = r.score
			ret = r.hand
		}
	}

	return ret
}
