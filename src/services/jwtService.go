package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct {
	secretKey string
	assinante string
}

func NewJwtService() *jwtService {
	return &jwtService{
		secretKey: "Jv410551",
		assinante: "RestApi",
	}
}

type Claim struct {
	Sum uint `json:"sum"`
	jwt.StandardClaims
}

func (s *jwtService) GerarToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.assinante,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	fmt.Println(token)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	fmt.Println(t)

	return t, nil
}

func (s *jwtService) ValidarToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ehValido := t.Method.(*jwt.SigningMethodHMAC); !ehValido {
			return nil, fmt.Errorf("token inválido: %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}
