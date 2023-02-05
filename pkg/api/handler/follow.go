package handler

import (
	"encoding/json"
	"github.com/follower-service/pkg/dto"
	"github.com/follower-service/pkg/service"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
)

const userHeaderIdentifierKey = "X-User-ID"

type FollowHandler struct {
	service     service.FollowProcessor
	userService service.UserProcessor
	validator   *validator.Validate
}

func (f *FollowHandler) Unfollow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIDHeader := r.Header.Get(userHeaderIdentifierKey)

	if userIDHeader == "" {
		responseForbiddenRender("Provide user identity", w, r)
		return
	}
	userID, err := strconv.Atoi(userIDHeader)
	if err != nil {
		responseForbiddenRender("Incorrect user identity", w, r)
		return
	}

	user, err := f.userService.GetUserInfo(ctx, userID)
	if err != nil {
		responseInternalErrorRender(w, r)
		return
	}

	if user == nil {
		responseNotFound(userIDHeader, w, r)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Error %s", err)
		responseBadRequestRender("incorrect request", w, r)
		return
	}

	model := new(unfollowRequest)
	err = json.Unmarshal(body, model)
	if err != nil {
		log.Errorf("Error %s", err)
		responseBadRequestRender("incorrect body", w, r)
		return
	}

	err = f.validator.Struct(model)
	if err != nil {
		responseUnprocessableEntityRender(err.Error(), w, r)
		return
	}

	unfollowModel := dto.Unfollow{
		UserID:   userID,
		FollowID: model.FollowID,
	}

	err = f.service.Unfollow(ctx, unfollowModel)
	if err != nil {
		responseInternalErrorRender(w, r)
		return
	}
	responseNoContent(w, r)

}

func (f *FollowHandler) CreateFollow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIDHeader := r.Header.Get(userHeaderIdentifierKey)

	if userIDHeader == "" {
		responseForbiddenRender("Provide user identity", w, r)
		return
	}
	userID, err := strconv.Atoi(userIDHeader)
	if err != nil {
		responseForbiddenRender("Incorrect user identity", w, r)
		return
	}

	user, err := f.userService.GetUserInfo(ctx, userID)
	if err != nil {
		responseInternalErrorRender(w, r)
		return
	}

	if user == nil {
		responseNotFound(userIDHeader, w, r)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Errorf("Error %s", err)
		responseBadRequestRender("incorrect request", w, r)
		return
	}

	model := new(followRequest)
	err = json.Unmarshal(body, model)
	if err != nil {
		log.Errorf("Error %s", err)
		responseBadRequestRender("incorrect body", w, r)
		return
	}

	err = f.validator.Struct(model)
	if err != nil {
		responseUnprocessableEntityRender(err.Error(), w, r)
		return
	}

	follow := dto.Follow{
		Follower: userID,
		Followee: model.Followee,
	}

	followR, err := f.service.CreateFollow(ctx, follow)
	if err != nil {
		responseInternalErrorRender(w, r)
		return
	}
	responseOKRender(followR, w, r)
}

func (f *FollowHandler) ListUserFollowers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIDHeader := r.Header.Get(userHeaderIdentifierKey)

	if userIDHeader == "" {
		responseForbiddenRender("Provide user identity", w, r)
		return
	}
	userID, err := strconv.Atoi(userIDHeader)
	if err != nil {
		responseForbiddenRender("Incorrect user identity", w, r)
		return
	}

	user, err := f.userService.GetUserInfo(ctx, userID)
	if err != nil {
		responseInternalErrorRender(w, r)
		return
	}

	if user == nil {
		responseNotFound(userIDHeader, w, r)
		return
	}

	followers, err := f.service.RetrieveFollowers(ctx, userID)
	if err != nil {
		responseInternalErrorRender(w, r)
		return
	}
	responseOKRender(followers, w, r)
}

func (f *FollowHandler) ListUserFollowee(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	//TODO move to the middleware
	userIDHeader := r.Header.Get(userHeaderIdentifierKey)

	if userIDHeader == "" {
		responseForbiddenRender("Provide user identity", w, r)
		return
	}
	userID, err := strconv.Atoi(userIDHeader)
	if err != nil {
		responseForbiddenRender("Incorrect user identity", w, r)
		return
	}

	user, err := f.userService.GetUserInfo(ctx, userID)
	if err != nil {
		responseInternalErrorRender(w, r)
		return
	}

	if user == nil {
		responseNotFound(userIDHeader, w, r)
		return
	}

	followee, err := f.service.RetrieveFollowee(ctx, userID)
	if err != nil {
		responseInternalErrorRender(w, r)
		return
	}
	responseOKRender(followee, w, r)
}

func NewFollowHandler(followService service.FollowProcessor, validator *validator.Validate, userService service.UserProcessor) *FollowHandler {
	return &FollowHandler{
		userService: userService,
		service:     followService,
		validator:   validator,
	}
}
