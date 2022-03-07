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
func CreateHouse(appCtx component.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {
		var data housemodel.HouseCreate
		if err := c.ShouldBind(&data); err != nil {
      panic(common.ErrInvalidRequest(err))
		}
    requester := c.MustGet(common.CurrentUser).(common.Requester)
    data.OwnerId = requester.GetUserId()

    store := housestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz   := housebiz.NewCreateHouseBiz(store)
		if err := biz.CreateHouse(c.Request.Context(), &data); err != nil {
      panic(err)
		}
		data.GenerateUID(common.DbTypeRestaurant)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
