package user

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"lowcode/internal/dao/query"
	"lowcode/internal/errcode"
	"lowcode/internal/helper"
	"lowcode/internal/svc"
	"lowcode/internal/types"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	if len(strings.TrimSpace(req.Name)) == 0 || len(strings.TrimSpace(req.Passwd)) == 0 {
		wErr := errors.New("name or passwd is empty.")
		return nil, errcode.NewCodeError(errcode.ErrAuthInvalidParam, wErr)
	}
	u := query.Use(l.svcCtx.DB).User
	user, err := u.WithContext(l.ctx).Where(u.Name.Eq(req.Name)).First()
	if err != nil {
		return nil, errcode.NewCodeError(errcode.ErrAuthNameNotExit, err)
	}
	err = helper.PwdVerify(req.Passwd, user.PwdHash)
	if err != nil {
		return nil, errcode.NewCodeError(errcode.ErrAuthPasswdNotMatch, err)
	}
	now := time.Now().Unix()

	accessSecret := l.svcCtx.Config.Auth.AccessSecret
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	token, err := helper.GetJwtToken(accessSecret, now, accessExpire, user.Ysid)
	if err != nil {
		return nil, errcode.NewDefaultError(err)
	}
	return &types.LoginResponse{
		JwtToken: "Bearer " + token,
	}, nil
}
