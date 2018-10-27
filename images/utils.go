package images

import (
	"math/rand"
	"time"
)

func initRand() {
	rand.Seed(time.Now().Unix())
}

// RandomNumber returns a randomly selected integer constrained by the length of the
// []ImageResult array passed to it
func RandomNumber(images []ImageResult) int {
	initRand()

	return rand.Intn(len(images))
}

// RandomNumberInt returns a randomly generated integeger constrained by the
// min and max int values passed to it.
func RandomNumberInt(min int, max int) int {
	initRand()

	return rand.Intn(max-min) + min
}
