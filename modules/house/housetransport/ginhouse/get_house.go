package ginhouse

import (
	"github.com/gin-gonic/gin"
	"hfs_backend/common"
	"hfs_backend/component"
  "hfs_backend/modules/house/housebiz"
  "hfs_backend/modules/house/housestorage"
	"net/http"
)

func GetHouse(appCtx component.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {
    uid, err := common.FromBase58(c.Param("id"))
    if err != nil {
      panic(common.ErrInvalidRequest(err))
    }
		store := housestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz   := housebiz.NewGetHouseBiz(store)

		data, err := biz.GetHouse(c.Request.Context(), int(uid.GetlocalID()))

		if err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
