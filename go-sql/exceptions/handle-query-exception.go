package exceptions

import "log"

func HandleQueryException(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
