package errorc

import "net/http"

var (
	// ErrUnauthorized is returned when a user is unauthorized
	ErrUnauthorized = New(http.StatusUnauthorized, "unauthorized")
	// ErrForbidden is returned when a user is forbidden
	ErrForbidden = New(http.StatusForbidden, "forbidden")
	// ErrBadRequest is returned when a request is invalid
	ErrBadRequest = New(http.StatusBadRequest, "bad request")
	// ErrConflict is returned when a resource is already exists
	ErrConflict = New(http.StatusConflict, "conflict")
	// ErrBind is returned when a request body is invalid
	ErrBind = New(http.StatusBadRequest, "invalid request body")
	// ErrNotFound is returned when a resource is not found
	ErrNotFound = New(http.StatusNotFound, "not found")
	// ErrInvalidID is returned when an ID is invalid
	ErrInvalidID = New(http.StatusBadRequest, "invalid id")
	// ErrInvalidEmail is returned when an email is invalid
	ErrInvalidEmail = New(http.StatusNotFound, "invalid email")
	// ErrInvalidPassword is returned when a password is invalid
	ErrInvalidPassword = New(http.StatusNotFound, "invalid password")
	// ErrInvalidToken is returned when a token is invalid
	ErrInvalidToken = New(http.StatusInternalServerError, "invalid token")
	// ErrInvalidToken is returned when a token is invalid
	ErrInvalidRefreshToken = New(http.StatusInternalServerError, "invalid refresh token")
	// ErrServer is returned when an internal server error occurs
	ErrServer = New(http.StatusInternalServerError, "internal server error")
)
