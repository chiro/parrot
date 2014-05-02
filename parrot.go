package main

import "flag"

// This variable means where the service is.
// const BASE_URL = "http://2048.semantics3.com/hi/"
// const BASE_URL = "http://ring:2048/hi/"
const BASE_URL = "http://localhost:8080/hi/"

func main() {
	// Command-line options
	var q = flag.Bool("q", false, "Suppress outputs. Show only final state.")
	flag.Parse()

	var m *Manager = new(Manager)
	// Please change the next line to change AI.
	var p Player = new(RandomPlayer)
	m.Initialize(p, *q)
	m.StartGame()
}
