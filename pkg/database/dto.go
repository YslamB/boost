package database

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx"
)

type Options struct {
	Host          string
	Port          string
	Database      string
	Username      string
	Password      string
	PgPoolMaxConn int
}

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	Close()
}
