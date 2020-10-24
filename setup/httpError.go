package setup

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"test-time-tracker/data"
	"test-time-tracker/models"
	"test-time-tracker/services"
)

func deferredErrorHandler(context *gin.Context) {
	if err := recover(); err != nil {
		switch t := err.(type) {
		case error:
			if errors.Is(t, data.ErrorDataLayer) {
				context.JSON(http.StatusNotFound, t)
			} else if errors.Is(t, services.ErrorTimeTrack) {
				context.JSON(http.StatusForbidden, t)
			} else if errors.Is(t, models.ErrorUnknown) {
				context.JSON(http.StatusNotAcceptable, t)
			} else if errors.Is(t, services.ErrorUser) {
				context.JSON(http.StatusConflict, t)
			} else {
				context.JSON(http.StatusBadRequest, t)
			}
		default:
			context.JSON(http.StatusInternalServerError, t)
		}

	}
}
