package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/houseme/bookstore/rpc/check/check"
	"github.com/houseme/bookstore/rpc/check/internal/svc"
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

func (l *CheckLogic) Check(in *check.CheckReq) (*check.CheckResp, error) {
	// todo: add your logic here and delete this line

	// 手动代码开始
	resp, err := l.svcCtx.Model.FindOne(in.Id)
	if err != nil {
		return nil, err
	}

	return &check.CheckResp{
		Found: true,
		Price: resp.Price,
	}, nil
	// 手动代码结束
}
