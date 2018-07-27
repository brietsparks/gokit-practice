package util

import "fmt"

type appError struct{
	code          string
	message       string
}

func (err *appError) Error() string {
	return fmt.Sprintln("error " + err.code + ": " + err.message)
}

const (
	UnexpectedError = "UNEXPECTED_ERROR"
)

func NewUnexpectedError(err error) error {
	return &appError{
		code: UnexpectedError,
		message: err.Error(),
	}
}

func NewError(code string, message string) error {
	return &appError{
		code: code,
		message: message,
	}
}