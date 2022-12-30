package handler

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		file, header, err := r.FormFile("file")
		if err != nil {
			return
		}

		// 判断文件是否存在
		buf := make([]byte, header.Size)
		_, err = file.Read(buf)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(buf))

		rp := new(models.RepositoryPool)
		exist, err := svcCtx.Engine.Where("hash = ?", hash).Get(rp)
		if err != nil {
			return
		}
		if exist {
			httpx.OkJson(w, &types.FileUploadResponse{Identity: rp.Identity, Ext: rp.Ext, Name: rp.Name})
			return
		}

		// 往minio中存储文件
		minioPath, err := helper.MinioUpload(r)
		if err != nil {
			return
		}

		req.Name = header.Filename
		req.Size = header.Size
		req.Ext = path.Ext(header.Filename)
		req.Hash = hash
		req.Path = minioPath

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
