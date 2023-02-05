package postgres

type Query string

type Queryx struct {
	Query  Query
	Params []interface{}
}

var (
	insertFollowRecordQuery Query = `INSERT INTO followers (follower, followee) VALUES (:follower, :followee);`
	insertUserQuery         Query = `INSERT INTO users (id) VALUES `
	getUserQuery            Query = `SELECT * FROM users WHERE user_id = :user_id`
)
