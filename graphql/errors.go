package graphql

import "fmt"

// ErrorMarshalConfig wraps errors returned by json.Marshal in the SendRequest method.
type ErrorMarshalConfig struct {
	err error
}

func newErrorMarshalConfig(err error) ErrorMarshalConfig {
	return ErrorMarshalConfig{
		err: fmt.Errorf("graphql: marshal config: %w", err),
	}
}

func (e ErrorMarshalConfig) Error() string {
	return e.err.Error()
}

// ErrorNewRequest wraps errors returned by http.NewRequest in the SendRequest method.
type ErrorNewRequest struct {
	err error
}

func newErrorNewRequest(err error) ErrorNewRequest {
	return ErrorNewRequest{
		err: fmt.Errorf("graphql: new request: %w", err),
	}
}

func (e ErrorNewRequest) Error() string {
	return e.err.Error()
}

// ErrorClientDo wraps errors returned by http.Client.Do in the SendRequest method.
type ErrorClientDo struct {
	err error
}

func newErrorClientDo(err error) ErrorClientDo {
	return ErrorClientDo{
		err: fmt.Errorf("graphql: client do: %w", err),
	}
}

func (e ErrorClientDo) Error() string {
	return e.err.Error()
}

// ErrorReadAll wraps errors returned by ioutil.ReadAll in the SendRequest method.
type ErrorReadAll struct {
	err error
}

func newErrorReadAll(err error) ErrorReadAll {
	return ErrorReadAll{
		err: fmt.Errorf("graphql: read all: %w", err),
	}
}

func (e ErrorReadAll) Error() string {
	return e.err.Error()
}
