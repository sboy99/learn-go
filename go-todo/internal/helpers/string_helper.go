package helpers

import (
	"fmt"
	"strings"
)

func GenSlugStr(str string, delimeter *string) string {
	var d string = ""
	if delimeter != nil {
		d = *delimeter
	}

	slug := strings.Replace(strings.ToLower(str), " ", d, -1)
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
