package models

import "time"

type RepositoryPool struct {
	Id        int
	Identity  int64
	Size      int64
	Hash      string
	Name      string
	Ext       string
	Path      string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (table RepositoryPool) TableName() string {
	return "repository_pool"
}
