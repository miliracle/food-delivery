package ginimage

import (
	"fooddelivery/common"
	"fooddelivery/component"
	"fooddelivery/module/image/imagebiz"
	"fooddelivery/module/image/imagestorage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteImage(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := imagestorage.NewSQLStore(appCtx.GetMainDBConnection())
		uploadProvider := appCtx.UploadProvider()
		biz := imagebiz.NewDeleteImageBiz(uploadProvider, store)

		if err := biz.DeleteImage(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
