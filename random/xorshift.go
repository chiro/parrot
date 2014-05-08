package random

import (
	"math/rand"
)

type Xorshift struct {
	r               []int
	x, y, z, w, len uint32
}

func (x *Xorshift) SetRange(i []int) {
	x.r = i
	x.len = uint32(len(x.r))
	x.x, x.y, x.z, x.w = 123456789, 362436069, 521288629, 88675123
	for j := 0; j < rand.Intn(50)+50; j++ {
		x.GenNext()
	}
}

func (x *Xorshift) GenNext() uint32 {
	var t uint32 = x.x ^ (x.x << 11)
	x.x, x.y, x.z = x.y, x.z, x.w
	x.w = (x.w ^ (x.w >> 19)) ^ (t ^ (t >> 8))
	return x.w
}

func (x *Xorshift) GetRandom() int {
	return x.r[x.GenNext()%x.len]
}
