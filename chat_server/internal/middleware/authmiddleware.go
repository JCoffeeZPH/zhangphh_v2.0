package middleware

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strings"
	"time"
	"zhangphh/lib/client/userrpcservice"
)

type AuthMiddleware struct {
	UserCheckClient userrpcservice.UserRPCService
}

func NewAuthMiddleware(c userrpcservice.UserRPCService) *AuthMiddleware {
	return &AuthMiddleware{
		UserCheckClient: c,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logx.Infof("start authMiddleware")
		token := r.Header.Get("Authorization")
		splits := strings.Split(token, " ")
		if len(splits) == 2 && strings.ToLower(splits[0]) == "bearer" {
			ctx, _ := context.WithTimeout(r.Context(), 10*time.Second)
			userInfo, err := m.UserCheckClient.GetUserInfoByToken(ctx, &userrpcservice.GetUserByTokenRequest{Token: splits[1]})
			if err != nil {
				panic(err)
			}
			if userInfo != nil && userInfo.UserInfo.UserId <= 0 {
				logx.Errorf("user not exists!")
				panic("auth failed!")
			}
			logx.Infof("end authMiddleware")
			next(w, r)
		} else {
			panic("token is illegal")
		}
	}
}
