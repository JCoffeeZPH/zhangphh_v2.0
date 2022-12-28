package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zhangphh/chat_server/internal/logic"
	"zhangphh/chat_server/internal/svc"
	"zhangphh/chat_server/internal/types"
)

func SendMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendMsgRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSendMsgLogic(r.Context(), svcCtx)
		resp, err := l.SendMsg(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
