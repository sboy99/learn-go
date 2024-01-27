package helpers

import (
	"fmt"
	"strconv"
)

func ParseInt64(s string) int64 {
	sInt64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		// Handle the error if the conversion fails
		fmt.Println("Error:", err)
		return 0
	}

	return sInt64
}
