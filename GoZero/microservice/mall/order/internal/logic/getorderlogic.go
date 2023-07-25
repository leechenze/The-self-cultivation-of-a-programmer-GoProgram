package logic

import (
	"context"
	"user/types/user"

	"order/internal/svc"
	"order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderRequest) (resp *types.OrderResponse, err error) {
	// todo: add your logic here and delete this line
	userId := l.GetOrderById(req.Id)
	// 根据用户Id去user服务获取用户信息
	getUserResponse, err := l.svcCtx.UserRPC.GetUser(context.Background(), &user.IdRequest{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.OrderResponse{
		Id:       req.Id,
		Name:     "hello order name",
		UserName: getUserResponse.Name,
	}, nil
}

func (l *GetOrderLogic) GetOrderById(id string) string {
	return "1"
}
