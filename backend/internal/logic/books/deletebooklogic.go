package books

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"lowcode/internal/dao/query"
	"lowcode/internal/errcode"
	"lowcode/internal/svc"
	"lowcode/internal/types"
)

type DeleteBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBookLogic {
	return &DeleteBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBookLogic) DeleteBook(req *types.DeleteRequest) (resp *types.DeleteResponse, err error) {
	b := query.Use(l.svcCtx.DB).Book
	_, err = b.WithContext(l.ctx).Where(b.Ysid.Eq(req.ID)).Delete()
	if err != nil {
		return nil, errcode.NewDefaultError(err)
	}
	return &types.DeleteResponse{}, nil
}
