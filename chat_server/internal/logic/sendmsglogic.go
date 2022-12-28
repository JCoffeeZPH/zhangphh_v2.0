package logic

import (
	"context"
	"zhangphh/chat_server/internal/svc"
	"zhangphh/chat_server/internal/types"
	"zhangphh/lib/client/userrpcservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMsgLogic) SendMsg(req *types.SendMsgRequest) (*types.SendMsgResponse, error) {
	user, e := l.svcCtx.UserCheckClient.GetUserInfo(l.ctx, &userrpcservice.GetUserRequest{UserId: 1})
	if e != nil {
		return nil, e
	}
	logx.Infof("user: %+v", user)
	return &types.SendMsgResponse{
		Code: 200,
	}, nil
}
