package main

import (
	ginhelpers "github.com/Cabista/X/GinHelpers"
	"github.com/gin-gonic/gin"
)

func RegisterTrackingApiController(group *gin.RouterGroup) {
	group.POST("/", createTracking)
	group.GET("/:DriverID", getTracking)
}

func createTracking(c *gin.Context) {
	ginCtx := ginhelpers.NewGinContext(c)
	var tracking Tracking
	err := c.BindJSON(&tracking)

	if err != nil {
		ginCtx.AbortFailureErr(500, err)
		return
	}

	db.Save(&tracking)
	ginCtx.Created("Resource created", tracking.ID)
}

func getTracking(c *gin.Context) {
	ginCtx := ginhelpers.NewGinContext(c)

	id := c.Param("DriverID")
	if id == "" {
		ginCtx.AbortFailure(400, "id param was not provided as such this request was aborted")
		return
	}

	var trackings []Tracking
	err := db.Where("DriverID = ?", id).Find(&trackings).Error
	if err != nil {
		ginCtx.AbortFailureErr(404, err)
		return
	}
	// if vehicle == nil {
	// 	fmt.Println("vehicle")
	// 	ginCtx.AbortFailure(404, "a resource by the provided id could not be found")
	// }

	c.JSON(200, trackings)

}
