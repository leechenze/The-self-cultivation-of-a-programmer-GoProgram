package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"monolithic/internal/logic"
	"monolithic/internal/svc"
	"monolithic/internal/types"
)

func MonolithicHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMonolithicLogic(r.Context(), svcCtx)
		resp, err := l.Monolithic(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
