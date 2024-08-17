package shared

import (
	"crypto/rsa"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	jwt.RegisteredClaims
}

func Authorize(publicKey *rsa.PublicKey) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		TokenLookup:   "header:Authorization:Bearer ,cookie:accessToken",
		SigningMethod: "RS256",
		SigningKey:    publicKey,
		ContextKey:    "user",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
	})
}

func LoadPrivateKey() *rsa.PrivateKey {
	bs, err := os.ReadFile("rsa_private.pem")

	if err != nil {
		log.Fatalf("could not read private key pem file: %s", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(bs)

	if err != nil {
		log.Fatalf("unable to parse private key from pem file: %s", err)
	}

	return privateKey
}

func LoadPublicKey() *rsa.PublicKey {
	bs, err := os.ReadFile("rsa_public.pem")

	if err != nil {
		log.Fatalf("could not read public key pem file: %s", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(bs)

	if err != nil {
		log.Fatalf("unable to parse public key from pem file: %s", err)
	}

	return publicKey
}
