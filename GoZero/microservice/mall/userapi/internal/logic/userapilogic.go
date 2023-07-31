package logic

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"rpc-common/types/user"
	"rpc-common/types/userscore"
	"strconv"
	"time"
	"userapi/internal/customError"
	"userapi/internal/svc"
	"userapi/internal/types"
)

type UserapiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 定义UserapiLogic的构造函数
func NewUserapiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserapiLogic {
	return &UserapiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 定义生成token的方法
func (l *UserapiLogic) getToken(secretKey string, iat, seconds, userId int64) (string, error) {
	// seconds：过期时间
	// iat：表示签发时间，即当前时间
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *UserapiLogic) Register(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	// 超时上下文，不用 context.Background 了
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	userResponse, err := l.svcCtx.UserRPC.SaveUser(ctx, &user.UserRequest{
		Name:   req.Name,
		Gender: req.Gender,
	})
	if err != nil {
		return nil, err
	}
	// 加积分
	userId, _ := strconv.ParseInt(userResponse.Id, 10, 64)
	println("=============================================")
	fmt.Printf("userId: %v \n", userId)
	println("=============================================")
	userScore, err := l.svcCtx.UserScoreRPC.SaveUserScore(context.Background(), &userscore.UserScoreRequest{
		UserId: userId,
		Score:  10,
	})
	fmt.Printf("registed add score %d \n", userScore.Score)
	return &types.Response{
		Message: "success",
		Data:    userResponse,
	}, nil
}

func (l *UserapiLogic) GetUser(t *types.IdRequest) (resp *types.Response, err error) {

	// 模拟一个自定义错误的测试
	if t.Id == "1" {
		return nil, customError.ParamsError
	}

	// 认证通过后从token中获取用户id，获取之后会放到ctx上下文中，直接从ctx中回去即可。
	userId := l.ctx.Value("userId")
	// 这里用go-zero的打印日志，信息比较全面。
	logx.Infof("获取到的token的userId是：%v \n", userId)
	userResponse, err := l.svcCtx.UserRPC.GetUser(context.Background(), &user.IdRequest{
		Id: t.Id,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.Response{
		Message: "success",
		Data:    userResponse,
	}
	return
}

func (l *UserapiLogic) Login(t *types.LoginRequest) (token string, err error) {
	logx.Infof("正在执行login方法")
	userId := 100
	auth := l.svcCtx.Config.Auth
	return l.getToken(auth.AccessSecret, time.Now().Unix(), auth.AccessExpire, int64(userId))
}
