package service

import (
	"errors"
	"github.com/sourava/tiger/business/user/models"
	"github.com/sourava/tiger/business/user/request"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/mail"
)

var (
	ErrCreateUserRequestContainsEmptyParams error = errors.New("error create user request contains empty username, email or password")
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB: db,
	}
}

func (service *UserService) CreateUser(request *request.CreateUserRequest) error {
	err := validateCreateUserRequest(request)
	if err != nil {
		return err
	}

	hashedPassword, err := hashPassword(request.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: hashedPassword,
	}
	result := service.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func hashPassword(password string) (string, error) {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Hash password with bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

func validateCreateUserRequest(userRequest *request.CreateUserRequest) error {
	if userRequest.Username == "" || userRequest.Email == "" || userRequest.Password == "" {
		return ErrCreateUserRequestContainsEmptyParams
	}

	_, err := mail.ParseAddress(userRequest.Email)
	if err != nil {
		return err
	}

	return nil
}
