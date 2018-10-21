package models

import (
	"app"
)

type AdminRoleNode struct {
	Id                int64    `xorm:"pk autoincr BIGINT"`
	AdminRoleId       int64    `xorm:"not null BIGINT"`
	AdminNavigationId int64    `xorm:"not null BIGINT"`
	CreatedAt         app.Time `xorm:"created"`
	UpdatedAt         app.Time `xorm:"updated"`
}
