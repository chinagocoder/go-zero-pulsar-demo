package config

import (
	"github.com/chinagocoder/go-queue/pq"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Pulsar pq.Conf
}
