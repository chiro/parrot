package random

type Gen interface {
	SetRange([]int)
	GetRandom() int
	GetGenerator() func() uint32
}
