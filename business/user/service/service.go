package service

import (
	"github.com/sourava/tiger/business/user/models"
	"github.com/sourava/tiger/business/user/request"
	"github.com/sourava/tiger/business/user/response"
	"github.com/sourava/tiger/external/customErrors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"net/mail"
)

var (
	ErrEmptyParams     = customErrors.NewWithMessage(http.StatusBadRequest, "error request contains empty username, email or password")
	ErrInvalidEmail    = customErrors.NewWithMessage(http.StatusBadRequest, "error request contains invalid email")
	ErrHashingPassword = customErrors.NewWithMessage(http.StatusBadRequest, "error while hashing password")
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

func (service *UserService) CreateUser(request *request.CreateUserRequest) (*response.CreateUserHandlerResponse, *customErrors.CustomError) {
	err := validateCreateUserRequest(request)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := hashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	}
	result := service.DB.Create(&user)
	if result.Error != nil {
		return nil, customErrors.NewWithErr(http.StatusBadRequest, result.Error)
	}

	return &response.CreateUserHandlerResponse{
		Success: true,
		Payload: response.CreateUserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

func hashPassword(password string) (string, *customErrors.CustomError) {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Hash password with bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)
	if err != nil {
		return "", ErrHashingPassword
	}

	return string(hashedPasswordBytes), nil
}

func validateCreateUserRequest(userRequest *request.CreateUserRequest) *customErrors.CustomError {
	if userRequest.Username == "" || userRequest.Email == "" || userRequest.Password == "" {
		return ErrEmptyParams
	}

	_, err := mail.ParseAddress(userRequest.Email)
	if err != nil {
		return ErrInvalidEmail
	}

	return nil
}
