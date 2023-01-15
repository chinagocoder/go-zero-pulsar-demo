package handler

import (
	"github.com/chinagocoder/go-zero-pulsar-demo/demo/api/internal/logic"
	"github.com/chinagocoder/go-zero-pulsar-demo/demo/api/internal/svc"
	"github.com/chinagocoder/go-zero-pulsar-demo/demo/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ProduceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProduceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewProduceLogic(r.Context(), svcCtx)
		resp, err := l.Produce(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
