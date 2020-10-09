package auth

import "fmt"

// ErrAuth provides a wrap around low-level Go errors for the auth package.
type ErrAuth struct {
	wrapped error
}

func errAuth(err error) *ErrAuth {
	return &ErrAuth{
		wrapped: fmt.Errorf("auth: %w", err),
	}
}

func (e *ErrAuth) Error() string {
	return e.wrapped.Error()
}
