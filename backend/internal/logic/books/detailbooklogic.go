package books

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"lowcode/internal/dao/query"
	"lowcode/internal/errcode"
	"lowcode/internal/svc"
	"lowcode/internal/types"
)

type DetailBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailBookLogic {
	return &DetailBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailBookLogic) DetailBook(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	b := query.Use(l.svcCtx.DB).Book
	book, err := b.WithContext(l.ctx).Where(b.Ysid.Eq(req.ID)).First()
	if err != nil {
		return nil, errcode.NewDefaultError(err)
	}
	return &types.DetailResponse{
		Book: types.Book{ID: book.Ysid, Name: book.Name, Price: book.Price},
	}, nil
}
