package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"strings"
	"time"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

type JwtService struct {
}

const (
	SecretKey = "random123"
)

func NewJwtService() *JwtService {
	return &JwtService{}
}

// JwtMiddleware make middleware for jwt
func (j *JwtService) JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		header := ctx.Request().Header.Get("Authorization")

		splitString := strings.Split(header, " ")

		if len(splitString) != 2 {
			return ctx.JSON(401, map[string]interface{}{
				"message": "invalid token",
			})
		}

		tokenString := splitString[1]

		_, err := j.VerifyToken(tokenString)
		if err != nil {
			return err
		}
		return next(ctx)
	}
}

func (j *JwtService) VerifyToken(tokenString string) (bool, error) {

	byteSecret := []byte(SecretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return byteSecret, nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, fmt.Errorf("invalid token")
	}

	return true, nil

}

func (j *JwtService) GenerateToken(name, email string, isAdmin bool) (string, error) {
	// Set custom claims
	claims := &JwtCustomClaims{
		Name:  name,
		Email: email,
		Admin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}

	return t, nil
}
