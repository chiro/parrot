package main

import (
	"github.com/chiro/parrot/random"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	var initial [4][4]int = [4][4]int{{0, 1, 0, 1}, {1, 1, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}}
	var r random.Gen = &random.Std{}
	var sim Midori = Midori{encode(&initial), 0, false, r}
	if initial != sim.GetState() {
		t.Errorf("got %v\nwant %v\n", sim.GetState(), initial)
	}
}

func TestMidoriGet(t *testing.T) {
	var initial [4][4]int = [4][4]int{{0, 1, 0, 1}, {1, 1, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}}
	var r random.Gen = &random.Std{}
	var sim Midori = Midori{encode(&initial), 0, false, r}

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if sim.get(x, y) != initial[y][x] {
				t.Errorf("got %v, want %v, at (%v,%v)\n", sim.get(x, y), initial[y][x], x, y)
			}
		}
	}
}

func TestMidoriSet(t *testing.T) {
	var initial [4][4]int = [4][4]int{{0, 1, 0, 1}, {1, 1, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}}
	var r random.Gen = &random.Std{}
	var sim Midori = Midori{encode(&initial), 0, false, r}

	sim.set(0, 0, 2)
	if sim.get(0, 0) != 2 {
		t.Errorf("got %v, want %v\n", sim.get(0, 0), 2)
	}

	sim.set(3, 3, 8)
	if sim.get(3, 3) != 8 {
		t.Errorf("got %v, want %v\n", sim.get(3, 3), 8)
	}
}

func TestMidoriMoveUp(t *testing.T) {
	var initial [4][4]int = [4][4]int{{0, 1, 0, 1}, {1, 1, 0, 2}, {0, 1, 0, 3}, {0, 0, 0, 4}}
	var r random.Gen = &random.Std{}
	r.SetRange([]int{2, 4})
	var sim Midori = Midori{encode(&initial), 0, false, r}
	sim.Move(Up)

	var expected [4][4]int = [4][4]int{{1, 2, 0, 1}, {0, 1, 0, 2}, {0, 0, 0, 3}, {0, 0, 0, 4}}
	if sim.Grid != encode(&expected) {
		t.Errorf("got %v\n", sim.GetState())
		t.Errorf("got %v\nwant %v", sim.Grid, encode(&expected))
	}
}

func TestMidoriMoveRight(t *testing.T) {
	var initial [4][4]int = [4][4]int{{1, 1, 4, 5}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	var r random.Gen = &random.Std{}
	r.SetRange([]int{2, 4})
	var sim Midori = Midori{encode(&initial), 0, false, r}
	sim.Move(Right)

	var expected [4][4]int = [4][4]int{{0, 2, 4, 5}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	if sim.Grid != encode(&expected) {
		t.Errorf("got %v\n", sim.GetState())
		t.Errorf("got %v\nwant %v", sim.Grid, encode(&expected))
	}
}

func TestMidoriGetAvailableCells(t *testing.T) {
	var initial [4][4]int = [4][4]int{{0, 1, 0, 1}, {1, 1, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}}
	var r random.Gen = &random.Std{}
	r.SetRange([]int{2, 4})
	var sim Midori = Midori{encode(&initial), 0, false, r}

	if sim.GetAvailableCells() != 11 {
		t.Errorf("got %v, want %v\n", sim.GetAvailableCells(), 11)
	}

	x, _ := sim.getRandomAvailableCell()
	if x == -1 {
		t.Errorf("can't find empty cells!\n")
	}
}

func TestMidoriGetMaxCell(t *testing.T) {
	var initial [4][4]int = [4][4]int{{0, 1, 0, 1}, {1, 1, 0, 0}, {0, 1, 0, 0}, {0, 0, 3, 0}}
	var r random.Gen = &random.Std{}
	r.SetRange([]int{2, 4})
	var sim Midori = Midori{encode(&initial), 0, false, r}

	if sim.GetMaxTile() != 8 {
		t.Errorf("got %v, want %v\n", sim.GetMaxTile(), 8)
	}
}
