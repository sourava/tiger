package middlewares

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sourava/tiger/business/auth/request"
	"github.com/sourava/tiger/external/customErrors"
	"github.com/sourava/tiger/external/utils"
	"net/http"
	"time"
)

var (
	ErrRequestContainsInvalidToken = customErrors.NewWithMessage(http.StatusUnauthorized, "error invalid token")
)

func Auth(secret string) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			utils.ReturnError(context, ErrRequestContainsInvalidToken)
			context.Abort()
			return
		}

		err := ValidateToken(tokenString, secret)

		if err != nil {
			utils.ReturnError(context, ErrRequestContainsInvalidToken)
			context.Abort()
			return
		}
		context.Next()
	}
}

func GenerateToken(email string, username string, secret string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &request.JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(signedToken string, secret string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&request.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*request.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
