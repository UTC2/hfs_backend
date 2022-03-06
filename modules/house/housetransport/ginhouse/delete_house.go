package ginhouse

import (
  "github.com/gin-gonic/gin"
  "hfs_backend/common"
  "hfs_backend/component"
  "hfs_backend/modules/house/housebiz"
  "hfs_backend/modules/house/housestorage"
  "net/http"
  "strconv"
)

func DeleteHouse(appCtx component.AppContext) gin.HandlerFunc {

	return func(context *gin.Context) {
		id, error := strconv.Atoi(context.Param("id"))
		if error != nil {
			context.JSON(http.StatusHTTPVersionNotSupported, gin.H{
				"error": error.Error(),
			})
			return
		}
		store := housestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz   := housebiz.NewDeleteHouseBiz(store)

		if error := biz.DeleteHouse(context.Request.Context(), id); error != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": error.Error(),
			})
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
