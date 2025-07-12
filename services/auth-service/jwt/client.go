package jwt

import (
	"time"

	"main/services/auth-service/internal/values"

	"github.com/golang-jwt/jwt"
)

type Kind string

const (
	Access  Kind = "access"
	Refresh Kind = "refresh"
	Confirm Kind = "confirm"
)

type JWT interface {
	GenerateToken(string, Kind) (string, int, error)
	ParseJWTToken(string) (*TokenClaims, error)
}

type Config struct {
	APISecret     string
	RefreshSecret string
	Issuer        string
	AccessTTL     int
	RefreshTTL    int
}

type jwtClient struct {
	secret        string
	refreshSecret string
	issuer        string
	accessTTL     time.Duration
	refreshTTL    time.Duration
}

type TokenClaims struct {
	jwt.StandardClaims
	SecretType Kind
}

func (j *jwtClient) GenerateToken(id string, secretType Kind) (string, int, error) {
	var expiresAt time.Duration
	var secret string

	switch secretType {
	case Access:
		secret = j.secret
		expiresAt = j.accessTTL
	case Refresh, Confirm:
		secret = j.refreshSecret
		expiresAt = j.refreshTTL
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			Subject:   id,
			ExpiresAt: time.Now().Add(expiresAt).Unix(),
		},
		secretType,
	})

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, values.ErrInvalidSigningToken
	}

	return signedToken, int(expiresAt.Minutes()), nil
}

func (j *jwtClient) ParseJWTToken(accessToken string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, values.ErrInvalidSigningMethod
			}

			var secret string
			claims, ok := token.Claims.(*TokenClaims)
			if ok {
				switch claims.SecretType {
				case Access:
					secret = j.secret
				case Refresh, Confirm:
					secret = j.refreshSecret
				}
			}
			return []byte(secret), nil
		})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, values.ErrInvalidToken
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return nil, values.ErrInvalidClaims
	}

	return claims, nil
}

func New(cfg Config) JWT {
	return &jwtClient{
		secret:        cfg.APISecret,
		refreshSecret: cfg.RefreshSecret,
		issuer:        cfg.Issuer,
		accessTTL:     time.Minute * time.Duration(cfg.AccessTTL),
		refreshTTL:    time.Minute * time.Duration(cfg.RefreshTTL),
	}
}
