package setup

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"test-time-tracker/controller"
)

func RouterUser(r *gin.Engine) {
	r.GET("/user/:name", func(context *gin.Context) {
		defer deferredErrorHandler(context)
		context.JSON(http.StatusOK, controller.ThisMonth(context.Params.ByName("name")))
	})
	r.POST("/user/:name", func(context *gin.Context) {
		defer deferredErrorHandler(context)
		controller.AddUser(context.Params.ByName("name"))
		context.String(http.StatusOK, "OK")
	})
	r.PUT("/user/:name/:state", func(context *gin.Context) {
		defer deferredErrorHandler(context)
		controller.SetUserState(context.Params.ByName("name"), context.Params.ByName("state"))
		context.String(http.StatusOK, "OK")
	})
	r.DELETE("/user/:name", func(context *gin.Context) {
		defer deferredErrorHandler(context)
		controller.DeleteUser(context.Params.ByName("name"))
		context.String(http.StatusOK, "OK")
	})
	r.GET("/user", func(context *gin.Context) {
		defer deferredErrorHandler(context)
		context.JSON(http.StatusOK, controller.ListUsers())
	})
}
