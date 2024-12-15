package auth

import (
	"context"

	"github.com/Anton-Kraev/medods-test-assignment/internal/models/auth"
	"github.com/Anton-Kraev/medods-test-assignment/internal/models/user"
)

type (
	sessionRepository interface {
		Get(ctx context.Context, userID string) (user.User, auth.RefreshToken, error)
		Upsert(ctx context.Context, user user.User, refreshToken auth.RefreshToken) error
		UpdateToken(ctx context.Context, refreshToken auth.RefreshToken) error
	}

	emailClient interface {
		SendWarning(email string) error
	}

	Service struct {
		repository  sessionRepository
		emailClient emailClient
	}
)

func NewService(repository sessionRepository, emailClient emailClient) Service {
	return Service{
		repository:  repository,
		emailClient: emailClient,
	}
}

func (s Service) GenerateTokens(
	ctx context.Context, user user.User,
) (auth.AccessToken, auth.RefreshToken, error) {
	panic("not implemented")
}

func (s Service) RefreshTokens(
	ctx context.Context, userID string, userIP string, token auth.RefreshToken,
) (auth.AccessToken, auth.RefreshToken, error) {
	panic("not implemented")
}
