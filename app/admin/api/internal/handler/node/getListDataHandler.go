package node

import (
	"bodhiadmin/common/responsex"
	"net/http"

	"bodhiadmin/app/admin/api/internal/logic/node"
	"bodhiadmin/app/admin/api/internal/svc"
)

func GetListDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := node.NewGetListDataLogic(r.Context(), svcCtx)
		resp, err := l.GetListData()
		responsex.HttpResult(r, w, resp, err)
	}
}
