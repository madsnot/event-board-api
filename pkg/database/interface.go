package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBInterface interface {
	Open(context.Context) error
	Close() error

	Ping(context.Context) error
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	Query(context.Context, string, ...any) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
	CopyFrom(context.Context, pgx.Identifier, []string, pgx.CopyFromSource) (int64, error)
	SendBatch(context.Context, *pgx.Batch) pgx.BatchResults

	BeginTx(context.Context) (pgx.Tx, error)
	BeginTxOpt(context.Context, pgx.TxOptions) (pgx.Tx, error)
}
