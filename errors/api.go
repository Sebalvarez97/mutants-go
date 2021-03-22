package errors

import "fmt"

const NotFoundCode = 404
const NotFoundMessage = "Not Found because of %s"

const InternalServerErrorCode = 500
const InternalServerErrorMessage = "Server failed to perform request because of %s"

const DBConnectionErrorCode = 500
const DBConnectionErrorMessage = "Server failed to connect/disconnect from db because of %s"

const BadRequestErrorCode = 400
const BadRequestErrorMessage = "Invalid value entered: %s"

type ApiError struct {
	error
	Code    int
	Message string
	Cause   error
}

func (e ApiError) Error() string {
	return e.Message
}

func BadRequestError(err error) ApiError {
	return ApiError{
		Code:    BadRequestErrorCode,
		Message: fmt.Sprintf(BadRequestErrorMessage, err.Error()),
		Cause:   err,
	}
}

func NotFoundError(err error) ApiError {
	return ApiError{
		Code:    NotFoundCode,
		Message: fmt.Sprintf(NotFoundMessage, err.Error()),
		Cause:   err,
	}
}

func DbConnectionError(err error) ApiError {
	return ApiError{
		Code:    DBConnectionErrorCode,
		Message: fmt.Sprintf(DBConnectionErrorMessage, err.Error()),
		Cause:   err,
	}
}

func GenericError(err error) ApiError {
	return ApiError{
		Code:    InternalServerErrorCode,
		Message: fmt.Sprintf(InternalServerErrorMessage, err.Error()),
		Cause:   err,
	}
}
