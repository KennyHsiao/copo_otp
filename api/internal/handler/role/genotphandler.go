package role

import (
	"net/http"

	"github.com/copo888/copo_otp/api/internal/logic/role"
	"github.com/copo888/copo_otp/api/internal/svc"
	"github.com/copo888/copo_otp/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GenOtpHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OtpGenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewGenOtpLogic(r.Context(), ctx)
		resp, err := l.GenOtp(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
