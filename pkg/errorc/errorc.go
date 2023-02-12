package errorc

type Error struct {
	message string
}

func New(message string) error {
	return &Error{
		message: message,
	}
}

func (e *Error) Error() string {
	return e.message
}
