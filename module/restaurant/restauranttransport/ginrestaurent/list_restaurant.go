package ginrestaurent

import (
	"fooddelivery/common"
	"fooddelivery/component"
	"fooddelivery/module/restaurant/restaurantbiz"
	"fooddelivery/module/restaurant/restaurantmodel"
	"fooddelivery/module/restaurant/restaurantstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		}

		paging.Fulfill()

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store)
		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, &paging, &filter))
	}
}