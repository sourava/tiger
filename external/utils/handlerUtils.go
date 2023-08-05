package utils

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/sourava/tiger/external/customErrors"
	"net/http"
)

func ReturnError(context *gin.Context, err *customErrors.CustomError) {
	log.Error(err.Error())
	context.JSON(err.GetHTTPStatus(), gin.H{
		"success": false,
		"error":   err.Error(),
	})
}

func ReturnSuccessResponse(context *gin.Context, data interface{}) {
	log.Info(data)
	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"payload": data,
	})
}
