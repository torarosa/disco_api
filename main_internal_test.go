package main

import (
	"testing"
)

func TestFact(t *testing.T) {
	if Fact(1) != 1 {
		t.Error("expected 1")
	}
	if Fact(7) != 5040 {
		t.Error("expected 5040")
	}
}
