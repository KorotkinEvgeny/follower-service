package service

import (
	"context"
	"github.com/follower-service/pkg/dto"
)

type UserProcessor interface {
	Create(ctx context.Context, user dto.User) (*dto.User, error)
	ListUsers(ctx context.Context) ([]*dto.User, error)
	GetUsers(ctx context.Context) ([]*dto.User, error)
	GetUserInfo(ctx context.Context) ([]*dto.User, error)
}

type FollowProcessor interface {
	CreateFollow(ctx context.Context) (dto.Follow, error)
	RetrieveFollowee(ctx context.Context) ([]dto.Follow, error)
	RetrieveFollowers(ctx context.Context) ([]dto.Follow, error)
	Unfollow(ctx context.Context, followId string) error
}
