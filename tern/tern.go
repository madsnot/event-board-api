package tern

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/tern/v2/migrate"
)

const timeFormat = "2006-01-02 15:04:05"

func RunMigrations(ctx context.Context, pgDSN, migrationsPath, tableVersion, destination string) error {
	conn, err := newPgConn(ctx, pgDSN)
	if err != nil {
		return err
	}

	defer conn.Close(ctx)

	migrator, err := migrate.NewMigrator(ctx, conn, tableVersion)
	if err != nil {
		return err
	}

	if err := migrator.LoadMigrations(os.DirFS(migrationsPath)); err != nil {
		return err
	}

	if len(migrator.Migrations) == 0 {
		return errors.New("no migrations found")
	}

	migrator.OnStart = func(sequence int32, startAt, name, sql string) {
		startAt = time.Now().Format(timeFormat)
		fmt.Printf("%s %s sql: %s\n", startAt, name, sql)
	}

	currentVersion, err := migrator.GetCurrentVersion(ctx)
	if err != nil {
		return err
	}

	mustParseDestination := func(d string) int32 {
		var n int64

		n, err = strconv.ParseInt(d, 10, 32)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bad destination: %v\n", err)

			os.Exit(1)
		}

		return int32(n)
	}

	switch {
	case destination == "last":
		err = migrator.Migrate(ctx)
	case len(destination) >= 3 && destination[0:2] == "-+":
		err = migrator.MigrateTo(ctx, currentVersion-mustParseDestination(destination[2:]))
		if err == nil {
			err = migrator.MigrateTo(ctx, currentVersion)
		}
	case len(destination) >= 2 && destination[0] == '-':
		err = migrator.MigrateTo(ctx, currentVersion-mustParseDestination(destination[1:]))
	case len(destination) >= 2 && destination[0] == '+':
		err = migrator.MigrateTo(ctx, currentVersion+mustParseDestination(destination[1:]))
	default:
		err = migrator.MigrateTo(ctx, mustParseDestination(destination))
	}

	return err
}

func newPgConn(ctx context.Context, dsn string) (*pgx.Conn, error) {
	connConfig, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	conn, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
