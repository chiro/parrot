package main

import (
	"math/rand"
)

const MASK uint64 = 15

type Midori struct {
	Grid     uint64
	score    int
	gameover bool
	gen      func() uint32
}

func (m *Midori) Score() int {
	return m.score
}

func popcount(x uint64) uint64 {
	x = (x & 0x5555555555555555) + ((x & 0xAAAAAAAAAAAAAAAA) >> 1)
	x = (x & 0x3333333333333333) + ((x & 0xCCCCCCCCCCCCCCCC) >> 2)
	x = (x & 0x0F0F0F0F0F0F0F0F) + ((x & 0xF0F0F0F0F0F0F0F0) >> 4)

	return x * 0x0101010101010101
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
			if m.get(x, y) == 0 {
				cnt++
			}
		}
	}
	return
}

func (m *Midori) getRandomAvailableCell() (int, int) {
	available := m.GetAvailableCells()
	if available == 0 {
		return -1, -1
	}

	p := rand.Intn(available)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if m.get(j, i) == 0 {
				p--
				if p == 0 {
					return j, i
				}
			}
		}
	}
	return -1, -1
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
	if m.gen()%10 == 1 {
		m.set(x, y, 2)
	} else {
		m.set(x, y, 1)
	}
	return true
}

func log(x int) uint64 {
	switch x {
	case 2:
		return 1
	case 4:
		return 2
	case 8:
		return 3
	case 16:
		return 4
	case 32:
		return 5
	case 64:
		return 6
	case 128:
		return 7
	case 256:
		return 8
	case 512:
		return 9
	case 1024:
		return 10
	case 2048:
		return 11
	case 4096:
		return 12
	case 8192:
		return 13
	}
	return 0
}

func encode2(grid *[4][4]int) (ret uint64) {
	ret = uint64(0)
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			ret |= log(grid[y][x]) << uint((y*4+x)*4)
		}
	}
	return
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

func (m *Midori) GetBoard() [4][4]int {
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
			if m.get(col, i) == 0 || p == i {
				continue
			}

			if m.get(col, p) == 0 {
				// fmt.Printf("moved to empty cell! %d,%d -> %d,%d\n", col, i, col, p)
				m.set(col, p, m.get(col, i))
				m.set(col, i, 0)
				p++
				moved = true
			} else if m.get(col, p) == m.get(col, i) {
				// fmt.Printf("merge! %d,%d -> %d,%d\n", col, i, col, p)
				m.score += 1 << uint(m.get(col, i)+1)
				m.set(col, i, 0)
				m.set(col, p, m.get(col, p)+1)
				p++
				moved = true
			} else {
				p++
				// fmt.Printf("moved! %d,%d -> %d,%d\n", col, i, col, p)
				m.set(col, p, m.get(col, i))
				if p < i {
					m.set(col, i, 0)
					moved = true
				}
			}
		}
	}
	return moved
}

func (m *Midori) moveRight() bool {
	moved := false
	for row := 0; row < 4; row++ {
		p := 3
		for i := 3; i >= 0; i-- {
			if m.get(i, row) == 0 || p == i {
				continue
			}

			if m.get(p, row) == 0 {
				m.set(p, row, m.get(i, row))
				m.set(i, row, 0)
				p--
				moved = true
			} else if m.get(p, row) == m.get(i, row) {
				m.score += 1 << uint(m.get(i, row)+1)
				m.set(i, row, 0)
				m.set(p, row, m.get(p, row)+1)
				p--
				moved = true
			} else {
				p--
				m.set(p, row, m.get(i, row))
				if p != i {
					m.set(i, row, 0)
					moved = true
				}
			}
		}
	}
	return moved
}

func (m *Midori) moveDown() bool {
	moved := false
	for col := 0; col < 4; col++ {
		p := 3
		for i := 3; i >= 0; i-- {
			if m.get(col, i) == 0 || p == i {
				continue
			}

			if m.get(col, p) == 0 {
				// fmt.Printf("moved to empty cell! %d,%d -> %d,%d\n", col, i, col, p)
				m.set(col, p, m.get(col, i))
				m.set(col, i, 0)
				p--
				moved = true
			} else if m.get(col, p) == m.get(col, i) {
				// fmt.Printf("merge! %d,%d -> %d,%d\n", col, i, col, p)
				m.score += 1 << uint(m.get(col, p)+1)
				m.set(col, i, 0)
				m.set(col, p, m.get(col, p)+1)
				p--
				moved = true
			} else {
				p--
				// fmt.Printf("moved! %d,%d -> %d,%d\n", col, i, col, p)
				m.set(col, p, m.get(col, i))
				if p != i {
					m.set(col, i, 0)
					moved = true
				}
			}
		}
	}
	return moved
}

func (m *Midori) moveLeft() bool {
	moved := false
	for row := 0; row < 4; row++ {
		p := 0
		for i := 0; i < 4; i++ {
			if m.get(i, row) == 0 || p == i {
				continue
			}

			if m.get(p, row) == 0 {
				m.set(p, row, m.get(i, row))
				m.set(i, row, 0)
				p++
				moved = true
			} else if m.get(p, row) == m.get(i, row) {
				m.score += 1 << uint(m.get(i, row)+1)
				m.set(i, row, 0)
				m.set(p, row, m.get(p, row)+1)
				p++
				moved = true
			} else {
				p++
				m.set(p, row, m.get(i, row))
				if p != i {
					moved = true
					m.set(i, row, 0)
				}
			}
		}
	}
	return moved
}
