package main

import "fmt"

type ManualPlayer struct {
	State GameState
}

func (p *ManualPlayer) SetState(s GameState) {
	p.State = s
}

func (p *ManualPlayer) NextHand() Hand {
	p.State.showState()
	var i string
	_, err := fmt.Scanf("%s", &i)
	if err != nil {
		panic(err)
	}

	if i == "k" {
		return Up
	} else if i == "l" {
		return Right
	} else if i == "j" {
		return Down
	} else if i == "h" {
		return Left
	} else {
		return Quit
	}
}
