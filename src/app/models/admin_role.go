package models

import (
	"app"
)

type AdminRole struct {
	Id        int64    `xorm:"pk autoincr BIGINT"`
	RoleName  string   `xorm:"not null VARCHAR(32)"`
	IsSuper   bool     `xorm:"not null default false BOOL"`
	IsSys     bool     `xorm:"not null default false BOOL"`
	CreatedAt app.Time `xorm:"created"`
	UpdatedAt app.Time `xorm:"updated"`
	UserNames string   `xorm:"- <- ->"`
}
