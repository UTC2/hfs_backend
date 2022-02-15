package ginproduct

import (
	"github.com/gin-gonic/gin"
	"hfs_backend/common"
	"hfs_backend/component"
	"hfs_backend/modules/product/productbiz"
	"hfs_backend/modules/product/productstorage"
	"net/http"
	"strconv"
)

func GetProduct(appCtx component.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := productstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productbiz.NewGetProductBiz(store)

		data, err := biz.GetProduct(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
