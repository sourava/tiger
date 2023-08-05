package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sourava/tiger/business/user/request"
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

func (h *UserHandler) CreateUser(context *gin.Context) {
	var createUserRequest *request.CreateUserRequest
	err := context.ShouldBindBodyWith(&createUserRequest, binding.JSON)
	if err != nil {
		utils.ReturnError(context, customErrors.NewWithErr(http.StatusBadRequest, err))
		return
	}

	createUserErr := h.userService.CreateUser(createUserRequest)
	if err != nil {
		utils.ReturnError(context, createUserErr)
		return
	}

	utils.ReturnSuccessResponse(context, nil)
	return
}
