package core_service_jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct{
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

type Manager struct{
	secret []byte
	ttl time.Duration
}

func NewManager(secret string, ttl time.Duration) *Manager{
	return &Manager{
		secret: []byte(secret),
		ttl: ttl,
	}
}

func (m *Manager) Generate(userID int) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.ttl)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(m.secret)
}

func (m *Manager) Parse(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return m.secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error while parse jwt token")
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}