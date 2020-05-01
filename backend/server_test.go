package backend

import "testing"

func TestNewServer(t *testing.T) {
	s := NewServer()

	if s == nil {
		t.Error("nil received creating new server")
	}
}

func Test_middleware(t *testing.T) {
	h := middleware()
}
