package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	authhandler "github.com/Anton-Kraev/medods-test-assignment/internal/handler/auth"
	"github.com/Anton-Kraev/medods-test-assignment/internal/repository/session"
	authservice "github.com/Anton-Kraev/medods-test-assignment/internal/service/auth"
	"github.com/Anton-Kraev/medods-test-assignment/pkg/auth"
	"github.com/Anton-Kraev/medods-test-assignment/pkg/logger"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := pgxpool.New(ctx, "")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	customLog := logger.Setup("local")

	tokenManager := auth.NewTokenManager("default")

	repository := session.NewRepository(db)
	service := authservice.NewService(repository, nil, tokenManager)
	handler := authhandler.NewHandler(service, customLog)

	if err = handler.InitRoutes().Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
