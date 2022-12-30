package models

import "time"

type ShareBasic struct {
	Id                     int
	ExpireTime             int
	ClickNum               int
	Identity               int64
	UserIdentity           int64
	UserRepositoryIdentity int64
	RepositoryIdentity     int64
	CreatedAt              time.Time `xorm:"created"`
	UpdatedAt              time.Time `xorm:"updated"`
	DeletedAt              time.Time `xorm:"deleted"`
}

func (table ShareBasic) TableName() string {
	return "share_basic"
}
