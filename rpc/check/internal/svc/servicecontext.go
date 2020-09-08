package svc

import "github.com/weblfe/ffmpeg-service/rpc/check/internal/config"

type ServiceContext struct {
	c config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c: c,
	}
}
