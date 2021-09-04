package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type Token struct {
	Token string `json:"token"`
}

type AuthService struct {
}

func (service *AuthService) GenerateToken() (Token, error) {
	var err error

	//Specifying claims
	authenticationClaims := jwt.MapClaims{}
	authenticationClaims["authorized"] = true
	authenticationClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, authenticationClaims)
	token, err := at.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	if err != nil {
		return Token{}, err
	}
	return Token{Token: token}, nil
}
