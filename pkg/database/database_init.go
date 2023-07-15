package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
	dsn  string
}

func NewDB(dsn string) *DB {
	return &DB{
		dsn: dsn,
	}
}

func (db *DB) Open(ctx context.Context) error {
	config, err := pgxpool.ParseConfig(db.dsn)
	if err != nil {
		return err
	}

	db.pool, err = pgxpool.NewWithConfig(ctx, config)

	return err
}

func (db *DB) Close() error {
	if db.pool != nil {
		db.pool.Close()
	}

	return nil
}

func (db *DB) Ping(ctx context.Context) error {
	return db.pool.Ping(ctx)
}

func (db *DB) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return db.pool.Exec(ctx, sql, args...)
}

func (db *DB) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return db.pool.Query(ctx, sql, args...)
}

func (db *DB) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return db.pool.QueryRow(ctx, sql, args...)
}

func (db *DB) CopyFrom(ctx context.Context, identifier pgx.Identifier, strings []string, source pgx.CopyFromSource) (int64, error) {
	return db.pool.CopyFrom(ctx, identifier, strings, source)
}

func (db *DB) SendBatch(ctx context.Context, batch *pgx.Batch) pgx.BatchResults {
	return db.pool.SendBatch(ctx, batch)
}

func (db *DB) BeginTx(ctx context.Context) (pgx.Tx, error) {
	return db.pool.Begin(ctx)
}

func (db *DB) BeginTxOpt(ctx context.Context, options pgx.TxOptions) (pgx.Tx, error) {
	return db.pool.BeginTx(ctx, options)
}
