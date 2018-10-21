package models

import (
	"app"
	"github.com/go-xorm/xorm"
)

type AdminNavigationNode struct {
	Id                int64    `xorm:"pk autoincr BIGINT"`
	AdminNavigationId int64    `xorm:"BIGINT"`
	RouteAction       string   `xorm:"not null VARCHAR(255)"`
	Title             string   `xorm:"not null VARCHAR(255)"`
	Sort              int64    `xorm:"not null default 99 BIGINT"`
	CreatedAt         app.Time `xorm:"created"`
	UpdatedAt         app.Time `xorm:"updated"`
}

/**
角色分配传入admin_navigation_node_id返回父亲以及本身的Id集合
*/
func GetAllNavigationIds(db *xorm.Session, id int64, ids []int64) []int64 {
	admin_navigation := new(AdminNavigation)
	db.Id(id).Get(admin_navigation)
	if admin_navigation.ParentId > 0 {
		ids = append(ids, admin_navigation.Id)
		return GetAllNavigationIds(db, admin_navigation.ParentId, ids)
	} else {
		ids = append(ids, admin_navigation.Id)
		return ids
	}
}
