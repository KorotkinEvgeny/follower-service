package handler

type followRequest struct {
	Followee int `json:"followee_id"`
}

type userCreateRequest struct {
	Nickname string `json:"nickname" validate:"required"`
}

type unfollowRequest struct {
	FollowID int `json:"follow_id"`
}
