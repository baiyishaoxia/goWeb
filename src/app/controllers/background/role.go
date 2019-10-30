package background

import (
	"app"
	"app/models"
	newredis "app/vendors/redis/models"
	"config"
	"databases"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//region Remark:角色列表 Author:tang
func GetRoleList(c *gin.Context) {
	admin_roles := make([]models.AdminRole, 0)
	databases.Orm.Find(&admin_roles)
	//获取当前角色对应的用户列表
	for key, val := range admin_roles {
		admin := new([]models.BlogAdmin)
		databases.Orm.Id(val.Id).Find(admin)
		for _, v := range *admin {
			if v.AdminRoleId == val.Id {
				admin_roles[key].UserNames += v.Username + "、"
			}
		}
	}
	//模版
	c.HTML(http.StatusOK, "admin/role", gin.H{
		"Title":      "Background Login",
		"AdminRoles": admin_roles,
		"Count":      len(admin_roles),
	})
}

//endregion

func GetRoleCreate(c *gin.Context) {
	//模版
	c.HTML(http.StatusOK, "admin/role_add", gin.H{
		"Title":              "Background Login",
		"NavigationRoleData": models.NavigationRoleData(),
	})
}
func PostRoleCreate(c *gin.Context) {
	db := databases.Orm.NewSession()
	defer db.Close()
	db.Begin()
	//region 保存AdminRole数据
	role := new(models.AdminRole)
	if c.PostForm("role_type") == "0" {
		role.IsSuper = true
	}
	role.RoleName = c.PostForm("role_name")
	if c.PostForm("is_sys") == "true" {
		role.IsSys = true
	}
	_, err := db.Insert(role)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion

	//region 保存AdminRoleNode数据
	nodes := c.PostFormArray("route[]")
	admin_navigation_nodes := make([]models.AdminNavigationNode, 0)
	db.In("id", nodes).Cols("admin_navigation_id").GroupBy("admin_navigation_id").Find(&admin_navigation_nodes)
	//----获取所有的AdminNavigationId，包括父级的
	var allids []int64
	for _, v := range admin_navigation_nodes {
		var ids []int64
		id := models.GetAllNavigationIds(db, v.AdminNavigationId, ids)
		for _, v1 := range id {
			allids = append(allids, v1)
		}
	}
	allids = app.RemoveDuplicateInt64(allids)
	//----到数据库
	rolenodes := make([]models.AdminRoleNode, len(allids))
	for k, v := range allids {
		rolenodes[k].AdminRoleId = role.Id
		rolenodes[k].AdminNavigationId = v
	}
	_, err = db.Insert(&rolenodes)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion

	//region 保存AdminRoleNodeRoute
	rolenoderoutes := make([]models.AdminRoleNodeRoutes, len(nodes))
	for k, v := range nodes {
		rolenoderoutes[k].AdminRoleId = role.Id
		admin_navigation_id, _ := strconv.ParseInt(v, 10, 64)
		rolenoderoutes[k].AdminNavigationNodeId = admin_navigation_id
	}
	_, err = db.Insert(&rolenoderoutes)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion
	db.Commit()
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "保存成功",
		"url":    "/admin/role/list",
	})
}

