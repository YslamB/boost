package repositories

import (
	pb "boost/im_repo_boost"
	"boost/internal/queries"
	"context"
	"fmt"
)

func (repo *Repository) BoostPost(ctx context.Context, rq *pb.BoostRequest) (*pb.BoostResponse, error) {
	fmt.Println(rq.PostId)
	_, err := repo.psqlPool.Exec(ctx, queries.InsertBoostPost, rq.PostId)

	if err != nil {
		return &pb.BoostResponse{Message: err.Error(), Success: false}, err
	}
	return &pb.BoostResponse{Message: "succesfully created", Success: true}, nil
}

func (r *Repository) Test(ctx context.Context) (bool, error) {
	return true, nil
}
