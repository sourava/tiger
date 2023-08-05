package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	request2 "github.com/sourava/tiger/business/auth/request"
	"github.com/sourava/tiger/business/tiger/request"
	"github.com/sourava/tiger/business/tiger/service"
	"github.com/sourava/tiger/external/customErrors"
	"github.com/sourava/tiger/external/utils"
	"net/http"
)

type TigerHandler struct {
	tigerService *service.TigerService
}

func NewTigerHandler(tigerService *service.TigerService) *TigerHandler {
	return &TigerHandler{
		tigerService: tigerService,
	}
}

func (h *TigerHandler) CreateTiger(context *gin.Context) {
	var createTigerRequest *request.CreateTigerRequest
	err := context.ShouldBindBodyWith(&createTigerRequest, binding.JSON)
	if err != nil {
		utils.ReturnError(context, customErrors.NewWithErr(http.StatusBadRequest, err))
		return
	}

	claims, claimExists := context.Get("token-claims")
	if !claimExists {
		utils.ReturnSomethingWentWrong(context)
		return
	}

	createdTiger, createTigerErr := h.tigerService.CreateTiger(createTigerRequest, claims.(*request2.JWTClaim))
	if createTigerErr != nil {
		utils.ReturnError(context, createTigerErr)
		return
	}

	utils.ReturnSuccessResponse(context, createdTiger)
	return
}