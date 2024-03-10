package errors

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	StatusCode int
	Error      error
}

func BadRequest(format string, args ...any) HttpError {
	return HttpError{
		StatusCode: http.StatusBadRequest,
		Error:      fmt.Errorf(format, args...),
	}
}

func Unauthorized(format string, args ...any) HttpError {
	return HttpError{
		StatusCode: http.StatusUnauthorized,
		Error:      fmt.Errorf(format, args...),
	}
}
