package ginhouse

import (
  "github.com/gin-gonic/gin"
  "hfs_backend/common"
  "hfs_backend/component"
  "hfs_backend/modules/house/housebiz"
  "hfs_backend/modules/house/housemodel"
  "hfs_backend/modules/house/housestorage"
  "net/http"
)

func ListHouse(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var filter housemodel.HouseFilter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}

		paging.Fulfill()

		store := housestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz   := housebiz.NewListHouseBiz(store)

		result, err := biz.ListHouse(c.Request.Context(), &filter, &paging)

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
