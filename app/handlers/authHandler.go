package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sourava/tiger/business/auth/request"
	_ "github.com/sourava/tiger/business/auth/response"
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

// Login godoc
// @Summary      login api
// @Description  provides token if email password combination is correct
// @Accept       json
// @Produce      json
// @Param      	 request body request.LoginRequest true "login request body params"
// @Success      200  {object}  response.LoginHandlerResponse
// @Failure      400  {object} 	utils.HandlerErrorResponse
// @Failure      500  {object}  utils.HandlerErrorResponse
// @Router       /login [post]
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
