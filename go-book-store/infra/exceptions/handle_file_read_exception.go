package exceptions

func HandleFileReadException(err error) {
	if err != nil {
		panic(err)
	}
}
