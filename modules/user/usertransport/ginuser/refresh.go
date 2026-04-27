package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hfs_backend/common"
	"hfs_backend/component"
	"hfs_backend/component/tokenprovider/jwt"
	"hfs_backend/modules/user/userbiz"
	"hfs_backend/modules/user/usermodel"
)

func Refresh(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req usermodel.RefreshRequest
		if err := c.ShouldBind(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		tp := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		biz := userbiz.NewRefreshBusiness(tp, 60*15, 60*60*24*30)
		acc, err := biz.Refresh(c.Request.Context(), req.RefreshToken)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(acc))
	}
}
