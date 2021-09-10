package auth

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type Service interface{
	GenerateToken(userId int) (string,error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {

}

func NewService() *jwtService  {
	return &jwtService{}
}

var SECRET_KEY = []byte("SECRET_KEY")

func (s *jwtService) GenerateToken(userId int) (string,error){
	claims := jwt.MapClaims{}
	claims["user_id"]  = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)

	signToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signToken, err
	}

	return signToken, nil
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error)  {
	newToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_,ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY),nil
	})

	if err != nil {
		return newToken, err
	}

	return newToken,nil
}