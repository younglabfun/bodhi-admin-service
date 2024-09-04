package role

import (
	"bodhiadmin/common/responsex"
	"net/http"

	"bodhiadmin/app/admin/api/internal/logic/role"
	"bodhiadmin/app/admin/api/internal/svc"
)

func GetListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := role.NewGetListLogic(r.Context(), svcCtx)
		resp, err := l.GetList()
		responsex.HttpResult(r, w, resp, err)
	}
}
