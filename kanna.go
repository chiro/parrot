package main

import (
	"math/rand"
	//	"fmt"
)

var freezed [4][4]bool

type Kanna struct {
	Grid     [4][4]int
	score    int
	gameover bool
}

func (s *Kanna) Score() int {
	return s.score
}

func (s *Kanna) Move(h Hand) bool {
	if s.gameover {
		return false
	}

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			freezed[y][x] = false
		}
	}

	var moved bool = false
	switch h {
	case Up:
		moved = s.moveUp()
	case Right:
		moved = s.moveRight()
	case Left:
		moved = s.moveLeft()
	case Down:
		moved = s.moveDown()
	}
	x, y := s.getRandomAvailableCell()
	if x == -1 || !moved {
		return false
	}

	s.Grid[y][x] = GetNextTile()
	return true
}

func (s *Kanna) moveUp() bool {
	var moved bool = false
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
						moved = true
						continue
					} else if freezed[yy][x] || s.Grid[yy][x] != s.Grid[y][x] {
						// can't merge
						tmp := s.Grid[y][x]
						s.Grid[y][x] = 0
						s.Grid[yy+1][x] = tmp
						merged = true
						if yy+1 != y {
							moved = true
						}
						break
					} else {
						// merge
						s.score += s.Grid[y][x] * 2
						s.Grid[yy][x] += s.Grid[y][x]
						s.Grid[y][x] = 0
						merged = true
						freezed[yy][x] = true
						moved = true
						break
					}
				}
				if !merged {
					moved = true
					s.Grid[0][x] = s.Grid[y][x]
					s.Grid[y][x] = 0
				}
			}
		}
	}
	return moved
}

func (s *Kanna) rotateRight() {
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

func (s *Kanna) rotateLeft() {
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

func (s *Kanna) moveRight() bool {
	s.rotateLeft()
	moved := s.moveUp()
	s.rotateRight()
	return moved
}

func (s *Kanna) moveDown() bool {
	s.rotateLeft()
	s.rotateLeft()
	moved := s.moveUp()
	s.rotateRight()
	s.rotateRight()
	return moved
}

func (s *Kanna) moveLeft() bool {
	s.rotateRight()
	moved := s.moveUp()
	s.rotateLeft()
	return moved
}

func (s *Kanna) GetAvailableCells() (cnt int) {
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

func (s *Kanna) getRandomAvailableCell() (x, y int) {
	cnt := s.GetAvailableCells()
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
