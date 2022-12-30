package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MySQL struct {
		Port         int
		MaxOpenConns int
		MaxIdleConns int
		Host         string
		User         string
		Password     string
		DbName       string
	}
	Redis struct {
		Port     int
		DB       int
		PoolSize int
		Host     string
		Password string
	}
}
