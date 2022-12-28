package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zhangphh/chat_server/internal/config"
	"zhangphh/lib/client/userrpcservice"
)

type ServiceContext struct {
	Config          config.Config
	UserCheckClient userrpcservice.UserRPCService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		UserCheckClient: userrpcservice.NewUserRPCService(zrpc.MustNewClient(c.UserCheckConf)),
	}
}
