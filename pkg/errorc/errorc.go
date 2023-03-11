package errorc

import "net/http"

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func New(code int, message string) error {
	return &Error{
		Status:  code,
		Message: message,
	}
}
func ResponseErr(message string) error {
	return &Error{
		Status:  http.StatusInternalServerError,
		Message: message,
	}
}

func (e *Error) Error() string {
	return e.Message
}
