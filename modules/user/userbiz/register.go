package userbiz

import (
	"context"
	"hfs_backend/common"
	"hfs_backend/modules/user/usermodel"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	// Hash computes a bcrypt hash of the plaintext password.
	Hash(plain string) (string, error)
	// Verify checks plaintext against either a bcrypt hash (preferred)
	// or, for migration, a salted MD5 (caller computes the salted form).
	Verify(plain, hashed string) bool
}

// LegacyMD5Hasher is the old salted-MD5 path, retained only as a verifier
// for users who haven't logged in since bcrypt was introduced.
type LegacyMD5Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		registerStorage: registerStorage,
		hasher:          hasher,
	}
}

func (business *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := business.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if user != nil {
		return usermodel.ErrEmailExisted
	}

	hashed, err := business.hasher.Hash(data.Password)
	if err != nil {
		return common.ErrInternal(err)
	}
	data.Password = hashed
	data.Salt = "" // bcrypt embeds salt
	data.Role = "user"
	data.Status = 1

	if err := business.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
