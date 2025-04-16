package errors

import "net/http"

func ErrInvalidInput(err error) error {
	return NewDomainError(
		http.StatusBadRequest,
		err,
		"invalid input",
		"ERR_INVALID_INPUT",
	)
}

func ErrUnauthorized() error {
	return NewDomainError(
		http.StatusUnauthorized,
		nil,
		"unauthorized user",
		"ERR_UNAUTHORIZED",
	)
}
