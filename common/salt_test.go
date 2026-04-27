package common

import "testing"

func TestGenSalt_Length(t *testing.T) {
	s := GenSalt(50)
	if len(s) != 50 {
		t.Errorf("got len %d, want 50", len(s))
	}
}

func TestGenSalt_DistinctPerCall(t *testing.T) {
	a := GenSalt(50)
	b := GenSalt(50)
	if a == b {
		t.Error("two consecutive 50-char salts collided — RNG is broken or output domain shrank")
	}
}

func TestGenSalt_NonPositiveLengthDefaultsTo50(t *testing.T) {
	for _, n := range []int{0, -1, -100} {
		if got := len(GenSalt(n)); got != 50 {
			t.Errorf("GenSalt(%d): got len %d, want 50", n, got)
		}
	}
}
