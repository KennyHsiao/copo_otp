package svc

import (
	"github.com/copo888/copo_otp/rpc/internal/config"
	ztrace "github.com/zeromicro/go-zero/core/trace"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {

	ztrace.StartAgent(ztrace.Config{
		Name:     c.Telemetry.Name,
		Endpoint: c.Telemetry.Endpoint,
		Batcher:  c.Telemetry.Batcher,
		Sampler:  c.Telemetry.Sampler,
	})

	return &ServiceContext{
		Config: c,
	}
}
