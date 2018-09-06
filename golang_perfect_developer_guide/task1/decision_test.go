package main

import "testing"

func TestNewNumbers(t *testing.T) {
	n := newNumbers()

	if len(n) != 11 {
		t.Errorf("Expected numbers length of 11, but got %v", len(n))
	}

	if n[0] != 0 {
		t.Errorf("Expected first numbers of 0, but got %v", n[0])
	}
}
