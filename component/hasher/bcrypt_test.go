package hasher

import "testing"

func TestBcrypt_RoundTrip(t *testing.T) {
	h := NewBcryptHash()
	hashed, err := h.Hash("hunter2")
	if err != nil {
		t.Fatalf("Hash: %v", err)
	}
	if !h.Verify("hunter2", hashed) {
		t.Error("Verify rejected correct password")
	}
	if h.Verify("wrong", hashed) {
		t.Error("Verify accepted wrong password")
	}
}

func TestBcrypt_HashesAreSalted(t *testing.T) {
	h := NewBcryptHash()
	a, _ := h.Hash("same")
	b, _ := h.Hash("same")
	if a == b {
		t.Error("expected different hashes for same input (salt)")
	}
}
