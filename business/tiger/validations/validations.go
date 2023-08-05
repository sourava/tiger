package validations

import (
	"github.com/sourava/tiger/business/tiger/constants"
	"github.com/sourava/tiger/business/tiger/request"
	"github.com/sourava/tiger/external/customErrors"
	"time"
)

func ValidateCreateTigerRequest(tigerRequest *request.CreateTigerRequest) *customErrors.CustomError {
	if tigerRequest.Name == "" || tigerRequest.DateOfBirth == "" {
		return constants.ErrEmptyNameOrDateOfBirth
	} else if !ValidateDateString(tigerRequest.DateOfBirth) {
		return constants.ErrInvalidDateOfBirth
	} else if !ValidateLatitude(tigerRequest.LastSeenLatitude) {
		return constants.ErrInvalidLatitude
	} else if !ValidateLongitude(tigerRequest.LastSeenLongitude) {
		return constants.ErrInvalidLongitude
	}

	return nil
}

func ValidateCreateTigerSightingRequest(tigerSightingRequest *request.CreateTigerSightingRequest) *customErrors.CustomError {
	if tigerSightingRequest.Image == "" {
		return constants.ErrEmptyImageBlob
	} else if !ValidateLatitude(tigerSightingRequest.Latitude) {
		return constants.ErrInvalidLatitude
	} else if !ValidateLongitude(tigerSightingRequest.Longitude) {
		return constants.ErrInvalidLongitude
	}

	return nil
}

func ValidateDateString(dateString string) bool {
	_, err := time.Parse(constants.DateOfBirthLayout, dateString)
	return err == nil
}

func ValidateLatitude(latitude float64) bool {
	return latitude >= -90 && latitude <= 90
}

func ValidateLongitude(longitude float64) bool {
	return longitude >= -180 && longitude <= 180
}
