package role

import (
	"context"
	"github.com/copo888/copo_otp/api/internal/svc"
	"github.com/copo888/copo_otp/api/internal/types"
	"github.com/copo888/copo_otp/helper/otpx"
	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewValidateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ValidateLogic {
	return ValidateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ValidateLogic) Validate(req types.OtpVaildReq) (resp *types.OtpVaildReply, err error) {

	return &types.OtpVaildReply{
		Vaild: otpx.Validate(req.PassCode, req.Secret),
	}, nil
}
