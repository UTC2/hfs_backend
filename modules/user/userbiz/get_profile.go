package userbiz

import (
  "github.com/gin-gonic/gin"
  "hfs_backend/common"
  "hfs_backend/component"
  "net/http"
)

func GetProfile(appCtx component.AppContext) func(*gin.Context) {
  return func(c *gin.Context) {
    data := c.MustGet(common.CurrentUser).(common.Requester)

    c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
  }
}
