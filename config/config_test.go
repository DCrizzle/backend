package config

import "testing"

func TestReadConfig(t *testing.T) {
	config, err := New("../etc/config/config.json")
	if err != nil {
		t.Fatalf("error reading config: %s, config: %+v\n", err.Error(), config)
	}
}
