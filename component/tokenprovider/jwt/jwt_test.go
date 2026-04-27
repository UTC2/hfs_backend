package jwt

import (
	"hfs_backend/component/tokenprovider"
	"testing"
	"time"
)

const testSecret = "test-secret-please-change-me-in-prod"

func TestGenerateValidate_RoundTrip(t *testing.T) {
	p := NewTokenJWTProvider(testSecret)
	tok, err := p.Generate(tokenprovider.TokenPayload{UserId: 42, Role: "user"}, 60)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	payload, err := p.Validate(tok.Token)
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	if payload.UserId != 42 || payload.Role != "user" {
		t.Errorf("payload mismatch: got %+v", payload)
	}
}

func TestValidate_TamperedToken(t *testing.T) {
	p := NewTokenJWTProvider(testSecret)
	tok, _ := p.Generate(tokenprovider.TokenPayload{UserId: 1, Role: "user"}, 60)
	tampered := tok.Token + "x"
	if _, err := p.Validate(tampered); err == nil {
		t.Fatal("expected error for tampered token")
	}
}

func TestValidate_ExpiredToken(t *testing.T) {
	p := NewTokenJWTProvider(testSecret)
	tok, _ := p.Generate(tokenprovider.TokenPayload{UserId: 1, Role: "user"}, -1)
	time.Sleep(10 * time.Millisecond)
	if _, err := p.Validate(tok.Token); err == nil {
		t.Fatal("expected error for expired token")
	}
}

func TestValidate_WrongSecret(t *testing.T) {
	p1 := NewTokenJWTProvider(testSecret)
	p2 := NewTokenJWTProvider("different-secret")
	tok, _ := p1.Generate(tokenprovider.TokenPayload{UserId: 1, Role: "user"}, 60)
	if _, err := p2.Validate(tok.Token); err == nil {
		t.Fatal("expected error for wrong secret")
	}
}
