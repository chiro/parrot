package main

import (
	"math"
	"time"
)

// Monte-Calro method with UCB1
type Shiro struct {
	State     GameState
	alpha     float64
	timeLimit int
}

func (s *Shiro) SetState(state GameState) {
	s.State = state
}

func (s *Shiro) Playout(first Hand, gen func() uint32) float64 {
	var sim Simulator = &Kanna{s.State.Grid, s.State.Score, s.State.Over, gen}
	if !sim.Move(first) {
		return -100
	}
	sim.AddRandomCell()

	for sim.Move(intToHand(int(gen() % 4))) {
		if !sim.AddRandomCell() {
			break
		}
	}
	return calcScore(sim)
}

func choiceMax(x *[4]float64, n *[4]int, total int, alpha float64) (ret int) {
	ret = 0
	m := x[0] + alpha*math.Sqrt(2*math.Log(float64(total))/float64(n[0]))

	for i := 1; i < 4; i++ {
		y := x[i] + alpha*math.Sqrt(2*math.Log(float64(total))/float64(n[i]))
		if m < y {
			ret = i
			m = y
		}
	}
	return
}

func (p *Shiro) NextHand(gen func() uint32) Hand {
	x := [4]float64{0, 0, 0, 0}
	n := [4]int{100, 100, 100, 100}

	for i := 0; i < 4; i++ {
		for j := 0; j < 100; j++ {
			x[i] += float64(p.Playout(intToHand(i), gen))
		}
		x[i] /= 100
	}

	playCnt := [4]int{100, 100, 100, 100}
	cnt := 0
	start := time.Now()
	for time.Since(start) < time.Duration(p.timeLimit)*time.Millisecond {
		cnt++
		i := choiceMax(&x, &n, cnt, p.alpha)
		playCnt[i]++
		y := float64(p.Playout(intToHand(i), gen))
		x[i] = (x[i]*float64(n[i]) + y) / float64(n[i]+1)
		n[i]++
	}

	// fmt.Printf("play %d, %d, %d, %d\n", playCnt[0], playCnt[1], playCnt[2], playCnt[3])
	return intToHand(choiceMax(&x, &n, 100, p.alpha))
}
