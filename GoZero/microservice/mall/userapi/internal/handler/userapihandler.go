package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"userapi/internal/logic"
	"userapi/internal/svc"
	"userapi/internal/types"
)

type UserHandler struct {
	svcCtx *svc.ServiceContext
}

func NewUserapiHandler(svcCtx *svc.ServiceContext) *UserHandler {
	return &UserHandler{
		svcCtx,
	}
}

func (h UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req types.Request
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	l := logic.NewUserapiLogic(r.Context(), h.svcCtx)
	resp, err := l.Userapi(&req)
	if err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
	} else {
		httpx.OkJsonCtx(r.Context(), w, resp)
	}
}
