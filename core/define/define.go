package define

import "github.com/dgrijalva/jwt-go"

type UserClaim struct {
	Id       int
	Identity int64
	Name     string
	jwt.StandardClaims
}

var (
	JwtKey         = "cloud-disk-key"
	CodeLength     = 6
	CodeExpireTime = 600

	MinIOBucket          = "127.0.0.1:9000"
	MinIOAccessKeyId     = "test"
	MinIOSecretAccessKey = "12345678"
	UseSSL               = false

	PageSize             = 20
	DateTimeFormat       = "2006-01-02 15:04:05"
	JwtExpireTime        = 1800
	JwtExpireRefreshTime = 1800
)
