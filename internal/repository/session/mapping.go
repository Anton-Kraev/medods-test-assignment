package session

import "github.com/Anton-Kraev/medods-test-assignment/internal/models/user"

type session struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	IP           string `json:"ip_address"`
	RefreshToken string `json:"refresh_token_hash"`
}

func (s session) toUser() user.User {
	return user.User{
		ID:    s.ID,
		Email: s.Email,
		IP:    s.IP,
	}
}
