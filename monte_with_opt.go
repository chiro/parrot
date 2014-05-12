package main

import (
	"time"
)

type OptMonte struct {
	State     GameState
	timeLimit int
}

func (p *OptMonte) SetState(s GameState) {
	p.State = s
}

func (p *OptMonte) Playout(first Hand, res chan PlayoutResult, gen func() uint32, now time.Time) {
	avg, cnt := 0.0, 0
	for time.Since(now) < time.Duration(p.timeLimit)*time.Millisecond {
		cnt++
		var sim Simulator = &Kanna{p.State.Grid, p.State.Score, p.State.Over, gen}
		if !sim.Move(first) {
			res <- PlayoutResult{-100, first}
			return
		}
		sim.AddRandomCell()

		for sim.Move(intToHand(int(gen() % 4))) {
			if !sim.AddRandomCell() {
				break
			}
		}
		avg += calcScore(sim)
	}
	res <- PlayoutResult{avg / float64(cnt), first}
}

func (p *OptMonte) NextHand(gen func() uint32) Hand {
	res := make(chan PlayoutResult)
	now := time.Now()
	go p.Playout(Up, res, gen, now)
	go p.Playout(Right, res, gen, now)
	go p.Playout(Down, res, gen, now)
	go p.Playout(Left, res, gen, now)

	bestAvg, ret := 0.0, Up
	for i := 0; i < 4; i++ {
		r := <-res
		if bestAvg < r.score {
			bestAvg = r.score
			ret = r.hand
		}
	}

	// fmt.Printf("after best move %v:\n", ret)
	// var sim Midori = Midori{encode2(&p.State.Grid), 0, p.State.Over, r}
	// sim.Move(ret)
	// fmt.Printf("%v\n", sim.GetState())

	return ret
}
