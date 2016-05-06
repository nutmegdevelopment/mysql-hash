package main

import (
	"testing"
)

func TestHash(t *testing.T) {

	output := hash("right")
	if output != "*920018161824B14A1067A69626595E68CB8284CB" {
		t.Fatal("Incorrect hash:", output)
	}

}
