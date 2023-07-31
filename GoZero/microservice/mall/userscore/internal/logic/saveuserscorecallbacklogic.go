package logic

import (
	"context"

	"rpc-common/types/userscore"
	"userscore/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveUserScoreCallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveUserScoreCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveUserScoreCallbackLogic {
	return &SaveUserScoreCallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveUserScoreCallbackLogic) SaveUserScoreCallback(in *userscore.UserScoreRequest) (*userscore.UserScoreResponse, error) {
	// todo: add your logic here and delete this line
	println("userscore SaveUserScoreCallback is Capture Execution")
	return &userscore.UserScoreResponse{}, nil
}
