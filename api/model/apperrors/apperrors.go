package apperrors

import (
	"errors"
	"fmt"
	"net/http"
)

// Type holds a type string AND integer code for the error
type Type string

// The set of valid error types

const (
	Authorization   Type = "AUTHORIZATION"
	BadRequest      Type = "BADREQUEST"
	Conflict        Type = "CONFLICT"
	Internal        Type = "INTERNAL"
	NotFound        Type = "NOTFOUND"
	PayloadTooLarge Type = "PAYLOADTOOLARGE"
)

// Error is a custom error handling struct
// Which makes it easier to implement consistent
// Error responses throughout the api
type Error struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
}

// Error satisfies the requirements needed
// to implement golangs built in error interface
// we can return errors from this package as
// we would like normal golang errors
func (e *Error) Error() string {
	return e.Message
}

// Status maps errors to status codes, a little redundancy
// placed on an Error type, this is a receiver
func (e *Error) Status() int {
	switch e.Type {
	case Authorization:
		return http.StatusUnauthorized
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	case PayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	default:
		return http.StatusInternalServerError
	}
}

// Status (non receiver) checks runtime type of the error
// returns an http status code if the error is of this package
func Status(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Status()
	}

	// Returns 500 if the error is not of this package
	return http.StatusInternalServerError
}

// Error "Factories"

// NewAuthorizationErr creates authorization error
func NewAuthorizationErr(reason string) *Error {
	return &Error{
		Type:    Authorization,
		Message: reason,
	}
}

// NewBadRequestErr creates bad request error
func NewBadRequestErr(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad request. Reason: %v", reason),
	}
}

// NewConflictErr creates conflict error
func NewConflictErr(name string, value string) *Error {
	return &Error{
		Type:    Conflict,
		Message: fmt.Sprintf("resource %v with value %v already exists", name, value),
	}
}

// NewInternalErr creates an internal server error
func NewInternalErr() *Error {
	return &Error{
		Type:    Internal,
		Message: fmt.Sprintf("Internal server error."),
	}
}

// NewNotFoundErr creates NewNotFound error
func NewNotFoundErr(name string, value string) *Error {
	return &Error{
		Type:    NotFound,
		Message: fmt.Sprintf("resource %v with value %v not found / unknown", name, value),
	}
}

// NewPayloadTooLargeErr creates payload too large error
func NewNotFound(maxBodySize int64, contentLength int64) *Error {
	return &Error{
		Type:    PayloadTooLarge,
		Message: fmt.Sprintf("Max payload size of %v exceeded. payload: %v", maxBodySize, contentLength),
	}
}
