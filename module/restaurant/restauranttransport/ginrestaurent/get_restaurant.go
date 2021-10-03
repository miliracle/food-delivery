package ginrestaurent

import (
	"fooddelivery/common"
	"fooddelivery/component/appctx"
	"fooddelivery/module/restaurant/restaurantbiz"
	"fooddelivery/module/restaurant/restaurantstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		result, err := biz.GetRestaurant(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
