package session

import (
	"context"
	"fmt"

	"github.com/georgysavva/scany/v2/pgxscan"

	"github.com/Anton-Kraev/medods-test-assignment/internal/models/errs"
	"github.com/Anton-Kraev/medods-test-assignment/internal/models/user"
)

func (r Repository) Get(ctx context.Context, userID string) (user.User, string, error) {
	const op = "Repository.Get"

	tx, err := r.db.BeginTx(ctx, txReadOptions)
	if err != nil {
		return user.User{}, "", fmt.Errorf("%s: %w", op, err)
	}
	defer tx.Rollback(ctx)

	const query = `SELECT * FROM sessions WHERE id = $1`

	var sessions []session
	if err = pgxscan.Select(ctx, tx, &sessions, query, userID); err != nil {
		return user.User{}, "", fmt.Errorf("%s: %w", op, err)
	}

	if len(sessions) == 0 {
		return user.User{}, "", fmt.Errorf("%s: %w", op, errs.ErrSessionNotFound)
	}

	if err = tx.Commit(ctx); err != nil {
		return user.User{}, "", fmt.Errorf("%s: %w", op, err)
	}

	return sessions[0].toUser(), sessions[0].RefreshToken, nil
}
