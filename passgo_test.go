package main

import (
	"testing"
)

func StandardArgs(t *testing.T, args []string) {
	if len(args) != 2 {
		t.Fatal("Args size incorrect")
	}
	if args[0] != "insert" {
		t.Fatal("Expecting command insert")
	}
	if args[1] != "site-path" {
		t.Fatal("Expecting site path")
	}
}

func TestNotMultiline(t *testing.T) {
	args := []string{"passgo", "insert", "site-path"}
	args = SubArgs(args)
	StandardArgs(t, args)
}
