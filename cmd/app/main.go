package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	authhandler "github.com/Anton-Kraev/medods-test-assignment/internal/handler/auth"
	"github.com/Anton-Kraev/medods-test-assignment/internal/logger"
	"github.com/Anton-Kraev/medods-test-assignment/internal/repository/session"
	authservice "github.com/Anton-Kraev/medods-test-assignment/internal/service/auth"
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

	repository := session.NewRepository(db)
	service := authservice.NewService(repository, nil)
	handler := authhandler.NewHandler(service, customLog)

	if err = handler.InitRoutes().Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
