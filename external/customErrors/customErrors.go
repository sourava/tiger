package customErrors

import "errors"

type CustomError struct {
	HTTPStatus int
	Err        error
}

func NewWithMessage(statusCode int, message string) *CustomError {
	return &CustomError{
		HTTPStatus: statusCode,
		Err:        errors.New(message),
	}
}

func NewWithErr(statusCode int, err error) *CustomError {
	return &CustomError{
		HTTPStatus: statusCode,
		Err:        err,
	}
}

func (e *CustomError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return ""
}

func (e *CustomError) GetHTTPStatus() int {
	return e.HTTPStatus
}
