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

type RunningConfig struct {
	DSN            string
	TimeFormat     string
	MigrationsPath string
	TableVersion   string
	Destination    string
}

func NewRunningConfig(DSN, migrationsPath, tableVersion, destination string) RunningConfig {
	return RunningConfig{
		DSN:            DSN,
		MigrationsPath: migrationsPath,
		TableVersion:   tableVersion,
		Destination:    destination,
		TimeFormat:     "2006-01-02 15:04:05",
	}
}

func (rc RunningConfig) RunMigrations(ctx context.Context) error {
	conn, err := newPgConn(ctx, rc.DSN)
	if err != nil {
		return err
	}

	defer conn.Close(ctx)

	migrator, err := migrate.NewMigrator(ctx, conn, rc.TableVersion)
	if err != nil {
		return err
	}

	if err := migrator.LoadMigrations(os.DirFS(rc.MigrationsPath)); err != nil {
		return err
	}

	if len(migrator.Migrations) == 0 {
		return errors.New("no migrations found")
	}

	migrator.OnStart = func(sequence int32, startAt, name, sql string) {
		startAt = time.Now().Format(rc.TimeFormat)
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
	case rc.Destination == "last":
		err = migrator.Migrate(ctx)
	case len(rc.Destination) >= 3 && rc.Destination[0:2] == "-+":
		err = migrator.MigrateTo(ctx, currentVersion-mustParseDestination(rc.Destination[2:]))
		if err == nil {
			err = migrator.MigrateTo(ctx, currentVersion)
		}
	case len(rc.Destination) >= 2 && rc.Destination[0] == '-':
		err = migrator.MigrateTo(ctx, currentVersion-mustParseDestination(rc.Destination[1:]))
	case len(rc.Destination) >= 2 && rc.Destination[0] == '+':
		err = migrator.MigrateTo(ctx, currentVersion+mustParseDestination(rc.Destination[1:]))
	default:
		err = migrator.MigrateTo(ctx, mustParseDestination(rc.Destination))
	}

	return err
}

func newPgConn(ctx context.Context, DSN string) (*pgx.Conn, error) {
	connConfig, err := pgx.ParseConfig(DSN)
	if err != nil {
		return nil, err
	}

	conn, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
