package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewClient(ctx context.Context, options *Options) (*pgxpool.Pool, error) {

	log.Println("new client options")
	log.Println("🔔 Host: ", options.Host)
	log.Println("🔔 Port: ", options.Port)
	log.Println("🔔 Database: ", options.Database)
	log.Println("🔔 Username: ", options.Username)
	log.Println("🔔 Password: ", options.Password)
	log.Println("🔔 PgPoolMaxConn: ", options.PgPoolMaxConn)

	connPool, err := pgxpool.NewWithConfig(ctx, getConfig(*options))

	if err != nil {
		return nil, errors.New("🚫 error while creating connection to the database")
	}

	connection, err := connPool.Acquire(ctx)

	if err != nil {
		return nil, errors.New("🚫 error while acquiring connection from the database pool")
	}

	defer connection.Release()
	err = connection.Ping(ctx)

	if err != nil {
		return nil, errors.New("🚫 Could not ping database")
	}

	log.Println("✅ postgresql connected success")

	//return &Pool{connPool}, nil
	return connPool, nil
}

func getConfig(options Options) *pgxpool.Config {

	DatabaseUrl := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		options.Username,
		options.Password,
		options.Host,
		options.Port,
		options.Database)

	log.Println("🔔 database url: ", DatabaseUrl)
	dbConfig, err := pgxpool.ParseConfig(DatabaseUrl)

	if err != nil {
		log.Println("🚫 Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = int32(options.PgPoolMaxConn)
	dbConfig.MinConns = int32(0)
	dbConfig.MaxConnLifetime = time.Hour
	dbConfig.MaxConnIdleTime = time.Minute * 30
	dbConfig.HealthCheckPeriod = time.Minute
	dbConfig.ConnConfig.ConnectTimeout = time.Second * 5

	return dbConfig
}
