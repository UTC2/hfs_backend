package ginhouse

import (
  "github.com/gin-gonic/gin"
  "hfs_backend/common"
  "hfs_backend/component"
  "hfs_backend/modules/house/housebiz"
  "hfs_backend/modules/house/housemodel"
  "hfs_backend/modules/house/housestorage"
  "net/http"
  "strconv"
)

func UpdateHouse(appCtx component.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		id, error := strconv.Atoi(context.Param("id"))
		if error != nil {
			context.JSON(http.StatusHTTPVersionNotSupported, gin.H{
				"error": error.Error(),
			})
			return
		}
		var data housemodel.HouseUpdate
		if error := context.ShouldBind(&data); error != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": error.Error(),
			})
		}
		store := housestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz   := housebiz.NewUpdateHouseBiz(store)

		if error := biz.UpdateHouse(context.Request.Context(), id, &data); error != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": error.Error(),
			})
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
