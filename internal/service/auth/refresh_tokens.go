package auth

import (
	"context"
	"fmt"

	"github.com/Anton-Kraev/medods-test-assignment/internal/models/auth"
	"github.com/Anton-Kraev/medods-test-assignment/internal/models/errs"
)

func (s Service) RefreshTokens(
	ctx context.Context, userID string, userIP string, token auth.RefreshToken,
) (auth.AccessToken, auth.RefreshToken, error) {
	const op = "Service.RefreshTokens"

	user, refreshToken, err := s.repository.Get(ctx, userID)
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	if !s.tokens.CompareRefreshTokens(string(token), refreshToken) {
		return "", "", fmt.Errorf("%s: %w", op, errs.ErrInvalidRefreshToken)
	}

	if userIP != user.IP {
		if err = s.emailClient.SendWarning(ctx, user.Email); err != nil {
			return "", "", fmt.Errorf("%s: %w", op, err)
		}
	}

	accessToken, err := s.tokens.GenerateAccessToken(map[string]any{
		"UserID": user.ID,
		"UserIP": user.IP,
	})
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	refreshToken, refreshTokenHash, err := s.tokens.GenerateRefreshToken()
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	err = s.repository.UpdateToken(ctx, userID, refreshTokenHash)
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	return auth.AccessToken(accessToken), auth.RefreshToken(refreshToken), nil
}
