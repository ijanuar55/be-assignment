package entity

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type User struct {
	Id        string
	Email     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserResponse struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserPostRequest struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
	Name     string `validate:"required" json:"name"`
}

type UserUpdateRequest struct {
	Id    string `validate:"required" json:"id"`
	Email string `validate:"required" json:"email"`
	Name  string `validate:"required" json:"name"`
}

type Claims struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}
