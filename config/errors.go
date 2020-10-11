package config

import "fmt"

// ErrorReadFile wraps errors returned by ioutil.ReadFile in the New function.
type ErrorReadFile struct {
	err error
}

func newErrorReadFile(err error) ErrorReadFile {
	return ErrorReadFile{
		err: fmt.Errorf("config: read file: %w", err),
	}
}

func (e ErrorReadFile) Error() string {
	return e.err.Error()
}

// ErrorUnmarshalConfig wraps errors returned by json.Unmarshal in the New function.
type ErrorUnmarshalConfig struct {
	err error
}

func newErrorUnmarshalConfig(err error) ErrorUnmarshalConfig {
	return ErrorUnmarshalConfig{
		err: fmt.Errorf("config: unmarshal config: %w", err),
	}
}

func (e ErrorUnmarshalConfig) Error() string {
	return e.err.Error()
}
