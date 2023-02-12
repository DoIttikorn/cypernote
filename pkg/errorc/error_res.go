package errorc

// Path: pkg/errorc/error_res.go
type ErrorRes struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrorRes(code int, message string) *ErrorRes {
	return &ErrorRes{
		Code:    code,
		Message: message,
	}
}
