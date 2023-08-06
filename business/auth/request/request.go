package request

import "github.com/dgrijalva/jwt-go"

type LoginRequest struct {
	Email    string `json:"email" example:"user@gmail.com"`
	Password string `json:"password" example:"password"`
}

type JWTClaim struct {
	UserID   uint   `json:"userID"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
