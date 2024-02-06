package helpers

import (
	"fmt"
	"strings"
)

func GenSlugStr(str string, delimeter string) string {
	slug := strings.Replace(strings.ToLower(str), " ", delimeter, -1)
	return slug
}

func GetSlugSalt(size int) string {
	randStr := fmt.Sprint(GetRandInt(0, size))
	randStrLen := len(randStr)
	sizeLen := len(fmt.Sprint(size))

	for l := 1; l < sizeLen-randStrLen; l++ {
		randStr = "0" + randStr
	}

	return randStr
}
