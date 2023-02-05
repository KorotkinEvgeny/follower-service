package handler

import (
	"github.com/follower-service/pkg/service"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type FollowHandler struct {
	service   service.FollowProcessor
	validator *validator.Validate
}

func (f *FollowHandler) CreateFollow(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (f *FollowHandler) ListUserFollowers(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (f *FollowHandler) ListUserFollowee(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewFollowHandler(followService service.FollowProcessor, validator *validator.Validate) *FollowHandler {
	return &FollowHandler{
		service:   followService,
		validator: validator,
	}
}
