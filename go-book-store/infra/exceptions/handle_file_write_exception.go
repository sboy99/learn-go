package exceptions

func HandleFileWriteException(err error) {
	if err != nil {
		panic(err)
	}
}
