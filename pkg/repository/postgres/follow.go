package postgres

import (
	"context"
	"github.com/follower-service/pkg/dto"
	"github.com/jmoiron/sqlx"
)

type FollowRepository struct {
	db *sqlx.DB
}

func (f *FollowRepository) Store(ctx context.Context, request dto.Follow) (dto.Follow, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FollowRepository) ListFollowers(ctx context.Context, userID string) ([]*dto.Follow, error) {
	//TODO implement me
	panic("implement me")
}

func (f *FollowRepository) ListFollowee(ctx context.Context, userID string) ([]*dto.Follow, error) {
	//TODO implement me
	panic("implement me")
}

func NewFollowRepository(db *sqlx.DB) *FollowRepository {
	return &FollowRepository{
		db: db,
	}
}
