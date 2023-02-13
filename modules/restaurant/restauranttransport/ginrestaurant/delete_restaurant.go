package ginrestaurant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-rest-api/common"
	"simple-rest-api/component/component"
	"simple-rest-api/modules/restaurant/restaurantbiz"
	"simple-rest-api/modules/restaurant/restaurantstorage"
	"strconv"
)

func DeleteRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			common.ErrInvalidRequest(err)
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		err = biz.DeleteRestaurant(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
