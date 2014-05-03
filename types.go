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
	Grid       [4][4]int
	Won        bool
	Moved      bool
	Over       bool
	Score      int
	Points     int
	Zen        string
	Session_id string
}
