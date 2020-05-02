package rint

import (
	"math/rand"
	"time"
)

func GetRint(num int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(num)
}

func GetRintRange(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
