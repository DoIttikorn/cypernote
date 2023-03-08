package errorc

import "net/http"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func New(code int, message string) error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
func ErrServerMessage(message string) error {
	return &Error{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func (e *Error) Error() string {
	return e.Message
}