//region Remark:修改 Author:tang
func GetRoleEdit(c *gin.Context) {
	id := c.Param("id")
	//region AdminRole数据
	admin_role := new(models.AdminRole)
	databases.Orm.Id(id).Get(admin_role)
	//endregion

	//region AdminRoleNodeRoute数据Id集合
	rolenoderoutes := make([]models.AdminRoleNodeRoutes, 0)
	databases.Orm.Where("admin_role_id=?", id).Cols("admin_navigation_node_id").Find(&rolenoderoutes)
	var checkid string = ""
	for _, v := range rolenoderoutes {
		pj := ","
		if checkid == "" {
			pj = ""
		}
		checkid = checkid + pj + strconv.FormatInt(v.AdminNavigationNodeId, 10)
	}
	fmt.Println(checkid)
	//endregion
	c.HTML(http.StatusOK, "admin/role_edit", gin.H{
		"Title":              "Background Login",
		"NavigationRoleData": models.NavigationRoleData(),
		"Role":               admin_role,
		"RoleRouteIds":       checkid,
	})
}
func PostRoleEdit(c *gin.Context) {
	db := databases.Orm.NewSession()
	defer db.Close()
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	db.Begin()
	//region 更新AdminRole数据
	role := new(models.AdminRole)
	if c.PostForm("role_type") == "0" {
		role.IsSuper = true
	}
	role.RoleName = c.PostForm("role_name")
	if c.PostForm("is_sys") == "true" {
		role.IsSys = true
	}
	_, err := db.ID(id).Cols("is_super", "role_name", "is_sys").UseBool("role_type", "is_sys").Update(role)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion

	//region 删除旧的AdminRoleNode数据
	admin_role_node := new(models.AdminRoleNode)
	_, err = databases.Orm.Where("admin_role_id=?", id).Delete(admin_role_node)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion

	//region 保存AdminRoleNode数据
	nodes := c.PostFormArray("route[]")
	admin_navigation_nodes := make([]models.AdminNavigationNode, 0)
	db.In("id", nodes).Cols("admin_navigation_id").GroupBy("admin_navigation_id").Find(&admin_navigation_nodes)
	//----获取所有的AdminNavigationId，包括父级的
	var allids []int64
	for _, v := range admin_navigation_nodes {
		var ids []int64
		id := models.GetAllNavigationIds(db, v.AdminNavigationId, ids)
		for _, v1 := range id {
			allids = append(allids, v1)
		}
	}
	allids = app.RemoveDuplicateInt64(allids)
	//----到数据库
	rolenodes := make([]models.AdminRoleNode, len(allids))
	for k, v := range allids {
		rolenodes[k].AdminRoleId = id
		rolenodes[k].AdminNavigationId = v
	}
	_, err = db.Insert(&rolenodes)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion

	//region 删除旧的AdminRoleNodeRoute数据
	rolenoderoute := new(models.AdminRoleNodeRoutes)
	_, err = databases.Orm.Where("admin_role_id=?", id).Delete(rolenoderoute)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion

	//region 保存AdminRoleNodeRoute
	rolenoderoutes := make([]models.AdminRoleNodeRoutes, len(nodes))
	for k, v := range nodes {
		rolenoderoutes[k].AdminRoleId = id
		admin_navigation_id, _ := strconv.ParseInt(v, 10, 64)
		rolenoderoutes[k].AdminNavigationNodeId = admin_navigation_id
	}
	_, err = db.Insert(&rolenoderoutes)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion
	db.Commit()
	newredis.DelKeyByPrefix("admin:power")
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "保存成功",
		"url":    "/admin/role/list",
	})
}

//endregion

//region Remark:删除 Author:tang
func PostRoleDel(c *gin.Context) {
	ids := c.PostFormArray("id[]")
	if len(ids) == 0 {
		//单个删除
		ids = []string{}
		ids = append(ids, c.PostForm("id"))
	}
	db := databases.Orm.NewSession()
	db.Begin()
	//region 删除AdminRoleNodeRoute
	admin_role_node_routes := new(models.AdminRoleNodeRoutes)
	_, err := db.In("admin_role_id", ids).Delete(admin_role_node_routes)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion

	//region AdminRoleNode数据
	admin_role_nodes := new(models.AdminRoleNode)
	_, err = db.In("admin_role_id", ids).Delete(admin_role_nodes)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion

	//region AdminRole数据
	admin_roles := new(models.AdminRole)
	_, err = db.In("id", ids).Delete(admin_roles)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion
	db.Commit()
	newredis.DelKeyByPrefix("admin:power")
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "操作成功",
		"url":    "/admin/role/list",
	})
}

//endregion
