package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Task : jwt and refresh token related task

type AccessTokenClaims struct {
	UserID     uint    `json:"user_id"`
	CoachingID *uint   `json:"coaching_id,omitempty"`
	Role       string  `json:"role"`
	SessionID  uint    `json:"session_id"`
	TokenType  string  `json:"token_typ"`

	jwt.RegisteredClaims
}

type TokenManager struct {
	AccessSecret []byte
	Issuer        string
	AccessExpires time.Duration
	RefreshExpires time.Duration
}

func NewTokenManager(accessSecret string,issuer string,accessExpires time.Duration,refreshExpires time.Duration,)*TokenManager{
	return &TokenManager{
		AccessSecret: []byte(accessSecret),
		Issuer: issuer,
		AccessExpires: accessExpires,
		RefreshExpires: refreshExpires,
	}
}

func (tm TokenManager) GenerateAccessToken(userId uint, coaschingId *uint, role string, sessionId uint)(string, int64, error) {
	now := time.Now()
	expireAt := now.Add(tm.AccessExpires)

	claim := AccessTokenClaims{
		UserID: userId,
		CoachingID: coaschingId,
		Role: role,
		SessionID: sessionId,
		TokenType: "access",

		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: tm.Issuer,
			Subject: strconv.FormatUint(uint64(userId), 10),
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expireAt),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(tm.AccessSecret)

	if err != nil {
		return "", 0, err
	}

	return signedToken, int64(tm.AccessExpires.Seconds()), nil
}

func (tm TokenManager) VerifyAccessToken(tokenString string) (*AccessTokenClaims, error){
	
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method")
		}
		return tm.AccessSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims,ok := token.Claims.(*AccessTokenClaims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid access token")
	}

	if claims.TokenType != "access" {
		return nil, errors.New("invalid token type")
	}

	return claims,nil
}

func GenerateRefreshToken() (string, error) {
	bytes := make([]byte, 80)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(bytes), nil
}

func HashRefreshToken(token string) string {
	hash := sha256.Sum256([]byte(token))

	return hex.EncodeToString(hash[:])
}