package config

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	Database *pgx.Conn
	Context  context.Context
}

func (d *Database) Init() error {
	ctx := context.Background()

	db, err := pgx.Connect(ctx, "host=localhost port=5432 user=postgres password=password dbname=database sslmode=disable")

	if err != nil {
		return err
	}

	if err = db.Ping(ctx); err != nil {
		return err
	}

	d.Database = db
	d.Context = ctx

	return nil
}
