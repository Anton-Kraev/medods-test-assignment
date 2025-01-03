package auth

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Anton-Kraev/medods-test-assignment/internal/models/auth"
	"github.com/Anton-Kraev/medods-test-assignment/internal/models/errs"
	"github.com/Anton-Kraev/medods-test-assignment/pkg/logger"
)

type refreshTokensRequest struct {
	UserID       string `json:"user_id" binding:"required,uuid"`
	RefreshToken string `json:"refresh_token" binding:"required,base64"`
}

func (h Handler) RefreshTokens(c *gin.Context) {
	const errMsg = "failed to refresh tokens"

	log := h.log.With(slog.String("endpoint", "Handler.RefreshTokens"))

	var req refreshTokensRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Error(errMsg, logger.Err(err))

		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})

		return
	}

	accessToken, refreshToken, err := h.service.RefreshTokens(
		c.Request.Context(),
		req.UserID,
		c.ClientIP(),
		auth.RefreshToken(req.RefreshToken),
	)
	if err != nil {
		log.Error(errMsg, logger.Err(err))

		switch {
		case errors.Is(err, errs.ErrInvalidRefreshToken):
			c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		case errors.Is(err, errs.ErrSessionNotFound):
			c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
		}

		return
	}

	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:  string(accessToken),
		RefreshToken: string(refreshToken),
	})
}
