// Code generated by goctl. DO NOT EDIT!
// Source: exec.proto

package server

import (
	"context"

	"github.com/weblfe/ffmpeg-service/rpc/exec/internal/logic"
	"github.com/weblfe/ffmpeg-service/rpc/exec/internal/svc"
	exec "github.com/weblfe/ffmpeg-service/rpc/exec/pb"
)

type ExecutorServer struct {
	svcCtx *svc.ServiceContext
}

func NewExecutorServer(svcCtx *svc.ServiceContext) *ExecutorServer {
	return &ExecutorServer{
		svcCtx: svcCtx,
	}
}

func (s *ExecutorServer) Exec(ctx context.Context, in *exec.Request) (*exec.Response, error) {
	l := logic.NewExecLogic(ctx, s.svcCtx)
	return l.Exec(in)
}
