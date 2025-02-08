package error

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrBadRequest     = errors.New("bad request")
	ErrInternalServer = errors.New("internal server error")
)

func GetErrorStatusCode(err error) int {
	switch {
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, ErrBadRequest):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
