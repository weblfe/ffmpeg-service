package logic

import (
	"context"

	"github.com/weblfe/ffmpeg-service/internal/svc"
	"github.com/weblfe/ffmpeg-service/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ExecLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExecLogic(ctx context.Context, svcCtx *svc.ServiceContext) ExecLogic {
	return ExecLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExecLogic) Exec(req types.ExecReq) (*types.ExecResp, error) {
	return &types.ExecResp{}, nil
}
