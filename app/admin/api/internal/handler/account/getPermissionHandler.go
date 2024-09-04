package account

import (
	"bodhiadmin/common/responsex"
	"net/http"

	"bodhiadmin/app/admin/api/internal/logic/account"
	"bodhiadmin/app/admin/api/internal/svc"
)

func GetPermissionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := account.NewGetPermissionLogic(r.Context(), svcCtx)
		resp, err := l.GetPermission()
		responsex.HttpResult(r, w, resp, err)
	}
}
