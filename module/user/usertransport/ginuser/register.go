package ginuser

import (
	"fooddelivery/common"
	"fooddelivery/component"
	"fooddelivery/module/user/userbiz"
	"fooddelivery/module/user/usermodel"
	"fooddelivery/module/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(appctx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {

		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := appctx.GetMainDBConnection()

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewRegisterBiz(store, md5)

		if err := biz.RegisterUser(c, &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
