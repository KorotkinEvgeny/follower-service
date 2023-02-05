package handler

import (
	"encoding/json"
	"github.com/follower-service/pkg/dto"
	"github.com/follower-service/pkg/service"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type UserHandler struct {
	service   service.UserProcessor
	validator *validator.Validate
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Error %s", err)
		responseBadRequestRender("incorrect request", w, r)
		return
	}

	model := new(userCreateRequest)
	err = json.Unmarshal(body, model)
	if err != nil {
		log.Errorf("Error %s", err)
		responseBadRequestRender("incorrect body", w, r)
		return
	}

	err = u.validator.Struct(model)
	if err != nil {
		responseUnprocessableEntityRender(err.Error(), w, r)
		return
	}

	user := dto.User{
		Nickname: model.Nickname,
	}
	createdUser, err := u.service.Create(ctx, user)
	if err != nil {
		responseInternalErrorRender(w, r)
		return
	}

	responseOKRender(createdUser, w, r)
	return
}

func (u *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := u.service.GetUsers(ctx)
	if err != nil {
		responseInternalErrorRender(w, r)
		return
	}
	responseOKRender(users, w, r)
	return
}

func (u *UserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	users, err := u.service.GetUsers(ctx)
	if err != nil {
		responseInternalErrorRender(w, r)
		return
	}
	responseOKRender(users, w, r)
	return
}

func NewUserHandler(userService service.UserProcessor, validator *validator.Validate) *UserHandler {
	return &UserHandler{
		service:   userService,
		validator: validator,
	}
}
