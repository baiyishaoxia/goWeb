package background

import (
	"app/models"
	newredis "app/vendors/redis/models"
	"config"
	"databases"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
)

func GetNavigation(c *gin.Context) {
	c.String(http.StatusOK, "%s", models.NavHtml(models.GetAdminInfo(c)))
}

func GetNavigationList(c *gin.Context) {
	var data []models.AdminNavigation
	//判断redis是否有缓存数据
	var redis_key string = "admin:navlist:list"
	if res, _ := newredis.Exists(redis_key); res == true {
		valueBytes, _ := redis.Bytes(newredis.Get(redis_key))
		json.Unmarshal(valueBytes, &data)
	} else {
		//获取数据
		admin_navigations := make([]models.AdminNavigation, 0)
		databases.Orm.Asc("sort").Find(&admin_navigations)
		for k, v := range admin_navigations {
			admin_navigations[k].Title = "<span class='folder-open'></span>" + v.Title
			admin_navigation_nodes := make([]models.AdminNavigationNode, 0)
			databases.Orm.Cols("title").Where("admin_navigation_id =?", v.Id).Find(&admin_navigation_nodes)
			for _, v1 := range admin_navigation_nodes {
				implode_fh := ""
				if admin_navigations[k].NodeTitles != "" {
					implode_fh = "，"
				}
				admin_navigations[k].NodeTitles = admin_navigations[k].NodeTitles + template.HTML(implode_fh+v1.Title)
			}
		}
		data = models.AdminNavigationMerge(admin_navigations, "<span class='folder-line'></span>", 0, 0)
		//缓存到redis
		value, _ := json.Marshal(data)
		newredis.Set(redis_key, value, 60*60)
	}
	//模版
	c.HTML(http.StatusOK, "admin/permission", gin.H{
		"Title": "Background Login",
		"Data":  data,
		"Count": len(data),
	})
}
func GetNavigationCreate(c *gin.Context) {
	//模板
	id, _ := strconv.Atoi(c.Param("id"))
	c.HTML(http.StatusOK, "admin/permission_add", gin.H{
		"Title":     "Background Login",
		"Data":      models.NavigationRoleTreeData(),
		"Parent_id": id,
	})
}
func PostNavigationCreate(c *gin.Context) {
	db := databases.Orm.NewSession()
	defer db.Close()
	db.Begin()

	//region Post数据
	parent_id, _ := strconv.ParseInt(c.PostForm("parent_id"), 10, 64)
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	is_show := false
	if c.PostForm("is_show") == "true" {
		is_show = true
	}
	is_sys := false
	if c.PostForm("is_sys") == "true" {
		is_sys = true
	}
	//endregion

	//region 创建admin_navigation数据-单条
	var err error
	var admin_navigation_id int64
	if parent_id == 0 {
		admin_navigation := new(models.AdminNavigation2)
		admin_navigation.Sort = sort
		admin_navigation.IsShow = is_show
		admin_navigation.IsSys = is_sys
		admin_navigation.Title = c.PostForm("title")
		admin_navigation.Ico = c.PostForm("ico")
		admin_navigation.Url = c.PostForm("url")
		_, err = db.Table("admin_navigation").Insert(admin_navigation)
		admin_navigation_id = admin_navigation.Id
	} else {
		admin_navigation := new(models.AdminNavigation1)
		admin_navigation.ParentId = parent_id
		admin_navigation.Sort = sort
		admin_navigation.IsShow = is_show
		admin_navigation.IsSys = is_sys
		admin_navigation.Title = c.PostForm("title")
		admin_navigation.Ico = c.PostForm("ico")
		admin_navigation.Url = c.PostForm("url")
		_, err = db.Table("admin_navigation").Insert(admin_navigation)
		admin_navigation_id = admin_navigation.Id
	}
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion

	//region 如果nodeids为空，直接事务提交成功
	if c.PostForm("nodeids") == "" {
		db.Commit()
		newredis.DelKeyByPrefix("admin:navlist") //删除缓存
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "保存成功",
			"url":    "/admin/permission/list",
		})
		return
	}
	//endregion

	//region 插入-AdminNavigationNode
	index_max, _ := strconv.Atoi(c.PostForm("nodeids"))
	admin_navigation_nodes := make([]models.AdminNavigationNode, index_max)
	for i := 0; i < index_max; i++ {
		title := c.PostForm("node[" + strconv.Itoa(i) + "][title]")
		route_action := c.PostForm("node[" + strconv.Itoa(i) + "][route_action]")
		sort1, _ := strconv.ParseInt(c.PostForm("node["+strconv.Itoa(i)+"][sort]"), 10, 64)
		admin_navigation_nodes[i].Title = title
		admin_navigation_nodes[i].RouteAction = route_action
		admin_navigation_nodes[i].Sort = sort1
		admin_navigation_nodes[i].AdminNavigationId = admin_navigation_id
	}
	_, err = db.Insert(&admin_navigation_nodes)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion

	//region 事务提交和返回信息
	db.Commit()
	newredis.DelKeyByPrefix("admin:navlist") //删除缓存
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "保存成功",
		"url":    "/admin/permission/list",
	})
	//endregion
}
func GetNavigationEdit(c *gin.Context) {
	//编辑数据
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	admin_navigation := new(models.AdminNavigation)
	result, err := databases.Orm.Id(id).Get(admin_navigation)
	if err != nil {
		c.String(http.StatusOK, "Error:%s", "数据异常")
		return
	}
	if result == false {
		c.String(http.StatusOK, "Error:%s", "没有数据")
		return
	}
	//节点数据
	admin_navigation_nodes := make([]models.AdminNavigationNode, 0)
	databases.Orm.Where("admin_navigation_id=?", id).Asc("sort").Find(&admin_navigation_nodes)
	//模版
	c.HTML(http.StatusOK, "admin/permission_edit", gin.H{
		"Title":           "Background Login",
		"Data":            models.NavigationRoleTreeData(),
		"Navigation":      admin_navigation,
		"NavigationNodes": admin_navigation_nodes,
		"NodeLen":         len(admin_navigation_nodes),
	})
}
func PostNavigationEdit(c *gin.Context) {
	admin_navigation_id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	db := databases.Orm.NewSession()
	defer db.Close()
	db.Begin()

	//region Post数据
	parent_id, _ := strconv.ParseInt(c.PostForm("parent_id"), 10, 64)
	sort, _ := strconv.ParseInt(c.PostForm("sort"), 10, 64)
	is_show := false
	if c.PostForm("is_show") == "true" {
		is_show = true
	}
	is_sys := false
	if c.PostForm("is_sys") == "true" {
		is_sys = true
	}
	//endregion

	//region 更新admin_navigation数据-单条
	var err error
	admin_navigation := new(models.AdminNavigation)
	if parent_id > 0 {
		admin_navigation.ParentId = parent_id
	}
	admin_navigation.Sort = sort
	admin_navigation.IsShow = is_show
	admin_navigation.IsSys = is_sys
	admin_navigation.Title = c.PostForm("title")
	admin_navigation.Ico = c.PostForm("ico")
	admin_navigation.Url = c.PostForm("url")
	_, err = db.ID(admin_navigation_id).UseBool("is_show", "is_sys").Update(admin_navigation)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion

	//region 如果nodeids为空，直接事务提交成功
	if c.PostForm("nodeids") == "" {
		db.Commit()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "修改成功",
			"url":    "/admin/permission/list",
		})
		return
	}
	//endregion

	//region 删除-AdminNavigationNode
	admin_navigation_node := new(models.AdminNavigationNode)
	_, err = databases.Orm.Where("admin_navigation_id=?", admin_navigation_id).Delete(admin_navigation_node)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion

	//region 插入-AdminNavigationNode
	index_max, _ := strconv.Atoi(c.PostForm("nodeids"))
	admin_navigation_nodes := make([]models.AdminNavigationNode, index_max)
	for i := 0; i < index_max; i++ {
		title := c.PostForm("node[" + strconv.Itoa(i) + "][title]")
		route_action := c.PostForm("node[" + strconv.Itoa(i) + "][route_action]")
		sort1, _ := strconv.ParseInt(c.PostForm("node["+strconv.Itoa(i)+"][sort]"), 10, 64)
		admin_navigation_nodes[i].Title = title
		admin_navigation_nodes[i].RouteAction = route_action
		admin_navigation_nodes[i].Sort = sort1
		admin_navigation_nodes[i].AdminNavigationId = admin_navigation_id
	}
	_, err = db.Insert(&admin_navigation_nodes)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion

	//region 事务提交和返回信息
	db.Commit()
	newredis.DelKeyByPrefix("admin:navlist") //删除缓存
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "修改成功",
		"url":    "/admin/permission/list",
	})
	//endregion
}
func PostNavigationDel(c *gin.Context) {
	ids := c.PostFormArray("id[]")
	if len(ids) == 0 {
		//单个删除
		ids = []string{}
		ids = append(ids, c.PostForm("id"))
	}
	db := databases.Orm.NewSession()
	db.Begin()
	//AdminNavigationNode操作
	admin_navigation_node := new(models.AdminNavigationNode)
	_, err := db.In("admin_navigation_id", ids).Delete(admin_navigation_node)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//AdminNavigation操作
	admin_navigation := new(models.AdminNavigation)
	_, err = db.In("id", ids).Delete(admin_navigation)
	if err != nil {
		db.Rollback()
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	db.Commit()
	newredis.DelKeyByPrefix("admin:navlist") //删除缓存
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "操作成功",
		"url":    "/admin/permission/list",
	})
}
func PostNavigationSave(c *gin.Context) {
	ids := c.PostFormArray("data[sort][]")
	for _, v := range ids {
		id, _ := strconv.ParseInt(v, 10, 64)
		admin_navigation := new(models.AdminNavigation)
		sort, _ := strconv.ParseInt(c.PostForm("data["+v+"][sort]"), 10, 64)
		admin_navigation.Sort = sort
		databases.Orm.ID(id).Cols("sort").Update(admin_navigation)
	}
	newredis.DelKeyByPrefix("admin:navlist") //删除缓存
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "操作成功",
		"url":    "/admin/permission/list",
	})
}
