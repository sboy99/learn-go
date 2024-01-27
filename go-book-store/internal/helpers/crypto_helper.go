package helpers

import "math/rand"

func GenRandomId() int64 {
	return rand.Int63()
}
