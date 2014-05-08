package main

import (
	"github.com/chiro/parrot/random"
	"testing"
)

func TestMoveUp(t *testing.T) {
	var initial [4][4]int = [4][4]int{{0, 2, 0, 2}, {2, 2, 0, 0}, {0, 2, 0, 0}, {0, 0, 0, 0}}
	var r random.Gen = &random.Std{}
	var sim Kanna = Kanna{initial, 0, false, r}
	sim.Move(Up)

	var expected [4][4]int = [4][4]int{{2, 4, 0, 2}, {0, 2, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	if sim.Grid != expected {
		t.Errorf("got %v\nwant %v", sim.Grid, expected)
	}
}

func TestRotateRight(t *testing.T) {
	var initial [4][4]int = [4][4]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}, {12, 13, 14, 15}}
	var r random.Gen = &random.Std{}
	var sim Kanna = Kanna{initial, 0, false, r}
	sim.rotateRight()

	var expected [4][4]int = [4][4]int{{12, 8, 4, 0}, {13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}}
	if sim.Grid != expected {
		t.Errorf("got %v\nwant %v", sim.Grid, expected)
	}
}

func TestRotateLeft(t *testing.T) {
	var initial [4][4]int = [4][4]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}, {12, 13, 14, 15}}
	var r random.Gen = &random.Std{}
	var sim Kanna = Kanna{initial, 0, false, r}
	sim.rotateLeft()

	var expected [4][4]int = [4][4]int{{3, 7, 11, 15}, {2, 6, 10, 14}, {1, 5, 9, 13}, {0, 4, 8, 12}}
	if sim.Grid != expected {
		t.Errorf("got %v\nwant %v", sim.Grid, expected)
	}
}
