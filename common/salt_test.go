package common

import "testing"

func TestGenSalt_Length(t *testing.T) {
	s := GenSalt(50)
	if len(s) != 50 {
		t.Errorf("got len %d, want 50", len(s))
	}
}

func TestGenSalt_NotPredictable(t *testing.T) {
	a := GenSalt(50)
	b := GenSalt(50)
	if a == b {
		t.Error("two consecutive salts collided — RNG is not seeded properly")
	}
}
