package role

import (
	"context"
	"github.com/copo888/copo_otp/helper/otpx"

	"github.com/copo888/copo_otp/api/internal/svc"
	"github.com/copo888/copo_otp/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GenOtpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenOtpLogic(ctx context.Context, svcCtx *svc.ServiceContext) GenOtpLogic {
	return GenOtpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenOtpLogic) GenOtp(req types.OtpGenReq) (resp *types.OtpGenReply, err error) {
	auth, err :=otpx.GenOtpKey(req.Issuer, req.Account)
	if err != nil {
		return &types.OtpGenReply{
			Code: "1",
			Message: "Fail",
		}, err
	}

	return &types.OtpGenReply{
		Code: "0",
		Message: "Success",
		Data: types.OtpData{
			Secret: auth.Code,
			Qrcode: auth.Path,
		},
	},nil
}
