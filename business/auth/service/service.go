package service

import (
	"github.com/sourava/tiger/app/middlewares"
	"github.com/sourava/tiger/business/auth/request"
	"github.com/sourava/tiger/business/auth/response"
	"github.com/sourava/tiger/business/constants"
	"github.com/sourava/tiger/business/user/models"
	"github.com/sourava/tiger/external/customErrors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"net/mail"
)

type AuthService struct {
	db        *gorm.DB
	jwtSecret string
}

func NewAuthService(db *gorm.DB, jwtSecret string) *AuthService {
	return &AuthService{
		db:        db,
		jwtSecret: jwtSecret,
	}
}

func (service *AuthService) Login(request *request.LoginRequest) (*response.LoginHandlerResponse, *customErrors.CustomError) {
	_, err := mail.ParseAddress(request.Email)
	if err != nil {
		return nil, constants.ErrInvalidEmail
	}

	var user models.User
	record := service.db.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		if record.Error.Error() == "record not found" {
			return nil, constants.ErrEmailPasswordMismatch
		}
		return nil, customErrors.NewWithErr(http.StatusInternalServerError, record.Error)
	}

	err = comparePassword(request.Password, user.Password)
	if err != nil {
		return nil, constants.ErrEmailPasswordMismatch
	}

	tokenString, err := middlewares.GenerateToken(user.ID, user.Email, user.Username, service.jwtSecret)
	if err != nil {
		return nil, constants.ErrGeneratingToken
	}

	return &response.LoginHandlerResponse{
		Success: true,
		Payload: response.LoginResponse{
			Token: tokenString,
		},
	}, nil
}

func comparePassword(providedPassword string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
