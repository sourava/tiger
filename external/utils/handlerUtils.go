package utils

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/sourava/tiger/external/customErrors"
	"net/http"
	"strconv"
)

func ReturnSomethingWentWrong(context *gin.Context) {
	ReturnError(context, customErrors.NewWithMessage(http.StatusInternalServerError, "Something Went Wrong"))
}

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

func ValidatePaginationQueryParams(context *gin.Context) (int, int, *customErrors.CustomError) {
	pageStr, pageExists := context.GetQuery("page")
	if !pageExists {
		return -1, -1, customErrors.NewWithMessage(http.StatusBadRequest, "page not found in query params")
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return -1, -1, customErrors.NewWithErr(http.StatusBadRequest, err)
	}

	pageSizeStr, pageSizeExists := context.GetQuery("pageSize")
	if !pageSizeExists {
		return -1, -1, customErrors.NewWithMessage(http.StatusBadRequest, "pageSize not found in query params")
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		return -1, -1, customErrors.NewWithErr(http.StatusBadRequest, err)
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	return offset, pageSize, nil
}
