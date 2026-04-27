package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"hfs_backend/common"
	"hfs_backend/component"
)

func Recover(_ component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				handlePanic(c, r)
			}
		}()
		c.Next()
	}
}

func handlePanic(c *gin.Context, r interface{}) {
	c.Header("content-type", "application/json")
	var appErr *common.AppError
	switch e := r.(type) {
	case *common.AppError:
		appErr = e
	case error:
		appErr = common.ErrInternal(e)
	default:
		appErr = common.ErrInternal(fmt.Errorf("%v", r))
	}
	c.AbortWithStatusJSON(appErr.StatusCode, appErr)
}
