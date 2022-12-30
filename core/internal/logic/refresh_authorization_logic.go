package logic

import (
	"cloud-disk/core/define"
	"cloud-disk/core/helper"
	"context"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshAuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshAuthorizationLogic {
	return &RefreshAuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshAuthorizationLogic) RefreshAuthorization(req *types.RefreshAuthorizationRequest, authorization string) (resp *types.RefreshAuthorizationResponse, err error) {
	uc, err := helper.ParseToken(authorization)
	if err != nil {
		return nil, err
	}
	token, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, define.JwtExpireTime)
	if err != nil {
		return nil, err
	}
	refreshToken, err := helper.GenerateToken(uc.Id, uc.Identity, uc.Name, define.JwtExpireRefreshTime)
	if err != nil {
		return nil, err
	}
	resp = &types.RefreshAuthorizationResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}
	return resp, nil
}
