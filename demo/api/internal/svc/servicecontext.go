package svc

import "github.com/chinagocoder/go-zero-pulsar-demo/demo/api/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
