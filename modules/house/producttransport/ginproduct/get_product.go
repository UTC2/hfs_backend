package ginproduct

import (
	"github.com/gin-gonic/gin"
	"hfs_backend/common"
	"hfs_backend/component"
	"hfs_backend/modules/product/productbiz"
	"hfs_backend/modules/product/productstorage"
	"net/http"
)

func GetProduct(appCtx component.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {
    uid, err := common.FromBase58(c.Param("id"))
    if err != nil {
      panic(common.ErrInvalidRequest(err))
    }
		store := productstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productbiz.NewGetProductBiz(store)

		data, err := biz.GetProduct(c.Request.Context(), int(uid.GetlocalID()))

		if err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
