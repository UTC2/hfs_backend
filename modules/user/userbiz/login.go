package userbiz

import (
	"context"
	"strings"

	"hfs_backend/common"
	"hfs_backend/component/tokenprovider"
	"hfs_backend/modules/user/usermodel"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	UpdateUser(ctx context.Context, conditions map[string]interface{}, updates map[string]interface{}) error
}

type loginBusiness struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	legacyHasher  LegacyMD5Hasher
	expiry        int
	refreshExpiry int
}

func NewLoginBusiness(
	storeUser LoginStorage,
	tokenProvider tokenprovider.Provider,
	hasher Hasher,
	legacyHasher LegacyMD5Hasher,
	accessExpirySec int,
	refreshExpirySec int,
) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		legacyHasher:  legacyHasher,
		expiry:        accessExpirySec,
		refreshExpiry: refreshExpirySec,
	}
}

// isBcryptHash reports whether s looks like a bcrypt hash.
func isBcryptHash(s string) bool {
	return strings.HasPrefix(s, "$2a$") || strings.HasPrefix(s, "$2b$") || strings.HasPrefix(s, "$2y$")
}

func (b *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	user, err := b.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	var ok bool
	if isBcryptHash(user.Password) {
		ok = b.hasher.Verify(data.Password, user.Password)
	} else {
		// Legacy MD5+salt path
		ok = b.legacyHasher.Hash(data.Password+user.Salt) == user.Password
		if ok {
			// Transparent migration: re-hash with bcrypt and update DB.
			newHash, err := b.hasher.Hash(data.Password)
			if err == nil {
				_ = b.storeUser.UpdateUser(ctx,
					map[string]interface{}{"id": user.Id},
					map[string]interface{}{"password": newHash, "salt": ""},
				)
			}
		}
	}
	if !ok {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{UserId: user.Id, Role: user.Role}
	access, err := b.tokenProvider.Generate(payload, b.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	refresh, err := b.tokenProvider.Generate(payload, b.refreshExpiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return usermodel.NewAccount(access, refresh), nil
}
