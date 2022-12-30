package logic

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"cloud-disk/core/pkg/snowflake"
	"context"
	"errors"
	"log"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	resp = new(types.UserRegisterResponse)
	// 判断code 是否一致
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		err = errors.New("该邮箱的验证码为空")
		resp.Success = false
		resp.Error = err.Error()
		return resp, err
	}
	if code != req.Code {
		err = errors.New("验证码错误")
		resp.Success = false
		resp.Error = err.Error()
		return resp, err
	}
	// 判断用户名是否存在
	cnt, err := l.svcCtx.Engine.Where("name = ?", req.Name).Count(new(models.UserBasic))
	if err != nil {
		resp.Success = false
		resp.Error = err.Error()
		return resp, err
	}
	if cnt > 0 {
		err = errors.New("用户名已经存在")
		resp.Success = false
		resp.Error = err.Error()
		return resp, err
	}

	count, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return nil, err
	}
	if count > 0 {
		err = errors.New("邮箱已经被注册")
		resp.Success = false
		resp.Error = err.Error()
		return resp, err
	}

	// 数据入库
	user := &models.UserBasic{
		Identity: snowflake.GenID(),
		Name:     req.Name,
		Password: helper.Md5(req.Password),
		Email:    req.Email,
	}
	n, err := l.svcCtx.Engine.Insert(user)
	if err != nil {
		return nil, err
	}
	log.Println("insert user row:", n)
	resp.Email = req.Email
	resp.Success = true
	resp.Name = req.Name
	return resp, nil
}
