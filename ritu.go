package main

import (
	"fmt"
	"math"
	"time"
)

// Monte-Calro method with UCB1-Tuned
type Ritu struct {
	State     GameState
	alpha     float64
	timeLimit int
}

func (r *Ritu) SetState(state GameState) {
	r.State = state
}

func (r *Ritu) Playout(first Hand, gen func() uint32) float64 {
	var sim Simulator = &Kanna{r.State.Grid, r.State.Score, r.State.Over, gen}
	if !sim.Move(first) {
		return 0.0
	}
	sim.AddRandomCell()

	for sim.Move(intToHand(int(gen() % 4))) {
		if !sim.AddRandomCell() {
			break
		}
	}
	return calcScore(sim)
}

func (r *Ritu) ucb(x float64, n int, total int, alpha float64) float64 {
	beta := math.Min(0.25, x-x*x+math.Sqrt(alpha*math.Log(float64(total))/float64(n)))
	return x + math.Sqrt(math.Log(float64(total))/float64(n)*beta)
}

func (r *Ritu) plainUcb(x float64, n int, total int, alpha float64) float64 {
	return x + math.Sqrt(alpha*math.Log(float64(total))/float64(n))
}

func (r *Ritu) choiceMax(x *[4]float64, n *[4]int, total int, alpha float64) (ret int) {
	maxScore := 0.0
	for i := 0; i < 4; i++ {
		maxScore = math.Max(maxScore, x[i])
	}
	maxScore *= 1.25
	ret = 0
	m := r.ucb(x[0]/maxScore, n[0], total, alpha)
	//m := r.plainUcb(x[0]/maxScore, n[0], total, alpha)

	for i := 1; i < 4; i++ {
		y := r.plainUcb(x[i]/maxScore, n[i], total, alpha)
		//y := r.ucb(x[i]/maxScore, n[i], total, alpha)
		if m < y {
			ret = i
			m = y
		}
	}
	return
}

func (r *Ritu) NextHand(gen func() uint32) Hand {
	x := [4]float64{0, 0, 0, 0}
	n := [4]int{100, 100, 100, 100}

	for i := 0; i < 4; i++ {
		for j := 0; j < 100; j++ {
			x[i] += float64(r.Playout(intToHand(i), gen))
		}
		x[i] /= 100
	}

	cnt := 0
	start := time.Now()
	for time.Since(start) < time.Duration(r.timeLimit)*time.Millisecond {
		cnt++
		i := r.choiceMax(&x, &n, cnt, r.alpha)
		y := float64(r.Playout(intToHand(i), gen))
		x[i] = (x[i]*float64(n[i]) + y) / float64(n[i]+1)
		n[i]++
	}

	Debug(fmt.Sprintf("score %f %f %f %f", x[0], x[1], x[2], x[3]))
	Debug(fmt.Sprintf("play %5d, %5d, %5d, %5d", n[0], n[1], n[2], n[3]))
	return intToHand(r.choiceMax(&x, &n, 100, r.alpha))
}
