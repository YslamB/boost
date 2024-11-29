package handlers

import (
	pb "boost/im_repo_boost"
	"boost/internal/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

func NewHandler(logger *logrus.Entry, pool *pgxpool.Pool) *Handler {
	return &Handler{}
}

type Handler struct {
	repo repositories.Repository
	pb.BoostServiceServer
}
