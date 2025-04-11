package books

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"lowcode/internal/dao/query"
	"lowcode/internal/errcode"
	"lowcode/internal/helper"
	"lowcode/internal/svc"
	"lowcode/internal/types"
)

type ListBooksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListBooksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListBooksLogic {
	return &ListBooksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListBooksLogic) ListBooks(req *types.ListRequest) (resp *types.ListResponse, err error) {
	logx.Infof("userid: %v", helper.GetJWTUserID(l.ctx))
	b := query.Use(l.svcCtx.DB).Book
	books, err := b.WithContext(l.ctx).Limit(10).Find()
	if err != nil {
		return nil, errcode.NewDefaultError(err)
	}
	tBooks := make([]types.Book, 0)
	for _, book := range books {
		tBooks = append(tBooks, types.Book{ID: book.Ysid, Name: book.Name, Price: book.Price})
	}
	return &types.ListResponse{
		Books: tBooks,
	}, nil
}
