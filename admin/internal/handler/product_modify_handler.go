package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerofp/admin/internal/logic"
	"zerofp/admin/internal/svc"
	"zerofp/admin/internal/types"
)

func ProductModifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductModifyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewProductModifyLogic(r.Context(), svcCtx)
		resp, err := l.ProductModify(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
