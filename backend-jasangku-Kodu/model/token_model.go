package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CreateToken struct {
	Email    string `json:"Email" validate:"required,max=30"`
	Password string `json:"Password" validate:"required"`
}

type ClaimToken struct {
	ID        string
	FirstName string
	LastName  string
	Username  string
	Email     string
	CreatedAt time.Time
	jwt.StandardClaims
}

type TokenResponse struct {
	ID        string    `json:"id"`
	Image     string    `json:"image"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `json:"token"`
}
