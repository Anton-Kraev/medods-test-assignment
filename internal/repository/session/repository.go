package session

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Anton-Kraev/medods-test-assignment/internal/models/auth"
	"github.com/Anton-Kraev/medods-test-assignment/internal/models/user"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return Repository{db: db}
}

func (r Repository) Get(ctx context.Context, userID string) (user.User, auth.RefreshToken, error) {
	panic("implement me")
}

func (r Repository) Upsert(ctx context.Context, user user.User, refreshToken auth.RefreshToken) error {
	panic("implement me")
}

func (r Repository) UpdateToken(ctx context.Context, refreshToken auth.RefreshToken) error {
	panic("implement me")
}
