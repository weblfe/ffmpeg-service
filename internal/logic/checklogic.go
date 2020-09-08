package logic

import (
	"context"

	"github.com/weblfe/ffmpeg-service/internal/svc"
	"github.com/weblfe/ffmpeg-service/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckLogic {
	return CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(req types.CheckReq) (*types.CheckResp, error) {
	return &types.CheckResp{}, nil
}
