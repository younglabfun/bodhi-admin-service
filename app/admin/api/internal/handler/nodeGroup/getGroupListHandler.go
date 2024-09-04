package nodeGroup

import (
	"bodhiadmin/common/responsex"
	"net/http"

	"bodhiadmin/app/admin/api/internal/logic/nodeGroup"
	"bodhiadmin/app/admin/api/internal/svc"
)

func GetGroupListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := nodeGroup.NewGetGroupListLogic(r.Context(), svcCtx)
		resp, err := l.GetGroupList()
		responsex.HttpResult(r, w, resp, err)
	}
}
