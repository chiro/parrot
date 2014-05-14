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

var moveTable [4][65536]uint64
var scoreTable [4][65536]int

func (m *Midori) Initialize() {
	moveTable = [4][65536]uint64{}
	scoreTable = [4][65536]int{}
	for y := 0; y < 65536; y++ {
		//table := decodeRow(uint64(y))
		var sim Kanna = Kanna{decodeRow(uint64(y)), 0, false, func() uint32 { return uint32(1) }}
		sim.Move(Right)
		moveTable[1][y] = encodeRow(sim.Grid)
		moveTable[2][y] = encodeRow(sim.Grid)
		scoreTable[1][y] = sim.score
		scoreTable[2][y] = sim.score

		sim.Grid = decodeRow(uint64(y))
		sim.score = 0
		sim.Move(Left)
		moveTable[0][y] = encodeRow(sim.Grid)
		moveTable[3][y] = encodeRow(sim.Grid)
		scoreTable[0][y] = sim.score
		scoreTable[3][y] = sim.score
	}
}

func decodeRow(r uint64) [4][4]int {
	row := [4][4]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	for x := 0; x < 4; x++ {
		row[0][x] = int((r >> uint(x*4)) & MASK)
	}
	return row
}

func encodeRow(table [4][4]int) uint64 {
	var ret uint64 = 0
	for x := 0; x < 4; x++ {
		ret |= uint64((uint64(table[0][x]) & MASK) << uint64(x*4))
	}
	return ret
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
	var nr uint64 = 0
	for col := 0; col < 4; col++ {
		var c uint64 = 0
		for y := 0; y < 4; y++ {
			c |= uint64(m.get(col, y) << uint(y*4))
		}
		d := moveTable[0][c]
		m.score += scoreTable[0][c]
		for y := 0; y < 4; y++ {
			nr |= ((d >> uint(y*4)) & MASK) << uint(y*16+col*4)
		}
	}
	moved := nr != m.Grid
	m.Grid = nr
	return moved
}

func (m *Midori) moveRight() bool {
	var nr uint64 = 0
	for row := 0; row < 4; row++ {
		var c uint64 = 0
		for x := 0; x < 4; x++ {
			c |= uint64(m.get(x, row) << uint(x*4))
		}
		d := moveTable[1][c]
		m.score += scoreTable[1][c]
		nr |= d << uint(row*16)
	}
	moved := nr != m.Grid
	m.Grid = nr
	return moved
}

func (m *Midori) moveDown() bool {
	var nr uint64 = 0
	for col := 0; col < 4; col++ {
		var c uint64 = 0
		for y := 0; y < 4; y++ {
			c |= uint64(m.get(col, y) << uint(y*4))
		}
		d := moveTable[2][c]
		m.score += scoreTable[2][c]
		for y := 0; y < 4; y++ {
			nr |= ((d >> uint(y*4)) & MASK) << uint(y*16+col*4)
		}

	}
	moved := nr != m.Grid
	m.Grid = nr
	return moved
}

func (m *Midori) moveLeft() bool {
	var nr uint64 = 0
	for row := 0; row < 4; row++ {
		var c uint64 = 0
		for x := 0; x < 4; x++ {
			c |= uint64(m.get(x, row) << uint(x*4))
		}
		d := moveTable[3][c]
		m.score += scoreTable[3][c]
		nr |= d << uint(row*16)
	}
	moved := nr != m.Grid
	m.Grid = nr
	return moved
}
