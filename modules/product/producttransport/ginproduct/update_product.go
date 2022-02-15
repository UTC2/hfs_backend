package ginproduct

import (
	"github.com/gin-gonic/gin"
	"hfs_backend/common"
	"hfs_backend/component"
	"hfs_backend/modules/product/productbiz"
	"hfs_backend/modules/product/productmodel"
	"hfs_backend/modules/product/productstorage"
	"net/http"
	"strconv"
)

func UpdateProduct(appCtx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, error := strconv.Atoi(context.Param("id"))
		if error != nil {
			context.JSON(http.StatusHTTPVersionNotSupported, gin.H{
				"error": error.Error(),
			})
			return
		}
		var data productmodel.ProductUpdate
		if error := context.ShouldBind(&data); error != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": error.Error(),
			})
		}
		store := productstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productbiz.NewUpdateProductBiz(store)

		if error := biz.UpdateProduct(context.Request.Context(), id, &data); error != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": error.Error(),
			})
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
