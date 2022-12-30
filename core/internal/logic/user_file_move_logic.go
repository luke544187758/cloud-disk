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

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveResponse, err error) {
	ui, _ := strconv.ParseInt(userIdentity, 10, 64)
	parentDate := new(models.UserRepository)
	has, err := l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.ParentIdentity, ui).Get(parentDate)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("the folder is not exist")
	}
	_, err = l.svcCtx.Engine.Where("identity = ?", req.Identity).Update(models.UserRepository{ParentId: parentDate.Id})

	return
}
