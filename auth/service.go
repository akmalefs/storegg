package auth

import "github.com/golang-jwt/jwt/v5"

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
	secretKey []byte
}

func NewService(secret string) *jwtService {
	return &jwtService{
		secretKey: []byte(secret),
	}
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(s.secretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, err
}
