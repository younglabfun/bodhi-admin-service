package account

import (
	"bodhiadmin/common/responsex"
	"net/http"

	"bodhiadmin/app/admin/api/internal/logic/account"
	"bodhiadmin/app/admin/api/internal/svc"
)

func GetAccountHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := account.NewGetAccountLogic(r.Context(), svcCtx)
		resp, err := l.GetAccount()
		responsex.HttpResult(r, w, resp, err)
	}
}
