package handler

import (
	"net/http"

	"github.com/weblfe/ffmpeg-service/internal/logic"
	"github.com/weblfe/ffmpeg-service/internal/svc"
	"github.com/weblfe/ffmpeg-service/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func execHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExecReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewExecLogic(r.Context(), ctx)
		resp, err := l.Exec(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
