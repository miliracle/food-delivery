package ginimage

import (
	"fmt"
	"fooddelivery/common"
	"fooddelivery/component/appctx"
	"fooddelivery/module/image/imagebiz"
	"fooddelivery/module/image/imagestorage"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		multipart, err := c.Request.MultipartReader()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		for {
			part, err := multipart.NextPart()

			if err == io.EOF {
				break
			}

			if err != nil {
				panic(common.ErrInvalidRequest(err))
			}

			if part.FormName() == "file" {
				db := appCtx.GetMainDBConnection()
				uploadProvider := appCtx.UploadProvider()
				imgStore := imagestorage.NewSQLStore(db)
				biz := imagebiz.NewUploadBiz(uploadProvider, imgStore)

				fmt.Print(part.Header)

				img, err := biz.UploadImage(c, part, "image", part.FileName())

				if err != nil {
					panic(err)
				}

				c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
			}
		}
	}
}
