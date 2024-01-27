package exceptions

import "log"

func HandleReqBodyReadException(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
