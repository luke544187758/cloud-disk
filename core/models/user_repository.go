package models

import "time"

type UserRepository struct {
	Id                 int
	ParentId           int
	Identity           int64
	UserIdentity       int64
	RepositoryIdentity int64
	Ext                string
	Name               string
	CreatedAt          time.Time `xorm:"created"`
	UpdatedAt          time.Time `xorm:"updated"`
	DeletedAt          time.Time `xorm:"deleted"`
}

func (table UserRepository) TableName() string {
	return "user_repository"
}
