package ripemd

import (
	"encoding/hex"
	"testing"
)

func TestRipemd160(t *testing.T) {
	// return DoRipemd160("") == "9c1185a5c5e9fc54612808977ee8f548b2258d31"
	// RIPEMD-160("The quick brown fox jumps over the lazy dog") = 37f332f68db77bd9d7edd4969571ad671cf9dd3b
	r := New160()
	result := r.ComputeBytes([]byte(""))
	t.Log(hex.EncodeToString(result))
}
