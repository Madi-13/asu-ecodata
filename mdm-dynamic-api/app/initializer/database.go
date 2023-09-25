package initializer

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func ConnectDB(databaseUrl *string) error {

	config, err := pgxpool.ParseConfig(*databaseUrl)
	if err != nil {
		return err
	}
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		log.Println("Connection: +1")
		return nil
	}
	config.MaxConns = 6
	config.MinConns = 2
	pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return err
	}
	_, err = pool.Exec(context.Background(), "SET TIME ZONE 'Asia/Almaty'")
	if err != nil {
		return err
	}
	_, err = pool.Exec(context.Background(), "SET search_path TO public")
	if err != nil {
		return err
	}

	return nil
}

func GetConnection() (*pgxpool.Conn, error) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func PoolClose() {
	pool.Close()
}
