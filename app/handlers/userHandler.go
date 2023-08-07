package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sourava/tiger/business/user/request"
	_ "github.com/sourava/tiger/business/user/response"
	"github.com/sourava/tiger/external/customErrors"
	"github.com/sourava/tiger/external/utils"

	"github.com/sourava/tiger/business/user/service"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser godoc
// @Summary      create user api
// @Description  creates a user.
// @Accept       json
// @Produce      json
// @Param      	 Authorization  header string					 true "token received in login api response"
// @Param      	 request 		body   request.CreateUserRequest true "create user request body params"
// @Success      200  {object}  response.CreateUserHandlerResponse
// @Failure      400  {object} 	utils.HandlerErrorResponse
// @Failure      500  {object}  utils.HandlerErrorResponse
// @Router       /users [post]
func (h *UserHandler) CreateUser(context *gin.Context) {
	var createUserRequest *request.CreateUserRequest
	err := context.ShouldBindBodyWith(&createUserRequest, binding.JSON)
	if err != nil {
		utils.ReturnError(context, customErrors.NewWithErr(http.StatusBadRequest, err))
		return
	}

	createUserResponse, createUserErr := h.userService.CreateUser(createUserRequest)
	if createUserErr != nil {
		utils.ReturnError(context, createUserErr)
		return
	}

	utils.ReturnSuccessResponse(context, createUserResponse)
	return
}
