package models

import (
	"app"
)

type AdminRoleNodeRoutes struct {
	Id                    int64    `xorm:"pk autoincr BIGINT"`
	AdminRoleId           int64    `xorm:"not null BIGINT"`
	AdminNavigationNodeId int64    `xorm:"not null BIGINT"`
	CreatedAt             app.Time `xorm:"DATETIME"`
	UpdatedAt             app.Time `xorm:"DATETIME"`
}
