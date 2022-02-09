package images

import (
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

// RandomNumber returns a randomly selected integer constrained by the length of the
// []ImageResult array passed to it
func RandomNumber(images []ImageResult) int {
	rand.Seed(time.Now().Unix())

	return rand.Intn(len(images))
}

func createTempFile(content []byte) (*os.File, error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "hasselhoffme-")
	if err != nil {
		return nil, err
	}
	if _, err = tmpFile.Write(content); err != nil {
		return nil, err
	}
	if err := tmpFile.Close(); err != nil {
		return nil, err
	}
	return tmpFile, nil
}
