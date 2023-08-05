package service

import (
	"github.com/sourava/tiger/business/user/models"
	"github.com/sourava/tiger/business/user/request"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"testing"
)

func setupTests() (*gorm.DB, func(t *testing.T)) {
	db, err := gorm.Open(sqlite.Open("test.db?_foreign_keys=on"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})

	return db, func(t *testing.T) {
		err := os.Remove("test.db")
		if err != nil {
			t.Log(err)
		}
	}
}

func Test_WhenCreateUserRequestContainsEmptyParams_ThenReturnErrCreateUserRequestContainsEmptyParams(t *testing.T) {
	gormDB, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	userService := NewUserService(gormDB)
	createUserRequest := &request.CreateUserRequest{
		Username: "",
		Email:    "",
		Password: "",
	}
	err := userService.CreateUser(createUserRequest)

	assert.NotNil(t, err)
	assert.Equal(t, "error request contains empty username, email or password", err.Error())
}

func Test_WhenCreateUserRequestContainsInvalidEmail_ThenReturnError(t *testing.T) {
	gormDB, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	userService := NewUserService(gormDB)
	createUserRequest := &request.CreateUserRequest{
		Username: "user",
		Email:    "invalid",
		Password: "password",
	}
	err := userService.CreateUser(createUserRequest)

	assert.NotNil(t, err)
	assert.Equal(t, "error request contains invalid email", err.Error())
}

func Test_WhenCreateUserRequestIsValid_ThenShouldCreateUser(t *testing.T) {
	gormDB, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	userService := NewUserService(gormDB)
	createUserRequest := &request.CreateUserRequest{
		Username: "user",
		Email:    "valid@email.com",
		Password: "password",
	}
	err := userService.CreateUser(createUserRequest)

	assert.Nil(t, err)

	// Validating whether the user was actually created in the db
	user := &models.User{}
	gormDB.First(&user)
	assert.Equal(t, "user", user.Username)
}

func Test_WhenCreateUserRequestIsValidAndUsernameAlreadyExists_ThenShouldReturnUniqueConstraintFailedError(t *testing.T) {
	gormDB, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	userService := NewUserService(gormDB)
	createUserRequest := &request.CreateUserRequest{
		Username: "user",
		Email:    "valid@email.com",
		Password: "password",
	}
	userService.CreateUser(createUserRequest)
	createUserRequest = &request.CreateUserRequest{
		Username: "user",
		Email:    "unique@email.com",
		Password: "password",
	}
	err := userService.CreateUser(createUserRequest)

	assert.NotNil(t, err)
	assert.Equal(t, "UNIQUE constraint failed: users.username", err.Error())
}

func Test_WhenCreateUserRequestIsValidAndEmailAlreadyExists_ThenShouldReturnUniqueConstraintFailedError(t *testing.T) {
	gormDB, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	userService := NewUserService(gormDB)
	createUserRequest := &request.CreateUserRequest{
		Username: "user1",
		Email:    "valid@email.com",
		Password: "password",
	}
	userService.CreateUser(createUserRequest)
	createUserRequest = &request.CreateUserRequest{
		Username: "user2",
		Email:    "valid@email.com",
		Password: "password",
	}
	err := userService.CreateUser(createUserRequest)

	assert.NotNil(t, err)
	assert.Equal(t, "UNIQUE constraint failed: users.email", err.Error())
}
