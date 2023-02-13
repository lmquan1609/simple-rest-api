package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"simple-rest-api/common"
	"simple-rest-api/component/component"
	"simple-rest-api/modules/restaurant/restaurantbiz"
	"simple-rest-api/modules/restaurant/restaurantmodel"
	"simple-rest-api/modules/restaurant/restaurantstorage"
	"strconv"
)

func UpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.JSON(401, common.SimpleSuccessResponse(true))
	}
}
