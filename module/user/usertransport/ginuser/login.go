package ginuser

import (
	"fooddelivery/common"
	"fooddelivery/component/appctx"
	"fooddelivery/component/hasher"
	"fooddelivery/component/tokenprovider/jwt"
	"fooddelivery/module/user/userbiz"
	"fooddelivery/module/user/usermodel"
	"fooddelivery/module/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		tokenprovider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewLoginBiz(store, tokenprovider, md5)
		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))

	}
}
