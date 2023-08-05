package service

import (
	"github.com/sourava/tiger/app/middlewares"
	"github.com/sourava/tiger/business/auth/request"
	"github.com/sourava/tiger/business/auth/response"
	"github.com/sourava/tiger/business/user/models"
	"github.com/sourava/tiger/external/customErrors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

var (
	ErrEmailPasswordMismatch = customErrors.NewWithMessage(http.StatusBadRequest, "error email and password combination mismatch")
	ErrGeneratingToken       = customErrors.NewWithMessage(http.StatusBadRequest, "error generating token")
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

func (service *AuthService) Login(request *request.LoginRequest) (*response.LoginResponse, *customErrors.CustomError) {
	var user models.User
	record := service.db.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		if record.Error.Error() == "record not found" {
			return nil, ErrEmailPasswordMismatch
		}
		return nil, customErrors.NewWithErr(http.StatusInternalServerError, record.Error)
	}

	err := comparePassword(request.Password, user.Password)
	if err != nil {
		return nil, ErrEmailPasswordMismatch
	}

	tokenString, err := middlewares.GenerateToken(user.Email, user.Username, service.jwtSecret)
	if err != nil {
		return nil, ErrGeneratingToken
	}

	return &response.LoginResponse{
		Token: tokenString,
	}, nil
}

func comparePassword(providedPassword string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
