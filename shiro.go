package main

import (
	"math"
	"math/rand"
)

// Monte-Calro method with UCB1
type Shiro struct {
	State GameState
	alpha float64
}

func (s *Shiro) Playout(first Hand, gen func() uint32) int {
	var sim Simulator = &Kanna{s.State.Grid, 0, s.State.Over, gen}
	if !sim.Move(first) {
		return 0
	}
	for sim.GetAvailableCells() > 0 {
		var moved bool = sim.Move(intToHand(rand.Intn(4)))
		if moved {
			continue
		}

		var goNext bool = false
		for hand := 0; hand < 4; hand++ {
			moved = sim.Move(intToHand(int(gen() % 4)))
			if moved {
				goNext = true
				break
			}
		}
		if !goNext {
			break
		}
	}
	return sim.Score()
}

func (s *Shiro) SetState(state GameState) {
	s.State = state
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
	INF := 100000000.0
	x := [4]float64{INF, INF, INF, INF}
	n := [4]int{1, 1, 1, 1}

	for i := 0; i < 4; i++ {
		x[i] = float64(p.Playout(intToHand(i), gen))
	}

	for cnt := 0; cnt < 500; cnt++ {
		i := choiceMax(&x, &n, cnt, p.alpha)
		y := float64(p.Playout(intToHand(i), gen))
		x[i] = (x[i]*float64(n[i]) + y) / float64(n[i]+1)
		n[i]++
	}

	return intToHand(choiceMax(&x, &n, 100, p.alpha))
}
