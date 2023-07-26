package repo

import (
	"context"
	"user/internal/model"
)

type UserRepo interface {
	Save(ctx context.Context, user *model.User) error
}
