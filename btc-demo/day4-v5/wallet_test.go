package main

import (
	"testing"
)

func TestWallet_NewAddress(t *testing.T) {
	w := NewWallet()
	address := w.NewAddress()
	t.Logf("%s", address)
}
