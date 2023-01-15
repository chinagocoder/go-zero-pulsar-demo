package svc

import (
	"github.com/chinagocoder/go-zero-pulsar-demo/demo/rmq/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
