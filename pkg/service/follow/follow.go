package follow

import (
	"context"
	"github.com/follower-service/pkg/dto"
	"github.com/follower-service/pkg/repository"
)

type Service struct {
	repo repository.FollowReaderWriter
}

func (s *Service) CreateFollow(ctx context.Context) (dto.Follow, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) RetrieveFollowee(ctx context.Context) ([]dto.Follow, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) RetrieveFollowers(ctx context.Context) ([]dto.Follow, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Unfollow(ctx context.Context, followId string) error {
	//TODO implement me
	panic("implement me")
}

func NewFollowService(followRepo repository.FollowReaderWriter) *Service {
	return &Service{repo: followRepo}
}
