package role

import (
	"net/http"

	"github.com/copo888/copo_otp/api/internal/logic/role"
	"github.com/copo888/copo_otp/api/internal/svc"
	"github.com/copo888/copo_otp/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ValidateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OtpVaildReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewValidateLogic(r.Context(), ctx)
		resp, err := l.Validate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
