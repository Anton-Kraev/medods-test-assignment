package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	authhandler "github.com/Anton-Kraev/medods-test-assignment/internal/handler/auth"
	"github.com/Anton-Kraev/medods-test-assignment/internal/repository/session"
	authservice "github.com/Anton-Kraev/medods-test-assignment/internal/service/auth"
	"github.com/Anton-Kraev/medods-test-assignment/pkg/auth"
	"github.com/Anton-Kraev/medods-test-assignment/pkg/email"
	"github.com/Anton-Kraev/medods-test-assignment/pkg/logger"
)

const envFilePath = "./.env"

func main() {
	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatalln("can't load environment file", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatalln("DATABASE_URL must be not empty")
	}

	db, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatalln("can't open postgres connection pool", err)
	}
	defer db.Close()

	customLog := logger.Setup(os.Getenv("ENV"))
	emailClient := email.MockClient{}

	jwtExp, err := time.ParseDuration(os.Getenv("JWT_EXP"))
	if err != nil {
		jwtExp = 15 * time.Minute
	}

	tokenManager := auth.NewTokenManager(os.Getenv("JWT_SIGN"), jwtExp)

	repository := session.NewRepository(db)
	service := authservice.NewService(repository, emailClient, tokenManager)
	handler := authhandler.NewHandler(service, customLog)

	if err = handler.InitRoutes().Run(":8080"); err != nil {
		log.Fatalln("can't start http server on port 8080", err)
	}
}
