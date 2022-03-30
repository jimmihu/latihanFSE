package jwt_usecase

import (
	"errors"
	"fmt"
	"latihanFSE/config"
	"latihanFSE/models/entity"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomClaim struct {
	jwt.StandardClaims
	UserID uuid.UUID `json:"user_id"`
	Name   string    `json:"name"`
}

func (j *JwtUsecase) GenerateToken(UserID uuid.UUID, Name string) (string, error) {

	claim := CustomClaim{
		UserID: UserID,
		Name:   Name,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    config.CONFIG["APP_NAME"],
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claim)
	return token.SignedString([]byte(config.CONFIG["SECRET_KEY"]))
}

func (j *JwtUsecase) ValidateToken(authHeader string) (*jwt.Token, error) {

	if !strings.Contains(authHeader, "Bearer") {
		return nil, fmt.Errorf("invalid Authorization header")
	}
	token := strings.Replace(authHeader, "Bearer ", "", 1)
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}
		return []byte(config.CONFIG["SECRET_KEY"]), nil
	})

}

func (j *JwtUsecase) ValidateTokenAndGetPayload(token string) (*entity.JwtPayload, error) {

	validatedToken, err := j.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	claims, ok := validatedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to claim token")
	}

	UserID, _ := uuid.Parse(claims["user_id"].(string))
	payload := entity.JwtPayload{
		UserID: UserID,
		Name:   claims["name"].(string),
	}
	return &payload, nil
}

func (j *JwtUsecase) UserHasAuthorization(userID uuid.UUID, roles []string) (bool, error) {

	UserDetail, result := j.UserRepo.GetUserDetail(userID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if result.Error != nil {

		return false, result.Error
	}

	for _, role := range roles {
		if UserDetail.Role.Title == role {
			return true, nil
		}
	}
	return false, nil
}
