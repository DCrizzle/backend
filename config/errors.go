package config

import "fmt"

// ErrConfig provides a wrap around low-level Go errors for the config package
type ErrConfig struct {
	wrapped error
}

func errConfig(err error) *ErrConfig {
	return &ErrConfig{
		wrapped: fmt.Errorf("config: %w", err),
	}
}

func (e *ErrConfig) Error() string {
	return e.wrapped.Error()
}
