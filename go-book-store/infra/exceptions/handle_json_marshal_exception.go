package exceptions

func HandleJsonMarshalException(
	err error,
) {
	if err != nil {
		panic(err)
	}
}
