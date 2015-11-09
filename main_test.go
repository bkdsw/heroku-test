package main

import "testing"

func TestAddition(t *testing.T) {
	if 1+1 != 2 {
		t.Fatal("1+1 should equal 2")
	}
}
