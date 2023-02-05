package dto

type Follow struct {
	ID       int `json:"id"`
	Follower int `json:"follower"`
	Followee int `json:"followee"`
}

type User struct {
	ID          int    `json:"user_id"`
	Nickname    string `json:"nickname"`
	CreatedDate string `json:"created_date"`
}

type Unfollow struct {
	UserID   int `json:"user_id"`
	FollowID int `json:"follow_id"`
}
