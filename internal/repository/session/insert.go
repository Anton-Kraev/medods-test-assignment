package session

import (
	"context"
	"fmt"

	"github.com/Anton-Kraev/medods-test-assignment/internal/models/user"
)

func (r Repository) Insert(ctx context.Context, user user.User, refreshToken string) error {
	const op = "Repository.Insert"

	tx, err := r.db.BeginTx(ctx, txWriteOptions)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback(ctx)

	const query = `
		INSERT INTO sessions(id, email, ip_address, refresh_token_hash) 
		VALUES ($1, $2, $3, $4)
	`

	if _, err = tx.Exec(
		ctx,
		query,
		user.ID,
		user.Email,
		user.IP,
		refreshToken,
	); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
