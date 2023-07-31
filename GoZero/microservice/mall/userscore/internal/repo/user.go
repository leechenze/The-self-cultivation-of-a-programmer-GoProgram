package repo

import (
	"context"
	"userscore/internal/model"
)

type UserScoreRepo interface {
	SaveUserScore(ctx context.Context, userscore *model.UserScore) error
}
