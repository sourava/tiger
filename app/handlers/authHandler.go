package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sourava/tiger/business/auth/request"
	"github.com/sourava/tiger/external/customErrors"
	"github.com/sourava/tiger/external/utils"

	"github.com/sourava/tiger/business/auth/service"
	"net/http"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(context *gin.Context) {
	var loginRequest *request.LoginRequest
	err := context.ShouldBindBodyWith(&loginRequest, binding.JSON)
	if err != nil {
		utils.ReturnError(context, customErrors.NewWithErr(http.StatusBadRequest, err))
		return
	}

	resp, loginErr := h.authService.Login(loginRequest)
	if loginErr != nil {
		utils.ReturnError(context, loginErr)
		return
	}

	utils.ReturnSuccessResponse(context, resp)
	return
}
