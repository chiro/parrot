package main

import "math/rand"

func GetNextTile() (ret int) {
	i := rand.Intn(10)
	if i == 4 {
		ret = 4
	} else {
		ret = 2
	}
	return
}
