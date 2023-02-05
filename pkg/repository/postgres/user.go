package postgres

import (
	"context"
	"github.com/follower-service/pkg/dto"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func (u *UserRepository) RetrieveUser(ctx context.Context, userID string) (dto.User, error) {
	//res, err := u.db.NamedExec(getUserQuery)
	//dto.User{
	//	Nickname:    "",
	//	CreatedDate: "",
	//}
	panic("implement me")
}

func (u *UserRepository) RetrieveUsers(ctx context.Context) ([]*dto.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Store(ctx context.Context, user dto.User) (dto.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}
