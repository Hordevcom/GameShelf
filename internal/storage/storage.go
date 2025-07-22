package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGDB struct {
	DB *pgxpool.Pool
}

func NewStorage() *PGDB {
	db, err := pgxpool.New(context.Background(), "postgres://postgres:1@localhost:5432/postgres")

	if err != nil {
		fmt.Println("Problem with connection to db: ", err)
		return nil
	}

	err = db.Ping(context.Background())
	if err != nil {
		fmt.Println("Problem with ping to db: ", err)
		return nil
	}

	return &PGDB{DB: db}
}
