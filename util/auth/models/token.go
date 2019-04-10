package models

import (
	"github.com/dgrijalva/jwt-go"
)

// Token struct for JWT
type Token struct {
	UserID uint
	jwt.StandardClaims
}
