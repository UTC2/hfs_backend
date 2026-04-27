package userbiz

import (
	"context"
	"testing"

	"hfs_backend/component/tokenprovider"
)

type stubProvider struct {
	validatePayload *tokenprovider.TokenPayload
	validateErr     error
}

func (s stubProvider) Generate(p tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	return &tokenprovider.Token{Token: "new", Expiry: expiry}, nil
}
func (s stubProvider) Validate(_ string) (*tokenprovider.TokenPayload, error) {
	return s.validatePayload, s.validateErr
}

func TestRefresh_RotatesPair(t *testing.T) {
	p := stubProvider{validatePayload: &tokenprovider.TokenPayload{UserId: 9, Role: "user"}}
	biz := NewRefreshBusiness(p, 60, 600)
	acc, err := biz.Refresh(context.Background(), "old-refresh-token")
	if err != nil {
		t.Fatalf("refresh: %v", err)
	}
	if acc.AccessToken.Token != "new" || acc.RefreshToken.Token != "new" {
		t.Errorf("expected both tokens reissued, got %+v", acc)
	}
}

func TestRefresh_InvalidToken(t *testing.T) {
	p := stubProvider{validateErr: tokenprovider.ErrInvalidToken}
	biz := NewRefreshBusiness(p, 60, 600)
	_, err := biz.Refresh(context.Background(), "bad")
	if err != tokenprovider.ErrInvalidToken {
		t.Errorf("expected ErrInvalidToken, got %v", err)
	}
}
