package auth0

import "fmt"

// ErrAuth0 provides a wrap around low-level Go errors for the auth0 package
type ErrAuth0 struct {
	wrapped error
}

func errAuth0(err error) *ErrAuth0 {
	return &ErrAuth0{
		wrapped: fmt.Errorf("auth0: %w", err),
	}
}

func (e *ErrAuth0) Error() string {
	return e.wrapped.Error()
}
