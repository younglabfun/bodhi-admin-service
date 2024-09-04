package menu

import (
	"bodhiadmin/common/responsex"
	"net/http"

	"bodhiadmin/app/admin/api/internal/logic/menu"
	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetMenuListByPidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListMenuReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewGetMenuListByPidLogic(r.Context(), svcCtx)
		resp, err := l.GetMenuListByPid(&req)
		responsex.HttpResult(r, w, resp, err)
	}
}