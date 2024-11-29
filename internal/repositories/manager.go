package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func NewRepository(logger *logrus.Entry, pool *pgxpool.Pool) Repository {

	return Repository{
		logger:   logger,
		psqlPool: pool,
	}
}

type Repository struct {
	logger   *logrus.Entry
	psqlPool *pgxpool.Pool
}
