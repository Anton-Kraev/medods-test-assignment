package auth

import (
	"context"
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/Anton-Kraev/medods-test-assignment/internal/models/auth"
	"github.com/Anton-Kraev/medods-test-assignment/internal/models/user"
)

type (
	authService interface {
		GenerateTokens(
			ctx context.Context, user user.User,
		) (
			auth.AccessToken, auth.RefreshToken, error,
		)

		RefreshTokens(
			ctx context.Context, userID string, userIP string, token auth.RefreshToken,
		) (
			auth.AccessToken, auth.RefreshToken, error,
		)
	}

	Handler struct {
		service authService
		log     *slog.Logger
	}
)

func NewHandler(service authService, log *slog.Logger) Handler {
	return Handler{
		service: service,
		log:     log,
	}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	authRoute := router.Group("/auth")
	authRoute.POST("/", h.GenerateTokens)
	authRoute.POST("/refresh", h.RefreshTokens)

	return router
}

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
