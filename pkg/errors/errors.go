package errors

import "errors"

var (
	ServerError    = errors.New("internal server error please contact an administrator")
	InvalidData    = errors.New("invalid data")
	InvalidRequest = errors.New("invalid request")
	RecordNotFound = errors.New("record not found")

	InvalidID       = errors.New("invalid id")
	InvalidPassword = errors.New("invalid password")
)

func New(text string) error {
	return errors.New(text)
}
