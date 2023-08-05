package service

import (
	"github.com/sourava/tiger/business/auth/request"
	"github.com/sourava/tiger/business/user/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"testing"
)

func setupTests(t *testing.T) (*gorm.DB, func(t *testing.T)) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	t.Setenv("JWT_PRIVATE_KEY", "private_key")

	db.AutoMigrate(&models.User{})
	db.Create(&models.User{Username: "user1", Email: "user1@email.com", Password: "$2a$04$npZR8DN1y2I0VNRrrPG6XOk.C2lfQLzCOhK5T9lR40oQuecSEHkhm"})

	return db, func(t *testing.T) {
		err := os.Remove("test.db")
		if err != nil {
			t.Log(err)
		}
	}
}

func Test_WhenLoginRequestContainsEmailThatIsNotRegistered_ThenReturnErrEmailPasswordMismatch(t *testing.T) {
	gormDB, teardownTestCase := setupTests(t)
	defer teardownTestCase(t)

	authService := NewAuthService(gormDB)
	loginRequest := &request.LoginRequest{
		Email:    "user2@email.com",
		Password: "password",
	}
	_, err := authService.Login(loginRequest)

	assert.NotNil(t, err)
	assert.Equal(t, "error email and password combination mismatch", err.Error())
}

func Test_WhenLoginRequestContainsInvalidPasswordForAUser_ThenReturnErrEmailPasswordMismatch(t *testing.T) {
	gormDB, teardownTestCase := setupTests(t)
	defer teardownTestCase(t)

	authService := NewAuthService(gormDB)
	loginRequest := &request.LoginRequest{
		Email:    "user1@email.com",
		Password: "pass",
	}
	_, err := authService.Login(loginRequest)

	assert.NotNil(t, err)
	assert.Equal(t, "error email and password combination mismatch", err.Error())
}

func Test_WhenLoginRequestIsValid_ThenReturnLoginResponse(t *testing.T) {
	gormDB, teardownTestCase := setupTests(t)
	defer teardownTestCase(t)

	authService := NewAuthService(gormDB)
	loginRequest := &request.LoginRequest{
		Email:    "user1@email.com",
		Password: "password",
	}
	resp, err := authService.Login(loginRequest)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
}
