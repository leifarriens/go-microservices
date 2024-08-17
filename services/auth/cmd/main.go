package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/leifarriens/go-microservices/internal/shared"
	_ "github.com/leifarriens/go-microservices/services/auth/docs"
	"github.com/leifarriens/go-microservices/services/auth/handler"
	"github.com/leifarriens/go-microservices/services/auth/service"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

//	@title			Auth Service
//	@version		2.0
//	@description	This is the auth service

// @host	localhost:1324
func main() {
	privateKey := shared.LoadPrivateKey()
	publicKey := shared.LoadPublicKey()

	s := shared.Server(&shared.ServerConfig{
		Swagger: true,
	})

	tokenService := service.NewTokenService(&service.TokenServiceConfig{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	})

	handler.NewHandler(&handler.HandlerConfig{
		E:            s,
		TokenService: tokenService,
	})

	s.Logger.Fatal(s.Start(":1324"))
}
