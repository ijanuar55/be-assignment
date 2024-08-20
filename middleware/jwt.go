package middleware

import (
	"be-assignment/entity"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret123!@#009")

func GenerateJWT(email, userId string) (string, error) {

	// Set expiration time of the token
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the claims, which includes the username and expiry time
	claims := &entity.Claims{
		UserId: userId,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create the token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "No Authorization header provided"})
			c.Abort()
			return
		}

		claims := &entity.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		// Access claims data
		c.Set("email", claims.Email)
		c.Set("user_id", claims.UserId)
		c.Next()
	}
}
