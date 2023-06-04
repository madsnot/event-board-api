package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Init(url string) (_ *pgxpool.Pool, err error) {
	dataBasePool, errMakePool := pgxpool.New(context.Background(), url)

	if errMakePool != nil {
		return nil, errMakePool
	}

	return dataBasePool, err
}
