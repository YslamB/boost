package repositories

import (
	pb "boost/im_repo_boost"
	"context"
	"fmt"
)

func (repo *Repository) BoostPost(ctx context.Context, rq *pb.BoostRequest) (*pb.BoostResponse, error) {
	fmt.Println(rq.PostId)
	return &pb.BoostResponse{}, nil
}

func (r *Repository) Test(ctx context.Context) (bool, error) {
	return true, nil
}
