package exceptions

func HandleJsonUnMarshalException(
	err error,
) {
	if err != nil {
		panic(err)
	}
}
