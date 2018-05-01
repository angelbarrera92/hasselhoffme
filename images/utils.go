package images

import (
	"time"
	"math/rand"
)

func initRand() {
	rand.Seed(time.Now().Unix())
}

func RandomNumber(images []ImageResult) int {
	initRand()

	return rand.Intn(len(images))
}

func RandomNumberInt(min int, max int) int {
	initRand()

	return rand.Intn(max - min) + min
}