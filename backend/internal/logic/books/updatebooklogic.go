package books

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"lowcode/internal/dao/model"
	"lowcode/internal/dao/query"
	"lowcode/internal/errcode"
	"lowcode/internal/svc"
	"lowcode/internal/types"
)

type UpdateBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBookLogic {
	return &UpdateBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateBookLogic) UpdateBook(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	b := query.Use(l.svcCtx.DB).Book
	logx.Info("req = ", req)
	_, err = b.WithContext(l.ctx).Where(b.Ysid.Eq(req.ID)).Updates(model.Book{Name: req.Name, Price: req.Price})
	if err != nil {
		return nil, errcode.NewDefaultError(err)
	}
	return &types.UpdateResponse{}, nil
}
