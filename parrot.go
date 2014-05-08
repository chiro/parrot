package main

import (
	"flag"
	"fmt"
	"github.com/chiro/parrot/random"
	"sort"
)

// This variable means where the service is.
// const BASE_URL = "http://2048.semantics3.com/hi/"
// const BASE_URL = "http://ring:2048/hi/"
const BASE_URL = "http://localhost:8080/hi/"

func playOnce(q bool, done chan GameState) {
	var m *Manager = new(Manager)
	// Please change the next line to change AI.
	// var p Player = new(RandomPlayer)
	var p Player = new(MonteCarloPlayer)
	var r random.Gen = new(random.Std)
	m.Initialize(p, q)
	m.StartGame(r)
	done <- m.state
}

func play(q bool, t int) {
	done := make(chan GameState)
	for i := 0; i < t; i++ {
		go playOnce(q, done)
	}

	var maxTile map[int]int = map[int]int{}
	var scores []int = make([]int, t)
	var avg float64 = 0.0
	for i := 0; i < t; i++ {
		var s GameState = <-done
		fmt.Print(".")
		maxTile[s.MaxTile()]++
		scores[i] = s.Score
		avg += float64(scores[i])
	}
	fmt.Println("")
	avg /= float64(t)
	sort.Sort(sort.IntSlice(scores))
	fmt.Printf("min = %d, max = %d, avg = %f\n", scores[0], scores[len(scores)-1], avg)
	fmt.Println("----------  Points ----------")

	var keys []int = make([]int, len(maxTile))
	i := 0
	for k, _ := range maxTile {
		keys[i] = k
		i++
	}
	sort.Sort(sort.IntSlice(keys))
	for i := 0; i < len(keys); i++ {
		fmt.Printf("%d : %d\n", keys[i], maxTile[keys[i]])
	}
}

func main() {
	// Command-line options
	var q = flag.Bool("q", false, "Suppress outputs. Show only final state.")
	var t = flag.Int("t", 1, "How many times we play the game.")
	flag.Parse()

	play(*q, *t)
}
