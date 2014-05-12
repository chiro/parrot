package random

import (
	"math/rand"
	"sync"
)

type Xorshift struct {
	r               []int
	x, y, z, w, len uint32
	generating      sync.Mutex
}

func (x *Xorshift) SetRange(i []int) {
	x.r = i
	x.len = uint32(len(x.r))
	x.x, x.y, x.z, x.w = 123456789, 362436069, 521288629, 88675123
	for j := 0; j < rand.Intn(50)+50; j++ {
		x.GenNext()
	}
	x.generating = sync.Mutex{}
}

func (x *Xorshift) GenNext() {
	x.generating.Lock()
	var t uint32 = x.x ^ (x.x << 11)
	x.x, x.y, x.z = x.y, x.z, x.w
	x.w = (x.w ^ (x.w >> 19)) ^ (t ^ (t >> 8))
	x.generating.Unlock()
}

func (x *Xorshift) GetRandom() int {
	x.GenNext()
	if x.w%10 == 0 {
		return 4
	}
	return 2
}

func (p *Xorshift) GetGenerator() func() uint32 {
	var x, y, z, w uint32 = 123456789, 362436069, 521288629, 88675123
	ret := func() uint32 {
		var t uint32 = x ^ (x << 11)
		x, y, z = y, z, w
		w = (w ^ (w >> 19)) ^ (t ^ (t >> 8))
		return w
	}
	for j := 0; j < 50+rand.Intn(50); j++ {
		ret()
	}
	return ret
}
