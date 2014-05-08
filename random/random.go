package random

type Gen interface {
	SetRange([]int)
	GetRandom() int
}
