package repository

import (
	"context"
	"github.com/follower-service/pkg/dto"
)

type FollowWriter interface {
	Store(ctx context.Context, request dto.Follow) (dto.Follow, error)
}

type FollowReader interface {
	ListFollowers(ctx context.Context, userID string) ([]*dto.Follow, error)
	ListFollowee(ctx context.Context, userID string) ([]*dto.Follow, error)
}

type FollowReaderWriter interface {
	FollowWriter
	FollowReader
}

type UserWriter interface {
	Store(ctx context.Context, user dto.User) (dto.User, error)
}

type UserReader interface {
	RetrieveUser(ctx context.Context, userID string) (dto.User, error)
	RetrieveUsers(ctx context.Context) ([]*dto.User, error)
}

type UserReaderWriter interface {
	UserReader
	UserWriter
}
