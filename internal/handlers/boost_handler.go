package handlers

import (
	pb "boost/im_repo_boost"
	"context"
)

func (h *Handler) BoostPost(ctx context.Context, rq *pb.BoostRequest) (*pb.BoostResponse, error) {
	return h.repo.BoostPost(ctx, rq)
}

func (h *Handler) GetPosts(ctx context.Context, rq *pb.PageLimitRequest) (*pb.PostsResponse, error) {
	offset := (rq.Page - 1) * rq.Limit
	return h.repo.GetPosts(ctx, offset, rq.Limit)
}
