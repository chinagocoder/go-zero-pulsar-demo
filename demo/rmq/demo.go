package main

import (
	"flag"
	"github.com/chinagocoder/go-queue/pq"
	"github.com/chinagocoder/go-zero-pulsar-demo/demo/rmq/internal/config"
	"github.com/chinagocoder/go-zero-pulsar-demo/demo/rmq/internal/handler"
	"github.com/chinagocoder/go-zero-pulsar-demo/demo/rmq/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/demo.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)

	// 注册消费者
	queue := pq.MustNewQueue()
	defer queue.Stop()
	handler.RegisterConsumer(queue, ctx)

	// 启动消费者
	queue.Start()

}
