package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerofp/admin/internal/logic"
	"zerofp/admin/internal/svc"
	"zerofp/admin/internal/types"
)

func DeviceModifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeviceModifyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeviceModifyLogic(r.Context(), svcCtx)
		resp, err := l.DeviceModify(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
