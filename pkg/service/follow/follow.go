package follow

import (
	"context"
	"github.com/follower-service/pkg/dto"
	"github.com/follower-service/pkg/repository"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	repo repository.FollowReaderWriter
}

func (s *Service) CreateFollow(ctx context.Context, follow dto.Follow) (*dto.Follow, error) {
	followStored, err := s.repo.Store(ctx, follow)
	if err != nil {
		log.Errorf("Follow Service error %s", err.Error())
		return nil, err
	}
	return followStored, nil
}

func (s *Service) RetrieveFollowee(ctx context.Context, userID int) ([]*dto.Follow, error) {
	followee, err := s.repo.ListFollowee(ctx, userID)
	if err != nil {
		return nil, err
	}
	return followee, nil
}

func (s *Service) RetrieveFollowers(ctx context.Context, userID int) ([]*dto.Follow, error) {
	followers, err := s.repo.ListFollowers(ctx, userID)
	if err != nil {
		return nil, err
	}
	return followers, nil
}

func (s *Service) Unfollow(ctx context.Context, unfollow dto.Unfollow) error {
	err := s.repo.Unfollow(ctx, unfollow)
	if err != nil {
		return err
	}
	return nil
}

func NewFollowService(followRepo repository.FollowReaderWriter) *Service {
	return &Service{repo: followRepo}
}
