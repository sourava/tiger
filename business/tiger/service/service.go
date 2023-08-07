package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	request2 "github.com/sourava/tiger/business/auth/request"
	"github.com/sourava/tiger/business/tiger/constants"
	"github.com/sourava/tiger/business/tiger/models"
	"github.com/sourava/tiger/business/tiger/request"
	"github.com/sourava/tiger/business/tiger/response"
	"github.com/sourava/tiger/business/tiger/validations"
	"github.com/sourava/tiger/external/customErrors"
	"github.com/sourava/tiger/external/utils"
	"gorm.io/gorm"
	"net/http"
)

type TigerService struct {
	db                               *gorm.DB
	tigerSightingNotificationChannel chan<- *request.SendTigerSightingNotificationRequest
}

func NewTigerService(db *gorm.DB, tigerSightingNotificationChannel chan<- *request.SendTigerSightingNotificationRequest) *TigerService {
	return &TigerService{
		db:                               db,
		tigerSightingNotificationChannel: tigerSightingNotificationChannel,
	}
}

func (service *TigerService) CreateTiger(request *request.CreateTigerRequest, claims *request2.JWTClaim) (*response.CreateTigerHandlerResponse, *customErrors.CustomError) {
	err := validations.ValidateCreateTigerRequest(request)
	if err != nil {
		return nil, err
	}

	tiger := &models.Tiger{
		UserID:            claims.UserID,
		Name:              request.Name,
		DateOfBirth:       request.DateOfBirth,
		LastSeenTimestamp: request.LastSeenTimestamp,
		LastSeenLatitude:  request.LastSeenLatitude,
		LastSeenLongitude: request.LastSeenLongitude,
	}

	transactionErr := service.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(tiger).Error; err != nil {
			return err
		}

		if err := tx.Create(&models.TigerSighting{
			TigerID:   tiger.ID,
			UserID:    claims.UserID,
			Latitude:  tiger.LastSeenLatitude,
			Longitude: tiger.LastSeenLongitude,
			Timestamp: tiger.LastSeenTimestamp,
			Image:     "",
		}).Error; err != nil {
			return err
		}

		return nil
	})

	if transactionErr != nil {
		return nil, customErrors.NewWithErr(http.StatusInternalServerError, transactionErr)
	}

	return response.BuildCreateTigerHandlerResponse(tiger), nil
}

func (service *TigerService) ListAllTigers(request *request.ListAllTigerRequest) (*response.ListAllTigersHandlerResponse, *customErrors.CustomError) {
	var tigers []*response.TigerResponse

	result := service.db.Table("tigers").Offset(request.Offset).Limit(request.PageSize).Order("last_seen_timestamp desc").Find(&tigers)
	if result.Error != nil {
		return nil, customErrors.NewWithErr(http.StatusInternalServerError, result.Error)
	}

	return &response.ListAllTigersHandlerResponse{
		Success: true,
		Payload: response.ListAllTigersResponse{
			Tigers: tigers,
		},
	}, nil
}

func (service *TigerService) CreateTigerSighting(tigerID uint, tigerSightingRequest *request.CreateTigerSightingRequest, claims *request2.JWTClaim) (*models.TigerSighting, *customErrors.CustomError) {
	validationErr := validations.ValidateCreateTigerSightingRequest(tigerSightingRequest)
	if validationErr != nil {
		return nil, validationErr
	}

	resizedImage, err := utils.ResizeImage(tigerSightingRequest.Image, 250, 200)
	if err != nil {
		return nil, customErrors.NewWithErr(http.StatusBadRequest, err)
	}

	var tiger *models.Tiger
	result := service.db.First(&tiger, tigerID)
	if result.Error != nil {
		return nil, customErrors.NewWithErr(http.StatusInternalServerError, result.Error)
	}

	if utils.Distance(tiger.LastSeenLatitude, tiger.LastSeenLongitude, tigerSightingRequest.Latitude, tigerSightingRequest.Longitude) < 5000 {
		return nil, constants.ErrTigerWithin5KM
	}

	tigerSighting := &models.TigerSighting{
		UserID:    claims.UserID,
		TigerID:   tiger.ID,
		Timestamp: tigerSightingRequest.Timestamp,
		Latitude:  tigerSightingRequest.Latitude,
		Longitude: tigerSightingRequest.Longitude,
		Image:     resizedImage,
	}

	err = service.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(tigerSighting).Error; err != nil {
			return err
		}

		if err := tx.Model(tiger).Updates(map[string]interface{}{
			"last_seen_timestamp": tigerSighting.Timestamp,
			"last_seen_latitude":  tigerSighting.Latitude,
			"last_seen_longitude": tigerSighting.Longitude,
		}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, customErrors.NewWithErr(http.StatusInternalServerError, err)
	}

	var tigerSightingReporters []*request.TigerSightingReporter
	err = service.db.
		Model(&models.TigerSighting{}).
		Select("users.username, users.email").
		Joins("left join users on users.id = tiger_sightings.user_id").
		Group("users.id").
		Having("tiger_sightings.tiger_id = ?", tiger.ID).
		Scan(&tigerSightingReporters).Error
	if err != nil {
		log.Error(err)
	}

	service.tigerSightingNotificationChannel <- &request.SendTigerSightingNotificationRequest{
		Reporters: tigerSightingReporters,
		Message:   fmt.Sprintf("%v is reported to be sighted at %v Latitude, %v Longitude", tiger.Name, tigerSighting.Latitude, tigerSighting.Longitude),
		Subject:   fmt.Sprintf("Sighting Reported for %v", tiger.Name),
	}

	return tigerSighting, nil
}

func (service *TigerService) ListAllSightingsForATiger(request *request.ListAllTigerSightingsRequest) (*response.ListAllTigerSightingsHandlerResponse, *customErrors.CustomError) {
	var tiger *models.Tiger
	result := service.db.First(&tiger, request.TigerID)
	if result.Error != nil {
		return nil, customErrors.NewWithErr(http.StatusInternalServerError, result.Error)
	}

	var tigerSightings []*response.TigerSightingResponse
	result = service.db.Table("tiger_sightings").Offset(request.Offset).Limit(request.PageSize).Order("timestamp desc").Find(&tigerSightings)
	if result.Error != nil {
		return nil, customErrors.NewWithErr(http.StatusInternalServerError, result.Error)
	}

	return &response.ListAllTigerSightingsHandlerResponse{
		Success: true,
		Payload: response.ListAllTigerSightingsResponse{
			TigerSightings: tigerSightings,
		},
	}, nil
}
