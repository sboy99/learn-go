package exceptions

import "log"

func HandleDBConnectionException(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
