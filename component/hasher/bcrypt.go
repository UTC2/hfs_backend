package hasher

import "golang.org/x/crypto/bcrypt"

const bcryptCost = 12

type bcryptHash struct{}

func NewBcryptHash() *bcryptHash {
	return &bcryptHash{}
}

func (h *bcryptHash) Hash(plain string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(plain), bcryptCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (h *bcryptHash) Verify(plain, hashed string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain)) == nil
}
