package logic

import (
	"context"
	"errors"
	"github.com/chinagocoder/go-queue/pq"
	"github.com/chinagocoder/go-zero-pulsar-demo/demo/api/internal/svc"
	"github.com/chinagocoder/go-zero-pulsar-demo/demo/api/internal/types"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProduceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProduceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProduceLogic {
	return &ProduceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProduceLogic) Produce(req *types.ProduceReq) (resp *types.ProduceResp, err error) {
	pusher, err := pq.NewPusher(l.svcCtx.Config.Pulsar)
	if err != nil {
		return nil, errors.New("连接pulsar失败")
	}
	messageId, err := pusher.Push(
		context.Background(), []byte(strconv.FormatInt(time.Now().UnixNano(), 10)), []byte(req.Msg),
	)

	if err != nil {
		return nil, errors.New("发送失败")
	} else {
		return &types.ProduceResp{
			Id: messageId.LedgerID,
		}, nil
	}
}
