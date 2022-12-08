package dto

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	UserID   int64  `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
}
