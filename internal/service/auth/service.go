package auth

import (
	"context"

	"github.com/Anton-Kraev/medods-test-assignment/internal/models/user"
)

type (
	sessionRepository interface {
		Get(ctx context.Context, userID string) (user.User, string, error)
		Upsert(ctx context.Context, user user.User, refreshToken string) error
		UpdateToken(ctx context.Context, userID string, refreshToken string) error
	}

	emailClient interface {
		SendWarning(ctx context.Context, email string) error
	}

	tokenManager interface {
		GenerateAccessToken(claims map[string]any) (string, error)
		GenerateRefreshToken() (string, string, error)
		CompareRefreshTokens(token string, tokenHash string) bool
	}

	Service struct {
		repository  sessionRepository
		emailClient emailClient
		tokens      tokenManager
	}
)

func NewService(repository sessionRepository, emailClient emailClient, tokens tokenManager) Service {
	return Service{
		repository:  repository,
		emailClient: emailClient,
		tokens:      tokens,
	}
}
