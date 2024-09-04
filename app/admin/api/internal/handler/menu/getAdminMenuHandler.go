package menu

import (
	"bodhiadmin/common/responsex"
	"net/http"

	"bodhiadmin/app/admin/api/internal/logic/menu"
	"bodhiadmin/app/admin/api/internal/svc"
)

func GetAdminMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewGetAdminMenuLogic(r.Context(), svcCtx)
		resp, err := l.GetAdminMenu()
		responsex.HttpResult(r, w, resp, err)
	}
}
