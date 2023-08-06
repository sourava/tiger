package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	request2 "github.com/sourava/tiger/business/auth/request"
	"github.com/sourava/tiger/business/tiger/request"
	_ "github.com/sourava/tiger/business/tiger/response"
	"github.com/sourava/tiger/business/tiger/service"
	"github.com/sourava/tiger/external/customErrors"
	"github.com/sourava/tiger/external/utils"
	"net/http"
	"strconv"
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

// ListAllTigers godoc
// @Summary      list all tigers api
// @Description  returns all tigers sorted by last time the tiger was seen.
// @Accept       json
// @Produce      json
// @Param page   query int true "Page"
// @Param pageSize query int true "Page Size"
// @Success      200  {object}  response.ListAllTigersHandlerResponse
// @Failure      400  {object} 	utils.HandlerErrorResponse
// @Failure      500  {object}  utils.HandlerErrorResponse
// @Router       /tigers [get]
func (h *TigerHandler) ListAllTigers(context *gin.Context) {
	offset, pageSize, err := utils.ValidatePaginationQueryParams(context)
	if err != nil {
		utils.ReturnError(context, err)
		return
	}

	tigerList, err := h.tigerService.ListAllTigers(&request.ListAllTigerRequest{
		Offset:   offset,
		PageSize: pageSize,
	})
	if err != nil {
		utils.ReturnError(context, err)
		return
	}

	utils.ReturnSuccessResponse(context, tigerList)
	return
}

func (h *TigerHandler) CreateTigerSighting(context *gin.Context) {
	var createTigerSightingRequest *request.CreateTigerSightingRequest
	err := context.ShouldBindBodyWith(&createTigerSightingRequest, binding.JSON)
	if err != nil {
		utils.ReturnError(context, customErrors.NewWithErr(http.StatusBadRequest, err))
		return
	}

	tigerIDStr := context.Param("tigerID")
	tigerID, err := strconv.Atoi(tigerIDStr)
	if err != nil {
		utils.ReturnError(context, customErrors.NewWithMessage(http.StatusBadRequest, "invalid tigerID"))
		return
	}

	claims, claimExists := context.Get("token-claims")
	if !claimExists {
		utils.ReturnSomethingWentWrong(context)
		return
	}

	createdTigerSighting, createTigerSightingErr := h.tigerService.CreateTigerSighting(uint(tigerID), createTigerSightingRequest, claims.(*request2.JWTClaim))
	if createTigerSightingErr != nil {
		utils.ReturnError(context, createTigerSightingErr)
		return
	}

	utils.ReturnSuccessResponse(context, createdTigerSighting)
	return
}

func (h *TigerHandler) ListAllTigerSightings(context *gin.Context) {
	offset, pageSize, validationErr := utils.ValidatePaginationQueryParams(context)
	if validationErr != nil {
		utils.ReturnError(context, validationErr)
		return
	}

	tigerIDStr := context.Param("tigerID")
	tigerID, err := strconv.Atoi(tigerIDStr)
	if err != nil {
		utils.ReturnError(context, customErrors.NewWithMessage(http.StatusBadRequest, "invalid tigerID"))
		return
	}

	tigerSightingList, tigerSightingListErr := h.tigerService.ListAllSightingsForATiger(&request.ListAllTigerSightingsRequest{
		TigerID:  tigerID,
		Offset:   offset,
		PageSize: pageSize,
	})
	if tigerSightingListErr != nil {
		utils.ReturnError(context, tigerSightingListErr)
		return
	}

	utils.ReturnSuccessResponse(context, tigerSightingList)
	return
}
