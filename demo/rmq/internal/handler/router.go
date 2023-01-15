package handler

import (
	"context"
	"github.com/chinagocoder/go-queue/pq"
	"github.com/chinagocoder/go-zero-pulsar-demo/demo/rmq/internal/logic"
	"github.com/chinagocoder/go-zero-pulsar-demo/demo/rmq/internal/svc"
)

func RegisterConsumer(queue *pq.PulsarQueues, svcCtx *svc.ServiceContext) {
	queue.AddQueue(
		svcCtx.Config.Pulsar,
		pq.WithHandle(logic.NewDemoConsumerLogic(context.Background(), svcCtx).Consumer),
	)
}
