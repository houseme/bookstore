package logic

import (
	"context"

	"github.com/houseme/bookstore/api/internal/svc"
	"github.com/houseme/bookstore/api/internal/types"
	"github.com/houseme/bookstore/rpc/add/adder"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddLogic {
	return AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req types.AddReq) (*types.AddResp, error) {
	// todo: add your logic here and delete this line

	// 手动代码开始
	resp, err := l.svcCtx.Adder.Add(l.ctx, &adder.AddReq{
		Book:  req.Book,
		Price: req.Price,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddResp{
		Ok: resp.Ok,
	}, nil
	// 手动代码结束
}
