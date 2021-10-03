package ginuser

import (
	"fooddelivery/common"
	"fooddelivery/component/appctx"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(appctx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
