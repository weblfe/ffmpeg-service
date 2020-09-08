package logic

import (
	"context"

	"github.com/weblfe/ffmpeg-service/rpc/check/internal/svc"
	check "github.com/weblfe/ffmpeg-service/rpc/check/pb"

	"github.com/tal-tech/go-zero/core/logx"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) Check(in *check.Request) (*check.Response, error) {
	// todo: add your logic here and delete this line

	return &check.Response{}, nil
}
