package exceptions

import "log"

func HandleBasicException(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
