package main

type Manager struct {
	player Player
	sessionId string
	state GameState
}

func (m *Manager) Initialize(p Player) {
	sid , _ := createGame()
	m.sessionId = sid
	newState, err := getState(m.sessionId)
	if err != nil {
		panic(err)
	}
	m.state = newState
	m.player = p
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

		m.state = nextState
		if nextState.Over || nextState.Won || !nextState.Moved {
			break
		}
	}
}
