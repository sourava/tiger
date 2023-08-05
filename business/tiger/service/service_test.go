package service

import (
	request2 "github.com/sourava/tiger/business/auth/request"
	models2 "github.com/sourava/tiger/business/tiger/models"
	"github.com/sourava/tiger/business/tiger/request"
	"github.com/sourava/tiger/business/user/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"testing"
)

func setupTests() (*gorm.DB, *request2.JWTClaim, func(t *testing.T)) {
	db, err := gorm.Open(sqlite.Open("test.db?_foreign_keys=on"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models2.Tiger{})
	db.AutoMigrate(&models2.TigerSighting{})
	db.Create(&models.User{Username: "user1", Email: "user1@email.com", Password: "$2a$04$npZR8DN1y2I0VNRrrPG6XOk.C2lfQLzCOhK5T9lR40oQuecSEHkhm"})

	claims := &request2.JWTClaim{
		UserID:   1,
		Username: "user1",
		Email:    "user1@email.com",
	}

	return db, claims, func(t *testing.T) {
		err := os.Remove("test.db")
		if err != nil {
			t.Log(err)
		}
	}
}

func Test_WhenCreateTigerRequestContainsEmptyName_ThenReturnErrEmptyParams(t *testing.T) {
	gormDB, claims, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB)
	createTigerRequest := &request.CreateTigerRequest{
		Name:              "",
		DateOfBirth:       "",
		LastSeenLatitude:  -1,
		LastSeenLongitude: -1,
		LastSeenTimestamp: 0,
	}
	_, err := tigerService.CreateTiger(createTigerRequest, claims)

	assert.NotNil(t, err)
	assert.Equal(t, "error request contains empty name or date of birth", err.Error())
}

func Test_WhenCreateTigerRequestContainsEmptyDateOfBirth_ThenReturnErrEmptyParams(t *testing.T) {
	gormDB, claims, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB)
	createTigerRequest := &request.CreateTigerRequest{
		Name:              "tiger1",
		DateOfBirth:       "",
		LastSeenLatitude:  -1,
		LastSeenLongitude: -1,
		LastSeenTimestamp: 0,
	}
	_, err := tigerService.CreateTiger(createTigerRequest, claims)

	assert.NotNil(t, err)
	assert.Equal(t, "error request contains empty name or date of birth", err.Error())
}

func Test_WhenCreateTigerRequestContainsInvalidDateOfBirth_ThenReturnErrInvalidDateOfBirth(t *testing.T) {
	gormDB, claims, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB)
	createTigerRequest := &request.CreateTigerRequest{
		Name:              "tiger1",
		DateOfBirth:       "invalid date",
		LastSeenLatitude:  -1,
		LastSeenLongitude: -1,
		LastSeenTimestamp: 0,
	}
	_, err := tigerService.CreateTiger(createTigerRequest, claims)

	assert.NotNil(t, err)
	assert.Equal(t, "error request contains invalid date of birth, format = YYYY-MM-DD", err.Error())
}

func Test_WhenCreateTigerRequestContainsInvalidLatitude_ThenReturnErrInvalidLatitude(t *testing.T) {
	gormDB, claims, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB)
	createTigerRequest := &request.CreateTigerRequest{
		Name:              "tiger1",
		DateOfBirth:       "2020-01-13",
		LastSeenLatitude:  -91,
		LastSeenLongitude: -1,
		LastSeenTimestamp: 0,
	}
	_, err := tigerService.CreateTiger(createTigerRequest, claims)

	assert.NotNil(t, err)
	assert.Equal(t, "error request contains invalid latitude", err.Error())
}

func Test_WhenCreateTigerRequestContainsInvalidLongitude_ThenReturnErrInvalidLongitude(t *testing.T) {
	gormDB, claims, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB)
	createTigerRequest := &request.CreateTigerRequest{
		Name:              "tiger1",
		DateOfBirth:       "2020-01-13",
		LastSeenLatitude:  -90,
		LastSeenLongitude: -181,
		LastSeenTimestamp: 0,
	}
	_, err := tigerService.CreateTiger(createTigerRequest, claims)

	assert.NotNil(t, err)
	assert.Equal(t, "error request contains invalid longitude", err.Error())
}

func Test_WhenCreateTigerRequestIsValid_ThenShouldSaveTigerInDB(t *testing.T) {
	gormDB, claims, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB)
	createTigerRequest := &request.CreateTigerRequest{
		Name:              "tiger1",
		DateOfBirth:       "2020-01-13",
		LastSeenLatitude:  -90,
		LastSeenLongitude: -180,
		LastSeenTimestamp: 0,
	}
	tiger, err := tigerService.CreateTiger(createTigerRequest, claims)

	assert.Nil(t, err)
	assert.NotNil(t, tiger)

	actualTigerInDB := &models2.Tiger{}
	gormDB.First(&actualTigerInDB)
	assert.Equal(t, "tiger1", actualTigerInDB.Name)
}

