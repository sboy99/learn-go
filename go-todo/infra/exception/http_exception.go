package exception

import "net/http"

type HttpException struct {
	StatusCode int
	Message    string
}

func (e *HttpException) Error() string {
	return e.Message
}

func BadRequestException(err error) *HttpException {
	var message string = err.Error()

	return &HttpException{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

func InternalServerErrorException(err error) *HttpException {
	var message string = err.Error()

	return &HttpException{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}
