package service

import (
	request2 "github.com/sourava/tiger/business/auth/request"
	"github.com/sourava/tiger/business/tiger/models"
	"github.com/sourava/tiger/business/tiger/request"
	"github.com/sourava/tiger/business/tiger/validations"
	"github.com/sourava/tiger/external/customErrors"
	"gorm.io/gorm"
	"net/http"
)

type TigerService struct {
	db *gorm.DB
}

func NewTigerService(db *gorm.DB) *TigerService {
	return &TigerService{
		db: db,
	}
}

func (service *TigerService) CreateTiger(request *request.CreateTigerRequest, claims *request2.JWTClaim) (*models.Tiger, *customErrors.CustomError) {
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

	result := service.db.Create(&tiger)
	if result.Error != nil {
		return nil, customErrors.NewWithErr(http.StatusInternalServerError, result.Error)
	}

	return tiger, nil
}

func (service *TigerService) ListAllTigers(request *request.ListAllTigerRequest) ([]*models.Tiger, *customErrors.CustomError) {
	var tigers []*models.Tiger

	result := service.db.Offset(request.Offset).Limit(request.PageSize).Order("last_seen_timestamp desc").Find(&tigers)
	if result.Error != nil {
		return nil, customErrors.NewWithErr(http.StatusInternalServerError, result.Error)
	}

	return tigers, nil
}

func (service *TigerService) CreateTigerSighting(request *request.CreateTigerSightingRequest, claims *request2.JWTClaim) (*models.TigerSighting, *customErrors.CustomError) {
	err := validations.ValidateCreateTigerSightingRequest(request)
	if err != nil {
		return nil, err
	}

	var tiger *models.Tiger
	result := service.db.First(&tiger, request.TigerID)
	if result.Error != nil {
		return nil, customErrors.NewWithErr(http.StatusInternalServerError, result.Error)
	}

	tigerSighting := &models.TigerSighting{
		UserID:    claims.UserID,
		TigerID:   request.TigerID,
		Timestamp: request.Timestamp,
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
	}

	result = service.db.Create(&tigerSighting)
	if result.Error != nil {
		return nil, customErrors.NewWithErr(http.StatusInternalServerError, result.Error)
	}

	return tigerSighting, nil
}
