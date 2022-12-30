package logic

import (
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"cloud-disk/core/models"
	"cloud-disk/core/pkg/snowflake"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadRequest) (resp *types.FileUploadResponse, err error) {
	rp := &models.RepositoryPool{
		Identity: snowflake.GenID(),
		Size:     req.Size,
		Hash:     req.Hash,
		Name:     req.Name,
		Ext:      req.Ext,
		Path:     req.Path,
	}
	_, err = l.svcCtx.Engine.Insert(rp)
	if err != nil {
		return nil, err
	}
	resp = new(types.FileUploadResponse)
	resp.Identity = rp.Identity
	resp.Ext = rp.Ext
	resp.Name = rp.Name
	return
}
