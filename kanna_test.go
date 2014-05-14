package main

import (
	"github.com/chiro/parrot/random"
	"testing"
)

func TestKannaAddRandomCell(t *testing.T) {
	var initial [4][4]int = [4][4]int{{0, 2, 0, 2}, {2, 2, 0, 0}, {0, 2, 0, 0}, {0, 0, 0, 0}}
	var r random.Gen = &random.Std{}
	var sim Kanna = Kanna{initial, 0, false, r.GetGenerator()}

	before := sim.GetAvailableCells()
	sim.AddRandomCell()
	after := sim.GetAvailableCells()
	if before-1 != after {
		t.Errorf("got %v, want %v\n", after, before-1)
	}
}

func TestMoveUp(t *testing.T) {
	var initial [4][4]int = [4][4]int{{0, 2, 0, 2}, {2, 2, 0, 0}, {0, 2, 0, 0}, {0, 0, 0, 0}}
	var r random.Gen = &random.Std{}
	var sim Kanna = Kanna{initial, 0, false, r.GetGenerator()}
	sim.Move(Up)

	var expected [4][4]int = [4][4]int{{2, 4, 0, 2}, {0, 2, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	if sim.Grid != expected {
		t.Errorf("got %v\nwant %v", sim.Grid, expected)
	}
}

func TestMoveUp2(t *testing.T) {
	var initial [4][4]int = [4][4]int{{8, 8, 4, 256}, {2, 0, 128, 128}, {0, 0, 2, 2}, {0, 0, 4, 8}}
	var r random.Gen = &random.Xorshift{}
	var gen func() uint32 = r.GetGenerator()
	var sim Kanna = Kanna{initial, 0, false, gen}
	p := sim.Move(Up)

	if p {
		t.Errorf("got %v\nwant %v\n", sim.Grid, p)
	}
}

func TestRotateRight(t *testing.T) {
	var initial [4][4]int = [4][4]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}, {12, 13, 14, 15}}
	var r random.Gen = &random.Std{}
	r.SetRange([]int{2, 4})
	var sim Kanna = Kanna{initial, 0, false, r.GetGenerator()}
	sim.rotateRight()

	var expected [4][4]int = [4][4]int{{12, 8, 4, 0}, {13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}}
	if sim.Grid != expected {
		t.Errorf("got %v\nwant %v", sim.Grid, expected)
	}
}

func TestRotateLeft(t *testing.T) {
	var initial [4][4]int = [4][4]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}, {12, 13, 14, 15}}
	var r random.Gen = &random.Std{}
	r.SetRange([]int{2, 4})
	var sim Kanna = Kanna{initial, 0, false, r.GetGenerator()}
	sim.rotateLeft()

	var expected [4][4]int = [4][4]int{{3, 7, 11, 15}, {2, 6, 10, 14}, {1, 5, 9, 13}, {0, 4, 8, 12}}
	if sim.Grid != expected {
		t.Errorf("got %v\nwant %v", sim.Grid, expected)
	}
}

func TestMoveLeft(t *testing.T) {
	var initial [4][4]int = [4][4]int{{1, 1, 1, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	var r random.Gen = &random.Std{}
	var sim Kanna = Kanna{initial, 0, false, r.GetGenerator()}
	sim.moveLeft()
	var expected [4][4]int = [4][4]int{{2, 1, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	if sim.Grid != expected {
		t.Errorf("got %v\nwant %v\n", sim.Grid, expected)
	}
}

func TestMoveUp3(t *testing.T) {
	var initial [4][4]int = [4][4]int{{1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {0, 0, 0, 0}}
	var r random.Gen = &random.Std{}
	var sim Kanna = Kanna{initial, 0, false, r.GetGenerator()}
	sim.moveUp()
	var expected [4][4]int = [4][4]int{{2, 0, 0, 0}, {1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	if sim.Grid != expected {
		t.Errorf("got %v\nwant %v\n", sim.Grid, expected)
	}
}
