package request

import "github.com/dgrijalva/jwt-go"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWTClaim struct {
	UserID   uint   `json:"userID"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
