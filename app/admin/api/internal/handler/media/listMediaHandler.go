package media

import (
	"bodhiadmin/common/responsex"
	"net/http"

	"bodhiadmin/app/admin/api/internal/logic/media"
	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListMediaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := media.NewListMediaLogic(r.Context(), svcCtx)
		resp, err := l.ListMedia(&req)
		responsex.HttpResult(r, w, resp, err)
	}
}
