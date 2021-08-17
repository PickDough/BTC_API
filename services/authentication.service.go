package services

import (
	"SE_School/models"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type AuthService struct {
}

func (service *AuthService) GenerateToken(userEmail string) (models.Token, error) {
	var err error

	//Specifying claims
	authenticationClaims := jwt.MapClaims{}
	authenticationClaims["authorized"] = true
	authenticationClaims["user_email"] = userEmail
	authenticationClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, authenticationClaims)
	token, err := at.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	if err != nil {
		return models.Token{}, err
	}
	return models.Token{Token: token}, nil
}
