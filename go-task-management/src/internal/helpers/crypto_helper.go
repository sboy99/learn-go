package helpers

import "math/rand"

func GetId() int64 {
	return rand.Int63()
}
