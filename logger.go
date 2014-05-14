package main

import (
	"fmt"
)

var isDebug bool = false

func Debug(s string) {
	if isDebug {
		fmt.Println(s)
	}
}
