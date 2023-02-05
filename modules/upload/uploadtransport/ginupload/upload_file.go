package ginupload

import (
	"github.com/gin-gonic/gin"
	"simple-rest-api/common"
	"simple-rest-api/component/component"
)

func Upload(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
