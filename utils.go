package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Creates a new game and returns a session-id.
func createGame() (string, error) {
	resp, err := http.Get(BaseURL + "start/size/4/tiles/2/victory/31/rand/2/json")
	if err != nil {
		return "", err
	}
	url := resp.Request.URL.String()
	slice := strings.Split(url, "/")
	return slice[len(slice)-2], nil
}

func getState(sessionId string) (GameState, error) {
	resp, err := http.Get(BaseURL + "state/" + sessionId + "/json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var state GameState
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &state)
	return state, err
}

func sendHand(sessionId string, hand Hand) (GameState, error) {
	url := BaseURL + "state/" + sessionId + "/move/" + handToString(hand) + "/json"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var state GameState
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &state)
	return state, err
}

func handToString(h Hand) string {
	if h == Up {
		return "0"
	} else if h == Right {
		return "1"
	} else if h == Down {
		return "2"
	} else {
		return "3"
	}
}

func showHand(h Hand) {
	s := "Move to "
	switch h {
	case Up:
		s += "Up"
	case Right:
		s += "Right"
	case Down:
		s += "Down"
	case Left:
		s += "Left"
	}
	fmt.Println(s)
}

func intToHand(i int) Hand {
	if i == 0 {
		return Up
	} else if i == 1 {
		return Right
	} else if i == 2 {
		return Down
	} else if i == 3 {
		return Left
	} else {
		return Quit
	}
}

func (s *GameState) showState() {
	fmt.Println("---------------------")
	for _, row := range s.Grid {
		fmt.Print("|")
		for _, v := range row {
			fmt.Printf("%4d|", v)
		}
		fmt.Println("")
		fmt.Println("---------------------")
	}
	fmt.Printf("points = %d, score = %d\n", s.Points, s.Score)
}

func Max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}
