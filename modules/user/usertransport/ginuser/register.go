package ginuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hfs_backend/common"
	"hfs_backend/component"
	"hfs_backend/component/hasher"
	"hfs_backend/modules/user/userbiz"
	"hfs_backend/modules/user/usermodel"
	"hfs_backend/modules/user/userstorage"
)

func Register(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		bcryptHasher := hasher.NewBcryptHash()
		biz := userbiz.NewRegisterBusiness(store, bcryptHasher)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
