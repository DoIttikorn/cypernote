package errorc

import "net/http"

var (
	ErrUnauthorized        = New(http.StatusUnauthorized, "unauthorized")                 // ErrUnauthorized is returned when a user is unauthorized
	ErrForbidden           = New(http.StatusForbidden, "forbidden")                       // ErrForbidden is returned when a user is forbidden
	ErrBadRequest          = New(http.StatusBadRequest, "bad request")                    // ErrBadRequest is returned when a request is invalid
	ErrConflict            = New(http.StatusConflict, "conflict")                         // ErrConflict is returned when a resource is already exists
	ErrBind                = New(http.StatusBadRequest, "invalid request body")           // ErrBind is returned when a request body is invalid
	ErrNotFound            = New(http.StatusNotFound, "not found")                        // ErrNotFound is returned when a resource is not found
	ErrInvalidID           = New(http.StatusBadRequest, "invalid id")                     // ErrInvalidID is returned when an ID is invalid
	ErrInvalidEmail        = New(http.StatusNotFound, "invalid email")                    // ErrInvalidEmail is returned when an email is invalid
	ErrInvalidPassword     = New(http.StatusNotFound, "invalid password")                 // ErrInvalidPassword is returned when a password is invalid
	ErrInvalidToken        = New(http.StatusInternalServerError, "invalid token")         // ErrInvalidToken is returned when a token is invalid
	ErrInvalidRefreshToken = New(http.StatusInternalServerError, "invalid refresh token") // ErrInvalidToken is returned when a token is invalid
	ErrServer              = New(http.StatusInternalServerError, "internal server error") // ErrServer is returned when an internal server error occurs
)
