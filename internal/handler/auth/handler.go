package auth

import (
	"context"

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
	}
)

func NewHandler(service authService) Handler {
	return Handler{service: service}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	authRoute := router.Group("/auth")
	authRoute.GET("/", h.GenerateTokens)
	authRoute.GET("/refresh", h.RefreshTokens)

	return router
}

func (h Handler) GenerateTokens(c *gin.Context) {
	panic("not implemented")
}

func (h Handler) RefreshTokens(c *gin.Context) {
	panic("not implemented")
}
