package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
func VerifyToken(tokenString string) error {
	parsedtoken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secretKey), nil
	})
	tokenIsValid := parsedtoken.Valid
	if err != nil {
		return errors.New("Could not parse the Token")
	}
	if !tokenIsValid {
		return errors.New("Invalid Token")
	}
	_, ok := parsedtoken.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("Invalid token Claims")
	}

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)
	return nil
}
