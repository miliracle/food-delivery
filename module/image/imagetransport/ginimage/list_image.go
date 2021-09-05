package ginimage

import (
	"fooddelivery/common"
	"fooddelivery/component"
	"fooddelivery/module/image/imagebiz"
	"fooddelivery/module/image/imagestorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListImage(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ids []int

		if err := c.ShouldBind(&ids); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := imagestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := imagebiz.NewListImageBiz(store)
		result, err := biz.ListImage(c.Request.Context(), &ids)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
