package session

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return Repository{db: db}
}

var (
	txReadOptions  = pgx.TxOptions{AccessMode: pgx.ReadOnly, IsoLevel: pgx.ReadCommitted}
	txWriteOptions = pgx.TxOptions{AccessMode: pgx.ReadWrite, IsoLevel: pgx.RepeatableRead}
)
