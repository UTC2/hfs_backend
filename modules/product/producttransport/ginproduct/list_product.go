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

func ListProduct(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var filter productmodel.ProductFilter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.Fulfill()

		store := productstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productbiz.NewListProductBiz(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
