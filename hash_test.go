package secret

import (
	"crypto"
	"testing"
)

var (
	h = NewHasher()
)

func TestMyHasher_HashToBytes(t *testing.T) {
	result, err := h.HashToBytes(dataStr, crypto.MD5)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	t.Logf("%s\n", result)
}

func TestMyHasher_HashToString(t *testing.T) {
	result, err := h.HashToString(dataStr, crypto.SHA256)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	t.Logf("%s\n", result)
}

func TestMyHasher_DoubleHashToBytes(t *testing.T) {
	result, err := h.DoubleHashToBytes(dataStruct, crypto.SHA384)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	t.Logf("%s\n", result)
}

func TestMyHasher_DoubleHashToString(t *testing.T) {
	result, err := h.DoubleHashToString(dataStr, crypto.RIPEMD160)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	t.Logf("%s\n", result)
}