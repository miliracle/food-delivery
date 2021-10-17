package ginrestaurent

import (
	"fooddelivery/common"
	"fooddelivery/component/appctx"
	"fooddelivery/module/restaurant/restaurantbiz"
	"fooddelivery/module/restaurant/restaurantmodel"
	"fooddelivery/module/restaurant/restaurantstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data.UserId = requester.GetUserId()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.GenUID(common.DBTypeRestaurant)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
