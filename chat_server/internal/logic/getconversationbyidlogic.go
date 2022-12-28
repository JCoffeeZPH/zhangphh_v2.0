package logic

import (
	"context"

	"zhangphh/chat_server/internal/svc"
	"zhangphh/chat_server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConversationByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConversationByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConversationByIdLogic {
	return &GetConversationByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConversationByIdLogic) GetConversationById(req *types.GetConvRequest) (resp *types.GetConvResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
