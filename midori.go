package main

import (
	"github.com/chiro/parrot/random"
	"math/rand"
)

const MASK uint64 = 15

type Midori struct {
	Grid     uint64
	score    int
	gameover bool
	rand     random.Gen
}

func (s *Midori) Initialize() {
	s.rand.SetRange([]int{1, 2})
}

func (m *Midori) Score() int {
	return m.score
}

func (m *Midori) GetMaxTile() int {
	g := m.Grid
	var p uint = 0
	for g > 0 {
		p = uint(Max(int(p), int(g&MASK)))
		g >>= 4
	}
	return 1 << p
}

func (m *Midori) GetAvailableCells() (cnt int) {
	cnt = 0
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if m.get(x, y) > 0 {
				cnt++
			}
		}
	}
	return
}

func (m *Midori) getRandomAvailableCell() (x, y int) {
	x, y = -1, -1
	available := m.GetAvailableCells()
	if available == 0 {
		return
	}

	p := rand.Intn(available)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if m.get(j, i) > 0 {
				p--
				if p == 0 {
					x, y = j, i
					return
				}
			}
		}
	}
	panic("!!!")
}

func (m *Midori) get(x, y int) int {
	return int((m.Grid >> uint((y*4+x)*4)) & MASK)
}

func (m *Midori) set(x, y, p int) {
	m.Grid = m.Grid&(^(MASK << uint((y*4+x)*4))) | (uint64(p) << uint64((y*4+x)*4))
}

func (m *Midori) Move(h Hand) bool {
	if m.gameover {
		return false
	}

	switch h {
	case Up:
		return m.moveUp()
	case Right:
		return m.moveRight()
	case Down:
		return m.moveDown()
	case Left:
		return m.moveLeft()
	}
	return false
}

func (m *Midori) AddRandomCell() bool {
	x, y := m.getRandomAvailableCell()
	if x == -1 {
		return false
	}
	m.set(x, y, m.rand.GetRandom())
	return true
}

func encode(grid *[4][4]int) (ret uint64) {
	ret = uint64(0)
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			ret |= (uint64(grid[y][x]) << uint((y*4+x)*4))
		}
	}
	return
}

func (m *Midori) GetState() [4][4]int {
	r := [4][4]int{}
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			r[y][x] = m.get(x, y)
		}
	}
	return r
}

func (m *Midori) moveUp() bool {
	moved := false
	for col := 0; col < 4; col++ {
		p := 0
		for i := 0; i < 4; i++ {
			if m.get(col, i) == 0 {
				continue
			}
			if p == i {
				continue
			}

			moved = true
			if m.get(col, p) == 0 {
				// fmt.Printf("moved to empty cell! %d,%d -> %d,%d\n", col, i, col, p)
				m.set(col, p, m.get(col, i))
				m.set(col, i, 0)
				p++
			} else if m.get(col, p) == m.get(col, i) {
				// fmt.Printf("merge! %d,%d -> %d,%d\n", col, i, col, p)
				m.set(col, i, 0)
				m.set(col, p, m.get(col, p)+1)
				p++
			} else {
				p++
				// fmt.Printf("moved! %d,%d -> %d,%d\n", col, i, col, p)
				m.set(col, p, m.get(col, i))
				m.set(col, i, 0)
			}
		}
	}
	return moved
}

func (m *Midori) moveRight() bool {
	return false
}

func (m *Midori) moveDown() bool {
	return false
}

func (m *Midori) moveLeft() bool {
	return false
}
