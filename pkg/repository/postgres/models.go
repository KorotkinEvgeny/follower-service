package postgres

type userDB struct {
	ID          int    `db:"id"`
	CreatedDate string `db:"created_date"`
	Nickname    string `db:"nickname"`
}

type followDB struct {
	ID         int `db:"id"`
	FollowerID int `db:"follower_id"`
	FolloweeID int `db:"followee_id"`
}
