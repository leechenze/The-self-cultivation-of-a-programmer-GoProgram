package logic

import (
	"context"

	"monolithic/internal/svc"
	"monolithic/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MonolithicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMonolithicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MonolithicLogic {
	return &MonolithicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MonolithicLogic) Monolithic(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = &types.Response{
		Message: "Hello go-zero",
	}
	return
}
