package main

type Hand int
const (
	Up Hand = iota
	Right
	Down
	Left
    Quit
)

type Player interface {
	NextHand() Hand
	SetState(GameState)
}

type GameState struct {
	Grid [][]int
	Won bool
	Moved bool
	Over bool
	Score int
	Points int
	Zen string
	Session_id string
}
