package auth

import (
	"fmt"
)

// ErrorNewRequest wraps errors returned by http.NewRequest in the Auth struct methods.
type ErrorNewRequest struct {
	err error
}

func newErrorNewRequest(err error) ErrorNewRequest {
	return ErrorNewRequest{
		err: fmt.Errorf("auth: new request: %w", err),
	}
}

func (e ErrorNewRequest) Error() string {
	return e.err.Error()
}

// ErrorClientDo wraps errors returned by http.Client.Do in the Auth struct methods.
type ErrorClientDo struct {
	err error
}

func newErrorClientDo(err error) ErrorClientDo {
	return ErrorClientDo{
		err: fmt.Errorf("auth: client do: %w", err),
	}
}

func (e ErrorClientDo) Error() string {
	return e.err.Error()
}

// ErrorReadResponseBody wraps errors returned by ioutil.ReadAll in the Auth struct methods.
type ErrorReadResponseBody struct {
	err error
}

func newErrorReadResponseBody(err error) ErrorReadResponseBody {
	return ErrorReadResponseBody{
		err: fmt.Errorf("auth: read all: %w", err),
	}
}

func (e ErrorReadResponseBody) Error() string {
	return e.err.Error()
}

// ErrorStatusUnauthorized wraps an error returned when a 401 status is received in the UpdateUserToken method.
type ErrorStatusUnauthorized struct {
	err error
}

func newErrorStatusUnauthorized(message string) ErrorStatusUnauthorized {
	return ErrorStatusUnauthorized{
		err: fmt.Errorf("auth: 401 status unauthorized: %s", message),
	}
}

func (e ErrorStatusUnauthorized) Error() string {
	return e.err.Error()
}

// ErrorNon200Status wraps an error returned when a non-200 status is received in the UpdateUserToken method.
type ErrorNon200Status struct {
	err error
}

func newErrorNon200Status(statusCode int) ErrorNon200Status {
	return ErrorNon200Status{
		err: fmt.Errorf("auth: non-200 status: received %d", statusCode),
	}
}

func (e ErrorNon200Status) Error() string {
	return e.err.Error()
}
