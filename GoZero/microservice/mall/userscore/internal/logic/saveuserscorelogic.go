package logic

import (
	"context"
	"userscore/internal/model"

	"rpc-common/types/userscore"
	"userscore/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveUserScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveUserScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveUserScoreLogic {
	return &SaveUserScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveUserScoreLogic) SaveUserScore(in *userscore.UserScoreRequest) (*userscore.UserScoreResponse, error) {
	// todo: add your logic here and delete this line
	score := &model.UserScore{
		UserId: in.UserId,
		Score:  int(in.Score),
	}
	err := l.svcCtx.UserScoreRepo.SaveUserScore(context.Background(), score)
	if err != nil {
		return nil, err
	}
	return &userscore.UserScoreResponse{
		UserId: score.UserId,
		Score:  int32(score.Score),
	}, nil
}
