package main

import (
	"math/rand"
	//	"fmt"
)

var freezed [4][4]bool

type Simulator struct {
	Grid   [4][4]int
	Score  int
	Finish bool
}

func (s *Simulator) move(h Hand) bool {
	if s.Finish {
		return false
	}

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			freezed[y][x] = false
		}
	}

	switch h {
	case Up:
		s.moveUp()
	case Right:
		s.moveRight()
	case Left:
		s.moveLeft()
	case Down:
		s.moveDown()
	}
	x, y := s.getRandomAvailableCell()
	if x == -1 {
		return false
	}

	s.Grid[y][x] = ((rand.Int() % 2) + 1) * 2
	return true
}

func (s *Simulator) moveUp() {
	for y, row := range s.Grid {
		if y == 0 {
			continue
		}

		for x, tile := range row {
			if tile != 0 {
				var merged bool = false
				// Search the next tile
				for yy := y - 1; yy >= 0; yy-- {
					if s.Grid[yy][x] == 0 {
						continue
					} else if freezed[yy][x] || s.Grid[yy][x] != s.Grid[y][x] {
						// can't merge
						tmp := s.Grid[y][x]
						s.Grid[y][x] = 0
						s.Grid[yy+1][x] = tmp
						merged = true
						break
					} else {
						// merge
						s.Score += s.Grid[y][x] * 2
						s.Grid[yy][x] += s.Grid[y][x]
						s.Grid[y][x] = 0
						merged = true
						freezed[yy][x] = true
						break
					}
				}
				if !merged {
					s.Grid[0][x] = s.Grid[y][x]
					s.Grid[y][x] = 0
				}
			}
		}
	}
}

func (s *Simulator) rotateRight() {
	var tmp [4][4]int
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			tmp[x][3-y] = s.Grid[y][x]
		}
	}
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			s.Grid[y][x] = tmp[y][x]
		}
	}
}

func (s *Simulator) rotateLeft() {
	var tmp [4][4]int
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			tmp[3-x][y] = s.Grid[y][x]
		}
	}
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			s.Grid[y][x] = tmp[y][x]
		}
	}
}

func (s *Simulator) moveRight() {
	s.rotateLeft()
	s.moveUp()
	s.rotateRight()
}

func (s *Simulator) moveDown() {
	s.rotateLeft()
	s.rotateLeft()
	s.moveUp()
	s.rotateRight()
	s.rotateRight()
}

func (s *Simulator) moveLeft() {
	s.rotateRight()
	s.moveUp()
	s.rotateLeft()
}

func (s *Simulator) getAvailableCells() (cnt int) {
	cnt = 0
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if s.Grid[y][x] == 0 {
				cnt++
			}
		}
	}
	return
}

func (s *Simulator) getRandomAvailableCell() (x, y int) {
	cnt := s.getAvailableCells()
	if cnt == 0 {
		return -1, -1
	}

	var sel = rand.Int() % cnt
	for y, row := range s.Grid {
		for x, v := range row {
			if v == 0 {
				if sel == 0 {
					return x, y
				}
				sel--
			}
		}
	}
	return -1, -1
}
