package version

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"lowcode/internal/errcode"
	"lowcode/internal/logic/version"
	"lowcode/internal/svc"
)

func VersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := version.NewVersionLogic(r.Context(), svcCtx)
		resp, err := l.Version()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, errcode.OK(resp))
		}
	}
}
