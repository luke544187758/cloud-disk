package helper

import (
	"cloud-disk/core/define"
	"cloud-disk/core/pkg/snowflake"
	"context"
	"crypto/md5"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"math/rand"
	"net/http"
	"net/smtp"
	"path"
	"time"
)

func Md5(val string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(val)))
}

func GenerateToken(id int, identity int64, name string, expireTime int) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(expireTime)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	return token.SignedString([]byte(define.JwtKey))
}

func ParseToken(token string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(tk *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, nil
}

func MailSendCode(mail, code string) error {
	e := email.NewEmail()
	e.From = "luke544187758@163.com"
	e.To = []string{mail}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码为: <h1>" + code + "</h1>")
	err := e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "luke544187758@163.com", "ZUXFVDUUGWSEOHGN", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	return nil
}

func RandCode() string {
	sets := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += string(sets[rand.Intn(len(sets))])
	}
	return code
}

func MinioUpload(r *http.Request) (string, error) {
	ctx := context.Background()
	cli, err := minio.New(define.MinIOBucket, &minio.Options{
		Creds:  credentials.NewStaticV4(define.MinIOAccessKeyId, define.MinIOSecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return "", err
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		return "", err
	}

	bucketName := "cloud-disk"
	location := "us-east-1"
	err = cli.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, err := cli.BucketExists(ctx, bucketName)
		if err != nil || !exists {
			return "", err
		}
	}

	key := fmt.Sprintf("%d%s", snowflake.GenID(), path.Ext(header.Filename))
	_, err = cli.PutObject(ctx, bucketName, key, file, header.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/%s/%s", define.MinIOBucket, bucketName, key), nil
}
