package util

import (
	"fmt"
	"encoding/json"
)

type appError struct{
	Code          string
	Message       string
}

func (err *appError) Error() string {
	str, _ := json.Marshal(err)
	return fmt.Sprintln(string(str))
}

const (
	UnexpectedError = "UNEXPECTED_ERROR"
)

func NewUnexpectedError(err error) error {
	return &appError{
		Code: UnexpectedError,
		Message: err.Error(),
	}
}

func NewError(code string, message string) error {
	return &appError{
		Code: code,
		Message: message,
	}
}