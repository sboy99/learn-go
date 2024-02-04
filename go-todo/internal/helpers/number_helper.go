package helpers

import "math/rand"

func GetRandInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
