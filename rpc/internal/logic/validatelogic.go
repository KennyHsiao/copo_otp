package logic

import (
	"context"
	"github.com/copo888/copo_otp/helper/otpx"
	"github.com/copo888/copo_otp/rpc/optclient"

	"github.com/copo888/copo_otp/rpc/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
)

type ValidateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateLogic {
	return &ValidateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ValidateLogic) Validate(in *optclient.OtpVaildRequest) (*optclient.OtpVaildResponse, error) {

	return &optclient.OtpVaildResponse{
		Vaild: otpx.Validate(in.PassCode, in.Secret),
	}, nil
}
