package logic

import (
	"context"
	"fmt"
	"github.com/chinagocoder/go-zero-pulsar-demo/demo/rmq/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type DemoConsumerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDemoConsumerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoConsumerLogic {
	return &DemoConsumerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DemoConsumerLogic) Consumer(k, v string) error {
	// 打印消息内容，根据实际业务自行发挥
	fmt.Printf("pulsar content: time:%s content:%s\n", time.Now().Format("2006-01-02 15:04:05"), v)
	return nil
}
