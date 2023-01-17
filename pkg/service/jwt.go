package service

import (
	"errors"
	"fmt"

	"github.com/doduykhang/hermes/chat/pkg/config"
	"github.com/golang-jwt/jwt/v4"
)


type JwtServcie interface {
	Generate(userId string) (string, error)
	Parse(token string) (bool, error)
}

type jwtService struct {
	
}

func NewJwtService() JwtServcie {
	return &jwtService{}
}

func (service *jwtService) Generate(userId string) (string, error) {
	
	var secretKey = []byte(config.GetEnv().JwtSecrect)
	token := jwt.New(jwt.SigningMethodHS256)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	fmt.Println(tokenString)

	return tokenString, nil	
}

func (service *jwtService) Parse(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv().JwtSecrect), nil
         })

	 if err != nil {
		 return false, err
	 }

	 _, ok := token.Claims.(jwt.MapClaims)
        if ok && token.Valid {
                return false, nil
        }

       	return true, errors.New("token not valid")
}


