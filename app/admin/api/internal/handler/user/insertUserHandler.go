package user

import (
	"bodhiadmin/common/responsex"
	"net/http"

	"bodhiadmin/app/admin/api/internal/logic/user"
	"bodhiadmin/app/admin/api/internal/svc"
	"bodhiadmin/app/admin/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func InsertUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NewUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewInsertUserLogic(r.Context(), svcCtx)
		resp, err := l.InsertUser(&req)
		responsex.HttpResult(r, w, resp, err)
	}
}