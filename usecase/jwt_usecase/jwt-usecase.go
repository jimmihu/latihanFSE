package jwt_usecase

import (
	"latihanFSE/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type CustomClaim struct {
	jwt.StandardClaims
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
}

func (jwtAuth *JwtUsecase) GenerateToken(UserID uuid.UUID, Email string) (string, error) {

	claim := CustomClaim{
		UserID: UserID,
		Email:  Email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    config.CONFIG["APP_NAME"],
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claim)
	return token.SignedString([]byte(config.CONFIG["SECRET_KEY"]))
}
