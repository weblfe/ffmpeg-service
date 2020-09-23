package logic

import (
	"context"

	"github.com/weblfe/ffmpeg-service/rpc/exec/internal/svc"
	exec "github.com/weblfe/ffmpeg-service/rpc/exec/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type ExecLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExecLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExecLogic {
	return &ExecLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExecLogic) Exec(in *exec.Request) (*exec.Response, error) {
	// todo: add your logic here and delete this line

	return &exec.Response{}, nil
}
