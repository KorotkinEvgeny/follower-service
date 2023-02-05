package postgres

var (
	insertFollowRecordQuery = `INSERT INTO follows (follower_id, followee_id) VALUES (:follower_id, :followee_id) RETURNING id, follower_id, followee_id;`
	insertUserQuery         = `INSERT INTO users (nickname) VALUES (:nickname) RETURNING id, created_date, nickname`
	getUserQuery            = `SELECT id, created_date, nickname FROM users WHERE id = :user_id`
	getFollowersQuery       = `SELECT id, follower_id, followee_id FROM follows WHERE followee_id = :user_id`
	getFolloweeQuery        = `SELECT id, follower_id, followee_id FROM follows WHERE follower_id = :user_id`
)
