package tern

import (
	"context"
	"errors"
	"fmt"
	"github.com/madsnot/event-board-api/internal/config"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/tern/v2/migrate"
)

const timeFormat = "2006-01-02 15:04:05"

func RunMigrations(ctx context.Context, cfg config.Config) error {
	connConfig, err := pgx.ParseConfig(cfg.PostgresCfg.DSN)
	if err != nil {
		return err
	}

	conn, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	migrator, err := migrate.NewMigrator(ctx, conn, cfg.MigrationsCfg.VersionTable)
	if err != nil {
		return err
	}

	if err := migrator.LoadMigrations(os.DirFS(cfg.MigrationsCfg.Path)); err != nil {
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

	destination := cfg.MigrationsCfg.DestinationVersion
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
