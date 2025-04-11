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

type AddBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddBookLogic {
	return &AddBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddBookLogic) AddBook(req *types.AddRequest) (resp *types.AddResponse, err error) {
	b := query.Use(l.svcCtx.DB).Book
	user := model.Book{Name: req.Name, Price: req.Price}
	err = b.WithContext(l.ctx).Create(&user)
	if err != nil {
		return nil, errcode.NewDefaultError(err)
	}
	return &types.AddResponse{}, nil
}
