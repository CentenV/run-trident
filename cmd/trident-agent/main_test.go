package main

import (
	"testing"
)

func Test(t *testing.T) {
	if 0 == 1 {
		t.Errorf("Testing works")
	}
}
