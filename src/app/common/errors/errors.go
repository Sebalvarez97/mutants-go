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

type ApiErrorImpl struct {
	Code    int
	Message string
	Cause   error
}

func (e ApiErrorImpl) Error() string {
	return e.Message
}

func BadRequestError(err error) ApiErrorImpl {
	return ApiErrorImpl{
		Code:    BadRequestErrorCode,
		Message: fmt.Sprintf(BadRequestErrorMessage, err.Error()),
		Cause:   err,
	}
}

func NotFoundError(err error) ApiErrorImpl {
	return ApiErrorImpl{
		Code:    NotFoundCode,
		Message: fmt.Sprintf(NotFoundMessage, err.Error()),
		Cause:   err,
	}
}

func DbConnectionError(err error) ApiErrorImpl {
	return ApiErrorImpl{
		Code:    DBConnectionErrorCode,
		Message: fmt.Sprintf(DBConnectionErrorMessage, err.Error()),
		Cause:   err,
	}
}

func GenericError(err error) ApiErrorImpl {
	return ApiErrorImpl{
		Code:    InternalServerErrorCode,
		Message: fmt.Sprintf(InternalServerErrorMessage, err.Error()),
		Cause:   err,
	}
}
