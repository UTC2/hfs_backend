package middleware

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"hfs_backend/common"
	"hfs_backend/component"
	"hfs_backend/component/tokenprovider/jwt"
	"hfs_backend/modules/user/userstorage"
)

func RequiredAuth(appCtx component.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractToken(c.GetHeader("Authorization"))
		if err != nil {
			handlePanic(c, err)
			return
		}

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			handlePanic(c, common.NewUnauthorized(err, "invalid token", "ErrInvalidToken"))
			return
		}

		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)
		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})
		if err != nil {
			handlePanic(c, common.ErrEntityNotFound("User", err))
			return
		}
		if user.Status == 0 {
			handlePanic(c, common.ErrNoPermission(errors.New("user has been deleted or banned")))
			return
		}

		user.Mask(false)
		c.Set(common.CurrentUser, user)
		c.Next()
	}
}

func extractToken(h string) (string, error) {
	parts := strings.Split(h, " ")
	if len(parts) < 2 || parts[0] != "Bearer" || strings.TrimSpace(parts[1]) == "" {
		return "", common.NewUnauthorized(
			errors.New("missing or malformed Authorization header"),
			"missing or malformed Authorization header",
			"ErrWrongAuthHeader",
		)
	}
	return parts[1], nil
}
