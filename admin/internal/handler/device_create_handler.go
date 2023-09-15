package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerofp/admin/internal/logic"
	"zerofp/admin/internal/svc"
	"zerofp/admin/internal/types"
)

func DeviceCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeviceCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeviceCreateLogic(r.Context(), svcCtx)
		resp, err := l.DeviceCreate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
