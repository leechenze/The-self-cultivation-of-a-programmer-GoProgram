package dao

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/database"
	"user/internal/model"
)

var cacheUserIdPrefix = "cache:user:id"

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

func (d *UserDao) FindById(ctx context.Context, id int64) (user *model.User, err error) {
	user = &model.User{}
	sqlQuery := fmt.Sprintf("select * from %s where id = ?", user.TableName())
	userIdKey := fmt.Sprintf("%s:%d", cacheUserIdPrefix, id)
	err = d.ConnCache.QueryRowCtx(ctx, user, userIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		return conn.QueryRowCtx(ctx, v, sqlQuery, id)
	})
	return
}
