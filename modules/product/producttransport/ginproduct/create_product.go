package ginproduct

import (
	"github.com/gin-gonic/gin"
	"hfs_backend/common"
	"hfs_backend/component"
	"hfs_backend/modules/product/productbiz"
	"hfs_backend/modules/product/productmodel"
	"hfs_backend/modules/product/productstorage"
	"net/http"
)

func CreateProduct(appCtx component.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {
		var data productmodel.ProductCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := productstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productbiz.NewCreateProductBiz(store)

		if err := biz.CreateProduct(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
