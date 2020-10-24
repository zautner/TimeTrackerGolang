package setup

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"test-time-tracker/controller"
)

func RouterTracker(r *gin.Engine) {
	// Track user activity
	r.POST("/track/:name", func(context *gin.Context) {
		defer deferredErrorHandler(context)
		controller.Track(context.Params.ByName("name"))
		context.String(http.StatusOK, "OK")
	})
	r.PUT("/track/:name", func(context *gin.Context) {
		defer deferredErrorHandler(context)
		controller.ResetForUser(context.Params.ByName("name"))
		context.String(http.StatusOK, "OK")
	})

	r.GET("/track/days/:name/:back", func(context *gin.Context) {
		defer deferredErrorHandler(context)
		iface := controller.GetDaysBack(context.Params.ByName("name"), context.Params.ByName("back"))
		context.JSON(http.StatusOK, iface)
	})

	r.GET("/track/months/:name/:back", func(context *gin.Context) {
		defer deferredErrorHandler(context)
		iface := controller.GetMonthsBack(context.Params.ByName("name"), context.Params.ByName("back"))
		context.JSON(http.StatusOK, iface)
	})
}
