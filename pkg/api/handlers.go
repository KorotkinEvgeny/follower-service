package api

import "net/http"

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	GetUserInfo(w http.ResponseWriter, r *http.Request)
}

type FollowerHandler interface {
	CreateFollow(w http.ResponseWriter, r *http.Request)
	Unfollow(w http.ResponseWriter, r *http.Request)
	ListUserFollowers(w http.ResponseWriter, r *http.Request)
	ListUserFollowee(w http.ResponseWriter, r *http.Request)
}

type HealthHandler interface {
	IndexController(w http.ResponseWriter, r *http.Request)
}
