package user

import (
	"context"
	"github.com/follower-service/pkg/dto"
	"github.com/follower-service/pkg/repository"
)

type Service struct {
	repo repository.UserReaderWriter
}

func (s *Service) Create(ctx context.Context, user dto.User) (*dto.User, error) {
	userStored, err := s.repo.Store(ctx, user)
	if err != nil {
		return nil, err
	}

	return userStored, nil
}

func (s *Service) ListUsers(ctx context.Context) ([]*dto.User, error) {
	users, err := s.repo.RetrieveUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) GetUsers(ctx context.Context) ([]*dto.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GetUserInfo(ctx context.Context, userID int) (*dto.User, error) {
	user, err := s.repo.RetrieveUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserService(userRepo repository.UserReaderWriter) *Service {
	return &Service{
		repo: userRepo,
	}
}
