package constants

import (
	"github.com/sourava/tiger/external/customErrors"
	"net/http"
)

var (
	ErrInvalidEmail          = customErrors.NewWithMessage(http.StatusBadRequest, "error request contains invalid email")
	ErrEmailPasswordMismatch = customErrors.NewWithMessage(http.StatusBadRequest, "error email and password combination mismatch")
	ErrGeneratingToken       = customErrors.NewWithMessage(http.StatusBadRequest, "error generating token")
)
