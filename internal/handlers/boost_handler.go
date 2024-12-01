package handlers

import (
	pb "boost/im_repo_boost"
	"context"
)

func (h *Handler) BoostPost(ctx context.Context, rq *pb.BoostRequest) (*pb.BoostResponse, error) {
	return h.repo.BoostPost(ctx, rq)
}
