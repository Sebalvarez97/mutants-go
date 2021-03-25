package web

import (
	"fmt"
	"net/http"
	"strings"
)

// Error is our custom error's interface implementation.
type Error struct {
	Message    string        `json:"message"`
	Code       string        `json:"error"`
	StatusCode int           `json:"status"`
	Cause      []interface{} `json:"cause"`
}

// Error returns a string message of the error. It is a concatenation of Code and Message fields.
// This means the Error implements the error interface.
func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	if e.Message == "" {
		return e.Code
	}

	if e.Code == "" {
		return e.Message
	}

	return fmt.Sprintf("%d %s: %s", e.StatusCode, e.Code, e.Message)
}

// NewError creates a new error with the given status code and message.
func NewError(statusCode int, message string) error {
	return NewErrorf(statusCode, message)
}

// NewErrorf creates a new error with the given status code and the message
// formatted according to args and format.
func NewErrorf(statusCode int, format string, args ...interface{}) error {
	return &Error{
		Code:       strings.ReplaceAll(strings.ToLower(http.StatusText(statusCode)), " ", "_"),
		Message:    fmt.Sprintf(format, args...),
		StatusCode: statusCode,
	}
}
