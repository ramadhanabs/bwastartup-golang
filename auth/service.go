package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type JWTService struct {
}

func NewService() *JWTService {
	return &JWTService{}
}

var SECRET_KEY = []byte("BWASTARTUP_S3CR3T_K3Y")

func (s *JWTService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Signed signature key for security purpose
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return []byte(SECRET_KEY), nil
}

func (s *JWTService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, keyFunc)
	if err != nil {
		return token, err
	}
	return token, nil
}
