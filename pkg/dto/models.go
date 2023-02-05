package dto

type Follow struct {
	Follower string `json:"follower"`
	Followee string `json:"followee"`
}

type User struct {
	Nickname    string `json:"nickname"`
	CreatedDate string `json:"created_date"`
}
