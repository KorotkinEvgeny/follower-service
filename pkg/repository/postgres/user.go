package postgres

import (
	"context"
	"github.com/follower-service/pkg/dto"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type UserRepository struct {
	db *sqlx.DB
}

func (u *UserRepository) RetrieveUser(ctx context.Context, userID int) (*dto.User, error) {
	args := map[string]interface{}{
		"user_id": userID,
	}

	rows, err := u.db.NamedQuery(getUserQuery, args)
	if err != nil {
		return nil, err
	}
	userDBResp := userDB{}
	for rows.Next() {
		err = rows.Scan(&userDBResp.ID, &userDBResp.CreatedDate, &userDBResp.Nickname)
		log.Infof("User insert result Rows  %v", userDBResp)
	}
	return &dto.User{
		ID:          userDBResp.ID,
		Nickname:    userDBResp.Nickname,
		CreatedDate: userDBResp.CreatedDate,
	}, err
}

func (u *UserRepository) RetrieveUsers(ctx context.Context) ([]*dto.User, error) {
	panic("implement me")
}

func (u *UserRepository) Store(ctx context.Context, user dto.User) (*dto.User, error) {
	args := map[string]interface{}{
		"nickname": user.Nickname,
	}
	rows, err := u.db.NamedQuery(insertUserQuery, args)
	if err != nil {
		return nil, err
	}
	userDBResp := userDB{}
	for rows.Next() {
		err = rows.Scan(&userDBResp.ID, &userDBResp.CreatedDate, &userDBResp.Nickname)
		log.Infof("User insert result Rows  %v", userDBResp)
	}
	return &dto.User{
		ID:          userDBResp.ID,
		Nickname:    userDBResp.Nickname,
		CreatedDate: userDBResp.CreatedDate,
	}, err
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}
