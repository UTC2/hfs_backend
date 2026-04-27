package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"hfs_backend/component/tokenprovider"
)

type jwtProvider struct {
	secret string
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.RegisteredClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		Payload: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Duration(expiry) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	})
	signed, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, tokenprovider.ErrEncodingToken
	}
	return &tokenprovider.Token{
		Token:   signed,
		Expiry:  expiry,
		Created: now,
	}, nil
}

func (j *jwtProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	claims := &myClaims{}
	parsed, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secret), nil
	})
	if err != nil || !parsed.Valid {
		return nil, tokenprovider.ErrInvalidToken
	}
	return &claims.Payload, nil
}
