package logic

import (
	"context"
	"github.com/copo888/copo_otp/helper/otpx"
	"github.com/copo888/copo_otp/rpc/optclient"

	"github.com/copo888/copo_otp/rpc/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
)

type GenOtpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenOtpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenOtpLogic {
	return &GenOtpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenOtpLogic) GenOtp(in *optclient.OtpGenRequest) (*optclient.OtpGenResponse, error) {
	auth, err :=otpx.GenOtpKey(in.Issuer, in.Account)

	if err != nil {
		return &optclient.OtpGenResponse{
			Code: "1",
			Message: err.Error(),
		}, err
	}

	return &optclient.OtpGenResponse{
		Code: "0",
		Message: "Success",
		Data: &optclient.OtpData{
			Secret: auth.Code,
			Qrcode: auth.Path,
		},
	}, nil

}
