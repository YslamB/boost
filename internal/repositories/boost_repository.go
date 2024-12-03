package repositories

import (
	pb "boost/im_repo_boost"
	"boost/internal/queries"
	"context"
	"fmt"
)

func (repo *Repository) BoostPost(ctx context.Context, rq *pb.BoostRequest) (*pb.BoostResponse, error) {
	_, err := repo.psqlPool.Exec(ctx, queries.InsertBoostPost, rq.PostId)

	if err != nil {
		return &pb.BoostResponse{Message: err.Error(), Success: false}, err
	}
	return &pb.BoostResponse{Message: "succesfully created", Success: true}, nil
}

func (repo *Repository) GetPosts(ctx context.Context, offset, limit int32) (*pb.PostsResponse, error) {
	rows, err := repo.psqlPool.Query(ctx, queries.GetPosts, limit, offset)
	if err != nil {
		return nil, err
	}
	var posts []*pb.Posts

	for rows.Next() {
		post := &pb.Posts{}

		err := rows.Scan(
			&post.Id,
			&post.Description,
			&post.Likes,
			&post.Paths,
			&post.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		posts = append(posts, post)
	}
	return &pb.PostsResponse{Status: true, Posts: posts}, nil
}

func (repo *Repository) GetBoostedPoosts(ctx context.Context, offset, limit int32) (*pb.PostsResponse, error) {
	rows, err := repo.psqlPool.Query(ctx, queries.GetPosts, offset, limit)
	if err != nil {
		return nil, err
	}
	var posts []*pb.Posts

	for rows.Next() {
		post := &pb.Posts{}

		err := rows.Scan(
			&post.Id,
			&post.Description,
			&post.Likes,
			&post.Paths,
			&post.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		posts = append(posts, post)
	}
	return &pb.PostsResponse{Status: true, Posts: posts}, nil
}

func (r *Repository) Test(ctx context.Context) (bool, error) {
	return true, nil
}
