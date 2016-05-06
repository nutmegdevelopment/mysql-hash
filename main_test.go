package main

import (
	"testing"
)

func TestHash(t *testing.T) {

	output := Hash("right")
	if output != "*920018161824B14A1067A69626595E68CB8284CB" {
		t.Fatal("Incorrect hash:", output)
	}

}

func TestNewPass(t *testing.T) {

	pass := NewPass()
	if len(pass) != passwordLen {
		t.Fatal("Bad length")
	}

}
