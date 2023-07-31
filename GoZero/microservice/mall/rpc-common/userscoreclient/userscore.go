// Code generated by goctl. DO NOT EDIT.
// Source: userscore.proto

package userscoreclient

import (
	"context"

	"rpc-common/types/userscore"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IdRequest         = userscore.IdRequest
	UserScoreRequest  = userscore.UserScoreRequest
	UserScoreResponse = userscore.UserScoreResponse

	UserScore interface {
		SaveUserScore(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error)
		SaveUserScoreCallback(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error)
	}

	defaultUserScore struct {
		cli zrpc.Client
	}
)

func NewUserScore(cli zrpc.Client) UserScore {
	return &defaultUserScore{
		cli: cli,
	}
}

func (m *defaultUserScore) SaveUserScore(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error) {
	client := userscore.NewUserScoreClient(m.cli.Conn())
	return client.SaveUserScore(ctx, in, opts...)
}

func (m *defaultUserScore) SaveUserScoreCallback(ctx context.Context, in *UserScoreRequest, opts ...grpc.CallOption) (*UserScoreResponse, error) {
	client := userscore.NewUserScoreClient(m.cli.Conn())
	return client.SaveUserScoreCallback(ctx, in, opts...)
}
