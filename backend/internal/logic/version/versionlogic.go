package version

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"lowcode/internal/svc"
	"lowcode/internal/types"
	"lowcode/internal/version"
)

type VersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VersionLogic {
	return &VersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VersionLogic) Version() (resp *types.VersionResponse, err error) {
	return &types.VersionResponse{
		Version: version.WebVersion(),
	}, nil
}
