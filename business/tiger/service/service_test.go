package service

import (
	"fmt"
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

var (
	claims = &request2.JWTClaim{
		UserID:   1,
		Username: "user1",
		Email:    "user1@email.com",
	}
	claims2 = &request2.JWTClaim{
		UserID:   2,
		Username: "user2",
		Email:    "user2@email.com",
	}
	claims3 = &request2.JWTClaim{
		UserID:   3,
		Username: "user3",
		Email:    "user3@email.com",
	}
)

func setupTests() (*gorm.DB, chan *request.SendTigerSightingNotificationRequest, func(t *testing.T)) {
	db, err := gorm.Open(sqlite.Open("test.db?_foreign_keys=on"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models2.Tiger{})
	db.AutoMigrate(&models2.TigerSighting{})
	db.Create(&models.User{Username: "user1", Email: "user1@email.com", Password: "$2a$04$npZR8DN1y2I0VNRrrPG6XOk.C2lfQLzCOhK5T9lR40oQuecSEHkhm"})
	db.Create(&models.User{Username: "user2", Email: "user2@email.com", Password: "$2a$04$npZR8DN1y2I0VNRrrPG6XOk.C2lfQLzCOhK5T9lR40oQuecSEHkhm"})
	db.Create(&models.User{Username: "user3", Email: "user3@email.com", Password: "$2a$04$npZR8DN1y2I0VNRrrPG6XOk.C2lfQLzCOhK5T9lR40oQuecSEHkhm"})

	tigerSightingNotificationChannel := make(chan *request.SendTigerSightingNotificationRequest, 1000)

	return db, tigerSightingNotificationChannel, func(t *testing.T) {
		err := os.Remove("test.db")
		if err != nil {
			t.Log(err)
		}
	}
}

func Test_WhenCreateTigerRequestContainsEmptyName_ThenReturnErrEmptyParams(t *testing.T) {
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
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
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
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
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
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
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
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
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
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
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
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
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
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

	listAllTigersResponse, err := tigerService.ListAllTigers(&request.ListAllTigerRequest{
		Offset:   2,
		PageSize: 2,
	})

	assert.Nil(t, err)
	assert.Equal(t, "tiger5", listAllTigersResponse.Payload.Tigers[0].Name)
	assert.Equal(t, "tiger2", listAllTigersResponse.Payload.Tigers[1].Name)
}

func Test_WhenCreateTigerSightingRequestContainsEmptyImageBlob_ThenReturnErrEmptyImageBlob(t *testing.T) {
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
	createTigerSightingRequest := &request.CreateTigerSightingRequest{
		Image:     "",
		Latitude:  -1,
		Longitude: -1,
		Timestamp: 0,
	}
	_, err := tigerService.CreateTigerSighting(1, createTigerSightingRequest, claims)

	assert.NotNil(t, err)
	assert.Equal(t, "error request contains empty image blob", err.Error())
}

func Test_WhenCreateTigerSightingRequestContainsInvalidLatitude_ThenReturnErrInvalidLatitude(t *testing.T) {
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
	createTigerSightingRequest := &request.CreateTigerSightingRequest{
		Image:     "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
		Latitude:  -91,
		Longitude: -1,
		Timestamp: 0,
	}
	_, err := tigerService.CreateTigerSighting(1, createTigerSightingRequest, claims)

	assert.NotNil(t, err)
	assert.Equal(t, "error request contains invalid latitude", err.Error())
}

func Test_WhenCreateTigerSightingRequestContainsInvalidLongitude_ThenReturnErrInvalidLongitude(t *testing.T) {
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
	createTigerSightingRequest := &request.CreateTigerSightingRequest{
		Image:     "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
		Latitude:  -90,
		Longitude: -181,
		Timestamp: 0,
	}
	_, err := tigerService.CreateTigerSighting(1, createTigerSightingRequest, claims)

	assert.NotNil(t, err)
	assert.Equal(t, "error request contains invalid longitude", err.Error())
}

func Test_WhenTigerIDIsInvalidInCreateTigerSightingRequest_ThenShouldReturnErr(t *testing.T) {
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
	createTigerSightingRequest := &request.CreateTigerSightingRequest{
		Image:     "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
		Latitude:  -90,
		Longitude: -180,
		Timestamp: 0,
	}
	_, err := tigerService.CreateTigerSighting(1, createTigerSightingRequest, claims)

	assert.NotNil(t, err)
	assert.Equal(t, "record not found", err.Error())
}

func Test_WhenCreateTigerSightingRequestIsValid_ThenShouldSaveTigerSightingInDB(t *testing.T) {
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
	createTigerRequest := &request.CreateTigerRequest{
		Name:              "tiger1",
		DateOfBirth:       "2020-01-13",
		LastSeenLatitude:  -90,
		LastSeenLongitude: -180,
		LastSeenTimestamp: 0,
	}
	tiger, _ := tigerService.CreateTiger(createTigerRequest, claims)
	createTigerSightingRequest := &request.CreateTigerSightingRequest{
		Image:     "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
		Latitude:  -83,
		Longitude: -180,
		Timestamp: 0,
	}
	tigerSighting, err := tigerService.CreateTigerSighting(1, createTigerSightingRequest, claims)

	assert.Nil(t, err)
	assert.NotNil(t, tigerSighting)

	actualTigerSightingInDB := &models2.TigerSighting{}
	gormDB.First(&actualTigerSightingInDB)
	assert.Equal(t, tiger.ID, actualTigerSightingInDB.TigerID)
}

func Test_WhenTigerIDIsInvalidInListAllTigerSightingsRequest_ThenShouldReturnErr(t *testing.T) {
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
	_, err := tigerService.ListAllSightingsForATiger(&request.ListAllTigerSightingsRequest{
		TigerID:  1,
		Offset:   0,
		PageSize: 2,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "record not found", err.Error())
}

func Test_WhenOffsetIs2AndPageSizeIs2_ThenShouldReturnTigerSightingsInCorrectOrder(t *testing.T) {
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
	tiger, _ := tigerService.CreateTiger(&request.CreateTigerRequest{
		Name:              "tiger1",
		DateOfBirth:       "2020-01-13",
		LastSeenLatitude:  -90,
		LastSeenLongitude: -180,
		LastSeenTimestamp: 5,
	}, claims)
	createTigerSightingRequest1 := &request.CreateTigerSightingRequest{
		Image:     "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
		Latitude:  -80,
		Longitude: -180,
		Timestamp: 6,
	}
	createTigerSightingRequest2 := &request.CreateTigerSightingRequest{
		Image:     "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
		Latitude:  -70,
		Longitude: -180,
		Timestamp: 7,
	}
	createTigerSightingRequest3 := &request.CreateTigerSightingRequest{
		Image:     "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
		Latitude:  -60,
		Longitude: -180,
		Timestamp: 8,
	}
	createTigerSightingRequest4 := &request.CreateTigerSightingRequest{
		Image:     "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
		Latitude:  -50,
		Longitude: -180,
		Timestamp: 9,
	}
	tigerService.CreateTigerSighting(tiger.ID, createTigerSightingRequest1, claims)
	tigerService.CreateTigerSighting(tiger.ID, createTigerSightingRequest2, claims)
	tigerService.CreateTigerSighting(tiger.ID, createTigerSightingRequest3, claims)
	tigerService.CreateTigerSighting(tiger.ID, createTigerSightingRequest4, claims)

	listAllTigerSightingsResponse, err := tigerService.ListAllSightingsForATiger(&request.ListAllTigerSightingsRequest{
		TigerID:  int(tiger.ID),
		Offset:   2,
		PageSize: 2,
	})

	assert.Nil(t, err)
	assert.Equal(t, 7, listAllTigerSightingsResponse.Payload.TigerSightings[0].Timestamp)
	assert.Equal(t, 6, listAllTigerSightingsResponse.Payload.TigerSightings[1].Timestamp)
}

func Test_WhenCreateTigerIsValid_ThenShouldSaveTigerAndTigerSightingInDB(t *testing.T) {
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
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

	actualTigerSightingInDB := &models2.TigerSighting{}
	gormDB.First(&actualTigerSightingInDB)
	assert.Equal(t, tiger.ID, actualTigerSightingInDB.TigerID)
}

func Test_WhenCreateTigerSightingIsValid_ThenShouldSaveTigerSightingAndUpdateTigerInDB(t *testing.T) {
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
	createTigerRequest := &request.CreateTigerRequest{
		Name:              "tiger1",
		DateOfBirth:       "2020-01-13",
		LastSeenLatitude:  -90,
		LastSeenLongitude: -180,
		LastSeenTimestamp: 0,
	}
	tiger, err := tigerService.CreateTiger(createTigerRequest, claims)
	assert.Nil(t, err)
	assert.Equal(t, float64(-90), tiger.LastSeenLatitude)
	assert.Equal(t, float64(-180), tiger.LastSeenLongitude)

	createTigerSightingRequest := &request.CreateTigerSightingRequest{
		Latitude:  0,
		Longitude: 0,
		Timestamp: 20,
		Image:     "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
	}
	tigerSighting, err := tigerService.CreateTigerSighting(tiger.ID, createTigerSightingRequest, claims)

	assert.Nil(t, err)
	assert.NotNil(t, tigerSighting)

	actualTigerInDB := &models2.Tiger{}
	gormDB.First(&actualTigerInDB)
	assert.Equal(t, float64(0), actualTigerInDB.LastSeenLatitude)
	assert.Equal(t, float64(0), actualTigerInDB.LastSeenLongitude)

	actualTigerSightingInDB := &models2.TigerSighting{}
	gormDB.First(&actualTigerSightingInDB)
	assert.Equal(t, tiger.ID, actualTigerSightingInDB.TigerID)
}

func Test_WhenCreateTigerSightingIsValidButTigerWithin5KM_ThenShouldReturnErrTigerWithin5KM(t *testing.T) {
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
	createTigerRequest := &request.CreateTigerRequest{
		Name:              "tiger1",
		DateOfBirth:       "2020-01-13",
		LastSeenLatitude:  0,
		LastSeenLongitude: 0,
		LastSeenTimestamp: 0,
	}
	tiger, _ := tigerService.CreateTiger(createTigerRequest, claims)
	createTigerSightingRequest := &request.CreateTigerSightingRequest{
		Latitude:  0.044,
		Longitude: 0,
		Timestamp: 20,
		Image:     "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
	}
	_, err := tigerService.CreateTigerSighting(tiger.ID, createTigerSightingRequest, claims)

	assert.NotNil(t, err)
	assert.Equal(t, "error tiger within 5km from last seen location", err.Error())
}

func Test_WhenCreateTigerSightingIsCalled_ThenShouldSendNotificationToAllReportersForThatTiger(t *testing.T) {
	gormDB, tigerSightingNotificationChannel, teardownTestCase := setupTests()
	defer teardownTestCase(t)

	tigerService := NewTigerService(gormDB, tigerSightingNotificationChannel)
	createTigerRequest := &request.CreateTigerRequest{
		Name:              "tiger1",
		DateOfBirth:       "2020-01-13",
		LastSeenLatitude:  0,
		LastSeenLongitude: 0,
		LastSeenTimestamp: 0,
	}
	tiger, _ := tigerService.CreateTiger(createTigerRequest, claims)
	createTigerSightingRequest1 := &request.CreateTigerSightingRequest{
		Latitude:  0.16,
		Longitude: 0,
		Timestamp: 20,
		Image:     "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
	}
	tigerService.CreateTigerSighting(tiger.ID, createTigerSightingRequest1, claims)
	createTigerSightingRequest2 := &request.CreateTigerSightingRequest{
		Latitude:  0.26,
		Longitude: 0,
		Timestamp: 21,
		Image:     "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
	}
	tigerService.CreateTigerSighting(tiger.ID, createTigerSightingRequest2, claims2)
	createTigerSightingRequest3 := &request.CreateTigerSightingRequest{
		Latitude:  0.36,
		Longitude: 0,
		Timestamp: 22,
		Image:     "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk+A8AAQUBAScY42YAAAAASUVORK5CYII=",
	}
	tigerService.CreateTigerSighting(tiger.ID, createTigerSightingRequest3, claims3)

	for i := 1; i <= 3; i++ {
		select {
		case message := <-tigerSightingNotificationChannel:
			assert.Equal(t, i, len(message.Reporters))
			assert.Equal(t, "Sighting Reported for tiger1", message.Subject)
			assert.Equal(t, fmt.Sprintf("tiger1 is reported to be sighted at 0.%d6 Latitude, 0 Longitude", i), message.Message)
		}
	}
}
