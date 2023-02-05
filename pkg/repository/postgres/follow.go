package postgres

import (
	"context"
	"github.com/follower-service/pkg/dto"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type FollowRepository struct {
	db *sqlx.DB
}

func (f *FollowRepository) Unfollow(ctx context.Context, unfollow dto.Unfollow) error {
	//TODO implement me
	panic("implement me")
}

func (f *FollowRepository) Store(ctx context.Context, request dto.Follow) (*dto.Follow, error) {
	args := map[string]interface{}{
		"follower_id": request.Follower,
		"followee_id": request.Followee,
	}
	log.Infof("follow store query preparation %v", args)
	rows, err := f.db.NamedQuery(insertFollowRecordQuery, args)
	if err != nil {
		log.Errorf("follow store query error %s", err.Error())
		return nil, err
	}
	followDBResp := followDB{}
	for rows.Next() {
		err = rows.Scan(&followDBResp.ID, &followDBResp.FollowerID, &followDBResp.FolloweeID)
		if err != nil {
			log.Errorf("follow store scan error %s", err.Error())
		}
	}

	return &dto.Follow{
		ID:       followDBResp.ID,
		Follower: followDBResp.FollowerID,
		Followee: followDBResp.FolloweeID,
	}, nil
}

func (f *FollowRepository) ListFollowers(ctx context.Context, userID int) ([]*dto.Follow, error) {
	args := map[string]interface{}{
		"user_id": userID,
	}
	rows, err := f.db.NamedQuery(getFollowersQuery, args)
	if err != nil {
		return nil, err
	}

	follows := make([]*dto.Follow, 0, 0)
	for rows.Next() {
		follow := dto.Follow{}
		err = rows.Scan(&follow.ID, &follow.Follower, &follow.Followee)
		if err != nil {
			log.Errorf("follow list scan error %s", err.Error())
			continue
		}
		follows = append(follows, &follow)
	}
	return follows, nil
}

func (f *FollowRepository) ListFollowee(ctx context.Context, userID int) ([]*dto.Follow, error) {
	args := map[string]interface{}{
		"user_id": userID,
	}
	rows, err := f.db.NamedQuery(getFolloweeQuery, args)
	if err != nil {
		return nil, err
	}

	follows := make([]*dto.Follow, 0, 0)
	for rows.Next() {
		follow := dto.Follow{}
		err = rows.Scan(&follow.ID, &follow.Follower, &follow.Followee)
		if err != nil {
			log.Errorf("followee list scan error %s", err.Error())
			continue
		}
		follows = append(follows, &follow)
	}
	return follows, nil
}

func NewFollowRepository(db *sqlx.DB) *FollowRepository {
	return &FollowRepository{
		db: db,
	}
}
