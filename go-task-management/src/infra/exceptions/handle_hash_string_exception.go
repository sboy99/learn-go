package exceptions

import "log"

func HandleHashStringException(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
