package test

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"testing"
)

func TestFileUploadFile(t *testing.T) {
	ctx := context.Background()
	endpoint := "127.0.0.1:9000"
	accessKeyId := "test"
	secretAccessKey := "12345678"
	useSSL := false
	cli, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln("创建MinIO客户端失败: ", err)
	}
	log.Println("创建MinIO客户端成功")
	bucketName := "test"
	location := "us-east-1"
	err = cli.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, err := cli.BucketExists(ctx, bucketName)
		if err == nil && exists {
			log.Printf("存储桶 %s 已经存在\n", bucketName)
		} else {
			log.Fatalln("查询存储桶状态异常： ", err)
		}
	}
	log.Printf("存储桶 %s 创建成功\n", bucketName)

	objectName := "zelda.jpeg"
	filePath := "zelda.jpeg"
	contentType := "application/jpeg"
	_, err = cli.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln("上传文件失败: ", err)
	}
	log.Printf("上传文件 %s 成功\n", objectName)
}
