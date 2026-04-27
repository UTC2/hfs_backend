package userbiz

import (
	"context"
	"testing"

	"hfs_backend/component/hasher"
	"hfs_backend/component/tokenprovider"
	"hfs_backend/modules/user/usermodel"
)

type fakeStore struct {
	user        *usermodel.User
	updateCalls []map[string]interface{}
}

func (f *fakeStore) FindUser(_ context.Context, _ map[string]interface{}, _ ...string) (*usermodel.User, error) {
	if f.user == nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}
	return f.user, nil
}
func (f *fakeStore) UpdateUser(_ context.Context, _ map[string]interface{}, updates map[string]interface{}) error {
	f.updateCalls = append(f.updateCalls, updates)
	return nil
}

type fakeProvider struct{}

func (fakeProvider) Generate(_ tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	return &tokenprovider.Token{Token: "tok", Expiry: expiry}, nil
}
func (fakeProvider) Validate(string) (*tokenprovider.TokenPayload, error) { return nil, nil }

func TestLogin_BcryptUser(t *testing.T) {
	bc := hasher.NewBcryptHash()
	hashed, _ := bc.Hash("hunter2")
	store := &fakeStore{user: &usermodel.User{Email: "a@b.c", Password: hashed, Role: "user"}}
	store.user.Id = 7

	biz := NewLoginBusiness(store, fakeProvider{}, bc, hasher.NewMd5Hash(), 60, 600)
	acc, err := biz.Login(context.Background(), &usermodel.UserLogin{Email: "a@b.c", Password: "hunter2"})
	if err != nil {
		t.Fatalf("login: %v", err)
	}
	if acc.AccessToken.Token != "tok" {
		t.Error("expected token in account")
	}
	if len(store.updateCalls) != 0 {
		t.Error("bcrypt user should not trigger migration update")
	}
}

func TestLogin_LegacyMD5User_Migrates(t *testing.T) {
	md5 := hasher.NewMd5Hash()
	salt := "saltsalt"
	legacyHash := md5.Hash("hunter2" + salt)
	store := &fakeStore{user: &usermodel.User{Email: "a@b.c", Password: legacyHash, Salt: salt, Role: "user"}}
	store.user.Id = 7

	biz := NewLoginBusiness(store, fakeProvider{}, hasher.NewBcryptHash(), md5, 60, 600)
	_, err := biz.Login(context.Background(), &usermodel.UserLogin{Email: "a@b.c", Password: "hunter2"})
	if err != nil {
		t.Fatalf("login: %v", err)
	}
	if len(store.updateCalls) != 1 {
		t.Fatalf("expected 1 update (migration), got %d", len(store.updateCalls))
	}
	if store.updateCalls[0]["salt"] != "" {
		t.Error("expected salt to be cleared")
	}
}

func TestLogin_WrongPassword(t *testing.T) {
	bc := hasher.NewBcryptHash()
	hashed, _ := bc.Hash("hunter2")
	store := &fakeStore{user: &usermodel.User{Email: "a@b.c", Password: hashed, Role: "user"}}
	store.user.Id = 7

	biz := NewLoginBusiness(store, fakeProvider{}, bc, hasher.NewMd5Hash(), 60, 600)
	_, err := biz.Login(context.Background(), &usermodel.UserLogin{Email: "a@b.c", Password: "wrong"})
	if err != usermodel.ErrUsernameOrPasswordInvalid {
		t.Errorf("expected invalid creds error, got %v", err)
	}
}
