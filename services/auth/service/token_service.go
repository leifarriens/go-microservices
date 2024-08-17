package service

import (
	"context"
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/leifarriens/go-microservices/internal/shared"
)

type Token struct {
	Token   string
	Expires time.Time
}

type TokenService interface {
	GenerateAccessToken(ctx context.Context) (*Token, error)
	GetPublicKey() *rsa.PublicKey
}

type tokenService struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

type TokenServiceConfig struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func NewTokenService(config *TokenServiceConfig) TokenService {
	return &tokenService{
		PrivateKey: config.PrivateKey,
		PublicKey:  config.PublicKey,
	}
}

func (s *tokenService) GenerateAccessToken(ctx context.Context) (*Token, error) {
	accessTokenExpires := time.Now().Add(time.Minute * 5)

	accessTokenClaims := shared.JwtCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "auth-service",
			// Subject:   u.ID,
			ExpiresAt: jwt.NewNumericDate(accessTokenExpires),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, accessTokenClaims)

	signedToken, err := token.SignedString(s.PrivateKey)

	if err != nil {
		return nil, err
	}

	return &Token{
		Token:   signedToken,
		Expires: time.Now().Add(time.Hour * 24 * 7),
	}, nil
}

func (s *tokenService) GetPublicKey() *rsa.PublicKey {
	return s.PublicKey
}
