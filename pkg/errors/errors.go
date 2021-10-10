package errors

import "errors"

var (
	ErrServerError    = errors.New("internal server error please contact an administrator")
	ErrInvalidData    = errors.New("invalid data")
	ErrInvalidRequest = errors.New("invalid request")
	ErrRecordNotFound = errors.New("record not found")

	ErrInvalidID       = errors.New("invalid id")
	ErrInvalidPassword = errors.New("invalid password")
)

func New(text string) error {
	return errors.New(text)
}
