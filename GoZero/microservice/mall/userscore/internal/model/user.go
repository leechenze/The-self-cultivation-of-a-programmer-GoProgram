package model

type UserScore struct {
	Id     int64 `db:"id"`
	UserId int64 `db:"user_id"`
	Score  int   `db:"score"`
}

func (UserScore) TableName() string {
	return "user_score"
}
