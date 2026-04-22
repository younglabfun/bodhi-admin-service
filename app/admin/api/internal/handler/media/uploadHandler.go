package media

import (
	"bodhiadmin/common/responsex"
	"net/http"

	"bodhiadmin/app/admin/api/internal/logic/media"
	"bodhiadmin/app/admin/api/internal/svc"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := media.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(r)
		responsex.HttpResult(r, w, resp, err)
	}
}
