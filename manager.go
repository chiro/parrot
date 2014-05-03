package main

import "fmt"

type Manager struct {
	player    Player
	sessionId string
	state     GameState
	Quiet     bool
}

func (m *Manager) Initialize(p Player, q bool) {
	sid, _ := createGame()
	m.sessionId = sid
	newState, err := getState(m.sessionId)
	if err != nil {
		panic(err)
	}
	m.state = newState
	m.player = p
	m.Quiet = q
}

func (m *Manager) StartGame() {
	for true {
		m.player.SetState(m.state)
		nextHand := m.player.NextHand()
		if nextHand == Quit {
			break
		}

		nextState, err := sendHand(m.sessionId, nextHand)
		if err != nil {
			panic(err)
		}

		if !m.Quiet {
			showHand(nextHand)
			nextState.showState()
		}

		m.state = nextState
		if nextState.Over || nextState.Won || !nextState.Moved {
			break
		}
	}

	if !m.Quiet {
		fmt.Println("----------  Finish  ----------")
		m.state.showState()
	}
}
