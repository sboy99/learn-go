package exception

import "net/http"

type HttpException struct {
	StatusCode int
	Message    *string
}

func (e *HttpException) Error() string {
	return *e.Message
}

func BadRequestException(msg *string) error {
	var message *string

	if msg == nil {
		*message = "Bad Request"
	} else {
		message = msg
	}

	return &HttpException{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}
