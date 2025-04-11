package books

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"lowcode/internal/errcode"
	"lowcode/internal/logic/books"
	"lowcode/internal/svc"
	"lowcode/internal/types"
	"lowcode/internal/validate"
)

func DeleteBookHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		if err := validate.Validate(req); err != nil {
			httpx.OkJson(w, errcode.ValidateErr(err))
			return
		}

		l := books.NewDeleteBookLogic(r.Context(), svcCtx)
		resp, err := l.DeleteBook(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, errcode.OK(resp))
		}
	}
}
