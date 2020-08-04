package main

import "testing"

func Test_readConfig(t *testing.T) {
	c, err := readConfig()
	if err != nil {
		t.Fatalf("error reading config: %s, config: %+v\n", err.Error(), c)
	}
}
