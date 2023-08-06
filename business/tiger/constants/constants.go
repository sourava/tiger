package constants

import (
	"github.com/sourava/tiger/external/customErrors"
	"net/http"
)

var (
	ErrEmptyNameOrDateOfBirth = customErrors.NewWithMessage(http.StatusBadRequest, "error request contains empty name or date of birth")
	ErrEmptyImageBlob         = customErrors.NewWithMessage(http.StatusBadRequest, "error request contains empty image blob")
	ErrInvalidLatitude        = customErrors.NewWithMessage(http.StatusBadRequest, "error request contains invalid latitude")
	ErrInvalidLongitude       = customErrors.NewWithMessage(http.StatusBadRequest, "error request contains invalid longitude")
	ErrInvalidDateOfBirth     = customErrors.NewWithMessage(http.StatusBadRequest, "error request contains invalid date of birth, format = YYYY-MM-DD")
	ErrTigerWithin5KM         = customErrors.NewWithMessage(http.StatusBadRequest, "error tiger within 5km from last seen location")
)

const (
	DateOfBirthLayout string = "2006-01-02"
)
