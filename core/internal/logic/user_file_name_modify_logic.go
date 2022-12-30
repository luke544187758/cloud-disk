package logic

import (
	"cloud-disk/core/models"
	"context"
	"errors"
	"strconv"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameModifyLogic {
	return &UserFileNameModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameModifyLogic) UserFileNameModify(req *types.UserFileNameModifyRequest, userIdentity string) (resp *types.UserFileNameModifyResponse, err error) {
	cnt, err := l.svcCtx.Engine.Where("name = ? AND parent_id = (SELECT parent_id FROM user_repository ur WHERE ur.identity = ?)", req.Name, req.Identity).Count(new(models.UserRepository))
	if err != nil {
		return nil, err
	}

	if cnt > 0 {
		return nil, errors.New("the name is exists")
	}

	data := &models.UserRepository{Name: req.Name}
	ui, _ := strconv.ParseInt(userIdentity, 10, 64)
	_, err = l.svcCtx.Engine.Where("identity=? AND user_identity = ?", req.Identity, ui).Update(data)
	if err != nil {
		return nil, err
	}
	return
}
