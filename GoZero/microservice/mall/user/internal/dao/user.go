package dao

import (
	"context"
	"fmt"
	"user/database"
	"user/internal/model"
)

type UserDao struct {
	*database.DBConn
}

func NewUserDao(Conn *database.DBConn) *UserDao {
	return &UserDao{
		Conn,
	}
}

func (d *UserDao) Save(ctx context.Context, user *model.User) error {

	sqlStat := fmt.Sprintf("insert into %s (name,gender) values(?,?)", user.TableName())

	execRes, err := d.Conn.ExecCtx(ctx, sqlStat, user.Name, user.Gender)
	if err != nil {
		return err
	}
	id, err := execRes.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = id

	return nil
}