func Test_WhenOffsetIs2AndPageSizeIs2_ThenShouldReturnTigersInCorrectOrder(t *testing.T) {
	gormDB, claims, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB)
	tiger1 := &request.CreateTigerRequest{
		Name:              "tiger1",
		DateOfBirth:       "2020-01-13",
		LastSeenLatitude:  -90,
		LastSeenLongitude: -180,
		LastSeenTimestamp: 5,
	}
	tiger2 := &request.CreateTigerRequest{
		Name:              "tiger2",
		DateOfBirth:       "2020-01-13",
		LastSeenLatitude:  -90,
		LastSeenLongitude: -180,
		LastSeenTimestamp: 2,
	}
	tiger3 := &request.CreateTigerRequest{
		Name:              "tiger3",
		DateOfBirth:       "2020-01-13",
		LastSeenLatitude:  -90,
		LastSeenLongitude: -180,
		LastSeenTimestamp: 4,
	}
	tiger4 := &request.CreateTigerRequest{
		Name:              "tiger4",
		DateOfBirth:       "2020-01-13",
		LastSeenLatitude:  -90,
		LastSeenLongitude: -180,
		LastSeenTimestamp: 1,
	}
	tiger5 := &request.CreateTigerRequest{
		Name:              "tiger5",
		DateOfBirth:       "2020-01-13",
		LastSeenLatitude:  -90,
		LastSeenLongitude: -180,
		LastSeenTimestamp: 3,
	}
	tigerService.CreateTiger(tiger1, claims)
	tigerService.CreateTiger(tiger2, claims)
	tigerService.CreateTiger(tiger3, claims)
	tigerService.CreateTiger(tiger4, claims)
	tigerService.CreateTiger(tiger5, claims)

	tigers, err := tigerService.ListAllTigers(&request.ListAllTigerRequest{
		Offset:   2,
		PageSize: 2,
	})

	assert.Nil(t, err)
	assert.Equal(t, "tiger5", tigers[0].Name)
	assert.Equal(t, "tiger2", tigers[1].Name)
}

func Test_WhenCreateTigerSightingRequestContainsEmptyImageBlob_ThenReturnErrEmptyImageBlob(t *testing.T) {
	gormDB, claims, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB)
	createTigerSightingRequest := &request.CreateTigerSightingRequest{
		TigerID:   1,
		Image:     "",
		Latitude:  -1,
		Longitude: -1,
		Timestamp: 0,
	}
	_, err := tigerService.CreateTigerSighting(createTigerSightingRequest, claims)

	assert.NotNil(t, err)
	assert.Equal(t, "error request contains empty image blob", err.Error())
}

func Test_WhenCreateTigerSightingRequestContainsInvalidLatitude_ThenReturnErrInvalidLatitude(t *testing.T) {
	gormDB, claims, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB)
	createTigerSightingRequest := &request.CreateTigerSightingRequest{
		TigerID:   1,
		Image:     "imageblob",
		Latitude:  -91,
		Longitude: -1,
		Timestamp: 0,
	}
	_, err := tigerService.CreateTigerSighting(createTigerSightingRequest, claims)

	assert.NotNil(t, err)
	assert.Equal(t, "error request contains invalid latitude", err.Error())
}

func Test_WhenCreateTigerSightingRequestContainsInvalidLongitude_ThenReturnErrInvalidLongitude(t *testing.T) {
	gormDB, claims, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB)
	createTigerSightingRequest := &request.CreateTigerSightingRequest{
		TigerID:   1,
		Image:     "imageblob",
		Latitude:  -90,
		Longitude: -181,
		Timestamp: 0,
	}
	_, err := tigerService.CreateTigerSighting(createTigerSightingRequest, claims)

	assert.NotNil(t, err)
	assert.Equal(t, "error request contains invalid longitude", err.Error())
}

func Test_WhenTigerIDIsInvalidInCreateTigerSightingRequest_ThenShouldReturnErr(t *testing.T) {
	gormDB, claims, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB)
	createTigerSightingRequest := &request.CreateTigerSightingRequest{
		TigerID:   1,
		Image:     "imageblob",
		Latitude:  -90,
		Longitude: -180,
		Timestamp: 0,
	}
	_, err := tigerService.CreateTigerSighting(createTigerSightingRequest, claims)

	assert.NotNil(t, err)
	assert.Equal(t, "record not found", err.Error())
}

func Test_WhenCreateTigerSightingRequestIsValid_ThenShouldSaveTigerSightingInDB(t *testing.T) {
	gormDB, claims, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB)
	createTigerRequest := &request.CreateTigerRequest{
		Name:              "tiger1",
		DateOfBirth:       "2020-01-13",
		LastSeenLatitude:  -90,
		LastSeenLongitude: -180,
		LastSeenTimestamp: 0,
	}
	tiger, _ := tigerService.CreateTiger(createTigerRequest, claims)
	createTigerSightingRequest := &request.CreateTigerSightingRequest{
		TigerID:   tiger.ID,
		Image:     "imageblob",
		Latitude:  -90,
		Longitude: -180,
		Timestamp: 0,
	}
	tigerSighting, err := tigerService.CreateTigerSighting(createTigerSightingRequest, claims)

	assert.Nil(t, err)
	assert.NotNil(t, tigerSighting)

	actualTigerSightingInDB := &models2.TigerSighting{}
	gormDB.First(&actualTigerSightingInDB)
	assert.Equal(t, tiger.ID, actualTigerSightingInDB.TigerID)
}
