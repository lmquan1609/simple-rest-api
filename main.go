package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"simple-rest-api/component/component"
	"simple-rest-api/modules/restaurant/restaurantmodel"
	"simple-rest-api/modules/restaurant/restauranttransport/ginrestaurant"
	"strconv"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//CRUD
	appCtx := component.NewAppContext(db)
	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))

		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))

		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))

		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))

		restaurants.DELETE("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}

			if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
				Where("id = ?", id).
				Delete(nil).Error; err != nil {
				c.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, map[string]interface{}{
				"ok": 1,
			})
		})
	}
	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
