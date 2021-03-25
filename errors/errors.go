package errors

import (
	"encoding/json"
	"fmt"
	"net"
)

type CustomError struct {
	Code    string
	Message string
	causes  []error
}

type ValidationError struct {
	CustomError
}

type NotFoundError struct {
	CustomError
	Domain     string
	Identifier string
}

type NotFoundErrorWithoutId struct {
	CustomError
	Domain string
}

type CommunicationError struct {
	CustomError
	RequestName string
	StatusCode  int
}

type JsonError struct {
	CustomError
}

type AuthError struct {
	CustomError
}

type DbError struct {
	CustomError
}

type MappingError struct {
	CustomError
}

func (e CustomError) Error() string {
	return e.Message
}

func (e ValidationError) Error() string {
	return e.Message
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf(e.Message, e.Domain, e.Identifier)
}

func (e NotFoundErrorWithoutId) Error() string {
	return fmt.Sprintf(e.Message, e.Domain)
}

func (e CommunicationError) Error() string {
	return fmt.Sprintf(e.Message)
}

func (e JsonError) Error() string {
	return fmt.Sprintf(e.Message)
}

func (a AuthError) Error() string {
	return a.Message
}

func (e DbError) Error() string {
	return e.Message
}

func (e MappingError) Error() string {
	return e.Message
}

func NewValidationErrorFromMessage(message string) error {
	return NewCustomValidationError(message, "")
}

func NewValidationError(message string) error {
	return ValidationError{CustomError{Message: message}}
}

func NewCustomValidationError(message, customError string) error {
	return ValidationError{CustomError{Message: message, Code: customError}}
}

func NewNotFoundError(domain string, identifier string) error {
	return NotFoundError{CustomError{Message: "The %s with identifier %s was not found"}, domain, identifier}
}

func NewNotFoundErrorWithoutId(domain string) error {
	return NotFoundErrorWithoutId{CustomError{Message: "The %s was not found"}, domain}
}

func NewCustomNotFoundError(message, domain, identifier string) error {
	return NotFoundError{CustomError{Message: message}, domain, identifier}
}

func NewCommunicationError(message, requestName string, status int) error {
	return CommunicationError{CustomError: CustomError{Message: message}, RequestName: requestName, StatusCode: status}
}

func NewAuthorizationError(message, code string) error {
	return AuthError{CustomError: CustomError{Message: message, Code: code}}
}

func NewJsonError(message string) error {
	return JsonError{CustomError: CustomError{Message: message}}
}

func NewDbError(message string) error {
	return DbError{CustomError: CustomError{Message: message}}
}

func NewMappingError(message string) error {
	return MappingError{CustomError: CustomError{Message: message}}
}

func TranslateUnmarshalError(err error) error {
	if casted, ok := err.(*json.UnmarshalTypeError); ok {
		return NewCustomValidationError(
			fmt.Sprintf("invalid type (%s) for field: %s", casted.Value, casted.Field),
			"invalid_items")
	} else {
		return NewValidationError(err.Error())
	}
}

func IsTimeout(err error) bool {
	netErr, ok := err.(net.Error)
	return ok && netErr.Timeout()
}
