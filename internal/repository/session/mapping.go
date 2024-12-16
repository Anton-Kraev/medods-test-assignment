package session

import "github.com/Anton-Kraev/medods-test-assignment/internal/models/user"

type session struct {
	ID           string `db:"id"`
	Email        string `db:"email"`
	IP           string `db:"ip_address"`
	RefreshToken string `db:"refresh_token_hash"`
}

func (s session) toUser() user.User {
	return user.User{
		ID:    s.ID,
		Email: s.Email,
		IP:    s.IP,
	}
}
