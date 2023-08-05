package service

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sourava/tiger/business/auth/request"
	"github.com/sourava/tiger/business/auth/response"
	"github.com/sourava/tiger/business/user/models"
	"github.com/sourava/tiger/external/customErrors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"time"
)

var (
	ErrEmailPasswordMismatch = customErrors.NewWithMessage(http.StatusBadRequest, "error email and password combination mismatch")
	ErrGeneratingToken       = customErrors.NewWithMessage(http.StatusBadRequest, "error generating token")
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		DB: db,
	}
}

func (service *AuthService) Login(request *request.LoginRequest) (*response.LoginResponse, *customErrors.CustomError) {
	var user models.User
	record := service.DB.Where("email = ?", request.Email).First(&user)
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

	tokenString, err := generateToken(user.Email, user.Username)
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

func generateToken(email string, username string) (string, error) {
	var jwtKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &request.JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
