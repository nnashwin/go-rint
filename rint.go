package rint

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	math_rand "math/rand"
)

func Init() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func Gen(max int) int {
	return math_rand.Intn(max)
}

func GenRange(min, max int) int {
	return math_rand.Intn(max-min) + min
}
