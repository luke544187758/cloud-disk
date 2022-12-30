package logic

import (
	"cloud-disk/core/models"
	"cloud-disk/core/pkg/snowflake"
	"context"
	"errors"
	"strconv"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateResponse, err error) {
	cnt, err := l.svcCtx.Engine.Where("name = ? AND parent_id = ?", req.Name, req.ParentId).Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("the name is exists")
	}
	ui, _ := strconv.ParseInt(userIdentity, 10, 64)
	data := &models.UserRepository{
		Identity:     snowflake.GenID(),
		UserIdentity: ui,
		ParentId:     req.ParentId,
		Name:         req.Name,
	}

	_, err = l.svcCtx.Engine.Insert(data)
	if err != nil {
		return nil, err
	}

	return &types.UserFolderCreateResponse{Identity: data.Identity}, nil
}
