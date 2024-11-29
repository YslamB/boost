package handlers

import (
	pb "boost/im_repo_boost"
	"context"
	"fmt"
)

func (h *Handler) BoostPost(ctx context.Context, rq *pb.BoostRequest) (*pb.BoostResponse, error) {
	fmt.Println(rq.PostId)
	return h.repo.BoostPost(ctx, rq)
}
