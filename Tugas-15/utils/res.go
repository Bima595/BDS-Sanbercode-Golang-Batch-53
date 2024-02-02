package utils

import "errors"

// NewBadRequestError creates a new bad request error with the given message.
func NewBadRequestError(message string) error {
	return errors.New("Bad Request: " + message)
}
