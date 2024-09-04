package node

import (
	"bodhiadmin/common/responsex"
	"net/http"

	"bodhiadmin/app/admin/api/internal/logic/node"
	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BatchMoveGroupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MoveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := node.NewBatchMoveGroupLogic(r.Context(), svcCtx)
		resp, err := l.BatchMoveGroup(&req)
		responsex.HttpResult(r, w, resp, err)
	}
}
