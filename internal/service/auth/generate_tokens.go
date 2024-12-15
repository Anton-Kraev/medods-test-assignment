package auth

import (
	"context"
	"fmt"

	"github.com/Anton-Kraev/medods-test-assignment/internal/models/auth"
	"github.com/Anton-Kraev/medods-test-assignment/internal/models/user"
)

func (s Service) GenerateTokens(
	ctx context.Context, user user.User,
) (auth.AccessToken, auth.RefreshToken, error) {
	const op = "Service.GenerateTokens"

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

	err = s.repository.Upsert(ctx, user, refreshTokenHash)
	if err != nil {
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	return auth.AccessToken(accessToken), auth.RefreshToken(refreshToken), nil
}
