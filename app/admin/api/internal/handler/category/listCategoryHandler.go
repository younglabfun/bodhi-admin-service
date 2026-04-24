package category

import (
	"bodhiadmin/common/responsex"
	"net/http"

	"bodhiadmin/app/admin/api/internal/logic/category"
	"bodhiadmin/app/admin/api/internal/svc"
)

func ListCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := category.NewListCategoryLogic(r.Context(), svcCtx)
		resp, err := l.ListCategory()
		responsex.HttpResult(r, w, resp, err)
	}
}
