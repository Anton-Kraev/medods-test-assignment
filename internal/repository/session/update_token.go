package session

import (
	"context"
	"fmt"
)

func (r Repository) UpdateToken(ctx context.Context, userID string, refreshToken string) error {
	const op = "Repository.UpdateToken"

	tx, err := r.db.BeginTx(ctx, txWriteOptions)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback(ctx)

	const query = `
		UPDATE sessions SET refresh_token_hash = $1 
		WHERE id = $2
	`

	if _, err = tx.Exec(
		ctx,
		query,
		refreshToken,
		userID,
	); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
