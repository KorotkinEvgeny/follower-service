package service

import (
	"context"
	"github.com/follower-service/pkg/dto"
)

type UserProcessor interface {
	Create(ctx context.Context, user dto.User) (*dto.User, error)
	ListUsers(ctx context.Context) ([]*dto.User, error)
	GetUsers(ctx context.Context) ([]*dto.User, error)
	GetUserInfo(ctx context.Context, userID int) (*dto.User, error)
}

type FollowProcessor interface {
	CreateFollow(ctx context.Context, follow dto.Follow) (*dto.Follow, error)
	RetrieveFollowee(ctx context.Context, userID int) ([]*dto.Follow, error)
	RetrieveFollowers(ctx context.Context, userID int) ([]*dto.Follow, error)
	Unfollow(ctx context.Context, unfollow dto.Unfollow) error
}
