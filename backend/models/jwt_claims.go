package models

import "github.com/dgrijalva/jwt-go"

// Claims struct used for JWT authentication
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
