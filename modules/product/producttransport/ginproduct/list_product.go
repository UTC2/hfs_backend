package ginproduct

import (
	"github.com/gin-gonic/gin"
	"hfs_backend/common"
	"hfs_backend/component"
	"hfs_backend/modules/product/productbiz"
	"hfs_backend/modules/product/productmodel"
	"hfs_backend/modules/product/productstorage"
  productlikestorage "hfs_backend/modules/productLike/storage"
  "net/http"
)

func ListProduct(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var filter productmodel.ProductFilter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}

		paging.Fulfill()

		store     := productstorage.NewSQLStore(appCtx.GetMainDBConnection())
		likeStore := productlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz       := productbiz.NewListProductBiz(store,likeStore)

		result, err := biz.ListProduct(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		for i := range result {
		  result[i].Mask(false)

      if i == len(result)-1 {
        paging.NextCursor = result[i].FakeId.String()
      }

    }
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
