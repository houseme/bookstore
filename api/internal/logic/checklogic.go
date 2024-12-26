package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/houseme/bookstore/api/internal/svc"
	"github.com/houseme/bookstore/api/internal/types"
	"github.com/houseme/bookstore/rpc/check/checker"
)

type CheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) CheckLogic {
	return CheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckLogic) Check(req types.CheckReq) (*types.CheckResp, error) {
	// todo: add your logic here and delete this line

	// 手动代码开始
	resp, err := l.svcCtx.Checker.Check(l.ctx, &checker.CheckReq{
		Id: req.Id,
	})
	if err != nil {
		logx.Error(err)
		return &types.CheckResp{}, err
	}

	return &types.CheckResp{
		Found: resp.Found,
		Price: resp.Price,
	}, nil
	// 手动代码结束
}
