package token

import (
	"github.com/chessnok/GoCalculator/orchestrator/internal/user"
	"github.com/golang-jwt/jwt"
	"time"
)

type TokenManager struct {
	secret []byte
}

func NewTokenManager(secret string) *TokenManager {
	return &TokenManager{secret: []byte(secret)}
}

func (t *TokenManager) GenerateToken(u *user.User) (string, error, time.Time) {
	now := time.Now()
	exp := now.Add(30 * 24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  u.ID,
		"exp": exp.Unix(),
	})
	tokenString, err := token.SignedString(t.secret)
	if err != nil {
		return "", err, time.Time{}
	}
	return tokenString, nil, exp
}

func (t *TokenManager) CheckToken(tokenString string) (bool, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return t.secret, nil
	})
	if err != nil {
		return false, "", err
	}
	if !token.Valid {
		return false, "", nil
	}
	return true, token.Claims.(jwt.MapClaims)["id"].(string), nil
}
