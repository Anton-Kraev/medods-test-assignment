package auth

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type TokenManager struct {
	jwtSign string
}

func NewTokenManager(jwtSign string) TokenManager {
	return TokenManager{jwtSign: jwtSign}
}

type Claims struct {
	UserID string `json:"user_id"`
	UserIP string `json:"ip_address"`
	jwt.RegisteredClaims
}

func (t TokenManager) GenerateAccessToken(userID, userIP string) (string, error) {
	claims := Claims{
		UserID: userID,
		UserIP: userIP,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := token.SignedString(t.jwtSign)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (t TokenManager) GenerateRefreshToken() (string, string, error) {
	refreshToken := fmt.Sprintf("%d", time.Now().UnixNano())
	encoded := base64.StdEncoding.EncodeToString([]byte(refreshToken))

	hashed, err := bcrypt.GenerateFromPassword([]byte(encoded), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	return encoded, string(hashed), nil
}

func (t TokenManager) CompareRefreshTokens(token string, tokenHash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(tokenHash), []byte(token)) == nil
}
