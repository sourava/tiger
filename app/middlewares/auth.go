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

		claims, err := ValidateToken(tokenString, secret)

		context.Set("token-claims", claims)

		if err != nil {
			utils.ReturnError(context, ErrRequestContainsInvalidToken)
			context.Abort()
			return
		}
		context.Next()
	}
}

func GenerateToken(userID uint, email string, username string, secret string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &request.JWTClaim{
		UserID:   userID,
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateToken(signedToken string, secret string) (*request.JWTClaim, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&request.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*request.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}

	return claims, nil
}
