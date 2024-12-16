package auth

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Anton-Kraev/medods-test-assignment/internal/models/user"
	"github.com/Anton-Kraev/medods-test-assignment/pkg/logger"
)

type generateTokensRequest struct {
	UserID string `json:"user_id" binding:"required,uuid"`
	Email  string `json:"email" binding:"required,email"`
}

func (h Handler) GenerateTokens(c *gin.Context) {
	const errMsg = "failed to generate tokens"

	log := h.log.With(slog.String("endpoint", "Handler.GenerateTokens"))

	var req generateTokensRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Error(errMsg, logger.Err(err))

		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})

		return
	}

	accessToken, refreshToken, err := h.service.GenerateTokens(c.Request.Context(), user.User{
		ID:    req.UserID,
		Email: req.Email,
		IP:    c.ClientIP(),
	})
	if err != nil {
		log.Error(errMsg, logger.Err(err))

		c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})

		return
	}

	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:  string(accessToken),
		RefreshToken: string(refreshToken),
	})
}
