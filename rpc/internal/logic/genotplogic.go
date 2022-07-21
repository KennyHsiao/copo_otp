package logic

import (
	"context"
	"github.com/copo888/copo_otp/helper/otpx"
	"github.com/copo888/copo_otp/rpc/internal/svc"
	"github.com/copo888/copo_otp/rpc/otpclient"
	"github.com/zeromicro/go-zero/core/logx"
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

func (l *GenOtpLogic) GenOtp(in *otpclient.OtpGenRequest) (*otpclient.OtpGenResponse, error) {
	auth, err := otpx.GenOtpKey(in.Issuer, in.Account)

	//span := trace.SpanFromContext(l.ctx)
	//
	//log.Println(">>>>>>>>",auth, span.SpanContext().TraceID(), span.SpanContext().SpanID())
	//span := trace.SpanFromContext(l.ctx)
	//defer span.End()
	//span.SetAttributes(attribute.KeyValue{
	//	Key:   "ccc",
	//	Value: attribute.StringValue("QQQQQ"),
	//})
	//TEST
	if err != nil {
		return &otpclient.OtpGenResponse{
			Code:    "1",
			Message: err.Error(),
		}, err
	}

	return &otpclient.OtpGenResponse{
		Code:    "0",
		Message: "Success",
		Data: &otpclient.OtpData{
			Secret: auth.Code,
			Qrcode: auth.Path,
		},
	}, nil

}
