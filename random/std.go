package random

import "math/rand"

type Std struct {
	r   []int
	len int
}

func (s *Std) SetRange(i []int) {
	s.r = i
	s.len = len(s.r)
}

func (s *Std) GetRandom() int {
	i := rand.Intn(s.len)
	return s.r[i]
}

func (s *Std) GetGenerator() func() uint32 {
	return func() uint32 {
		return uint32(rand.Int())
	}
}
