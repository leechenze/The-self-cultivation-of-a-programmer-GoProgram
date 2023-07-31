package dao

import (
	"context"
	"fmt"
	"userscore/database"
	"userscore/internal/model"
)

var cacheUserIdPrefix = "cache:user:id"

type UserScoreDao struct {
	*database.DBConn
}

func NewUserScoreDao(Conn *database.DBConn) *UserScoreDao {
	return &UserScoreDao{
		Conn,
	}
}

func (d *UserScoreDao) SaveUserScore(ctx context.Context, userscore *model.UserScore) error {

	sqlStat := fmt.Sprintf("insert into %s (user_id,score) values(?,?)", userscore.TableName())

	execRes, err := d.Conn.ExecCtx(ctx, sqlStat, userscore.UserId, userscore.Score)
	if err != nil {
		return err
	}
	id, err := execRes.LastInsertId()
	if err != nil {
		return err
	}
	userscore.Id = id

	return nil
}
