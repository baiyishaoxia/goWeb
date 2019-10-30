package background

import (
	"app"
	"app/models"
	newredis "app/vendors/redis/models"
	"config"
	"databases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// region Remark:管理员列表 Author:tang
func GetAdminList(c *gin.Context) {
	admins := make([]models.BlogAdmin, 0)
	databases.Orm.Find(&admins)
	for k, v := range admins {
		admin_role := new(models.AdminRole)
		databases.Orm.Id(v.AdminRoleId).Get(admin_role)
		admins[k].Role = admin_role
	}

	c.HTML(http.StatusOK, "admin/list", gin.H{
		"Title": "Background Index",
		"Data":  admins,
		"Count": len(admins),
	})
}

//endregion

//region Remark:新增 Author:tang
func GetAdminCreate(c *gin.Context) {
	admin_roles := make([]models.AdminRole, 0)
	databases.Orm.Find(&admin_roles)
	//模版
	c.HTML(http.StatusOK, "admin/add", gin.H{
		"Title":    "Background Login",
		"RoleData": admin_roles,
	})
}
func PostAdminCreate(c *gin.Context) {
	if c.PostForm("username") == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "管理员不能为空！",
		})
		return
	}
	if c.PostForm("password") == "" || c.PostForm("password_confirmation") == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "密码不能为空！",
		})
		return
	}
	if c.PostForm("password") != c.PostForm("password_confirmation") {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "密码确认错误",
		})
		return
	}
	admin := new(models.BlogAdmin)
	res, err := databases.Orm.Where("email=?", c.PostForm("email")).Get(admin)
	if res {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "该邮箱已经存在",
		})
		return
	}
	//region 保存Admin数据
	if c.PostForm("is_lock") == "1" {
		admin.IsLock = true
	}
	admin_role_id, _ := strconv.ParseInt(c.PostForm("admin_role_id"), 10, 64)
	admin.AdminRoleId = admin_role_id
	admin.Email = c.PostForm("email")
	admin.Password = app.Strmd5(c.PostForm("password"))
	admin.Username = c.PostForm("username")
	admin.Mobile = c.PostForm("mobile")
	_, err = databases.Orm.Insert(admin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "保存成功",
		"url":    "/admin/list",
	})
}

//endregion

//region Remark:删除 Author:tang
func PostAdminDel(c *gin.Context) {
	ids := c.PostFormArray("id[]")
	if len(ids) == 0 {
		//单个删除
		ids = []string{}
		ids = append(ids, c.PostForm("id"))
	}
	var admin = new(models.BlogAdmin)
	_, err := databases.Orm.In("id", ids).NotIn("id", models.GetAdminInfo(c).Id).Delete(admin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "删除失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "删除成功",
		"url":    "/admin/list",
	})
	return
}

//endregion

//region Remark:修改 Author:tang
func GetAdminEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	admin_roles := make([]models.AdminRole, 0)
	databases.Orm.Find(&admin_roles)
	//admin数据
	admin := new(models.BlogAdmin)
	databases.Orm.Id(id).Get(admin)
	//模版
	c.HTML(http.StatusOK, "admin/edit", gin.H{
		"Title":     "Background Login",
		"RoleData":  admin_roles,
		"AdminData": admin,
	})
}
func PostAdminEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if c.PostForm("username") == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "管理员不能为空！",
		})
		return
	}
	if c.PostForm("password") != c.PostForm("password_confirmation") {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "密码确认错误",
		})
		return
	}
	//region 保存Admin数据
	admin := new(models.BlogAdmin)
	res, err := databases.Orm.Where("email=?", c.PostForm("email")).Where("id !=?", id).Get(admin)
	if res {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "该邮箱已经存在",
		})
		return
	}
	if c.PostForm("is_lock") == "1" {
		admin.IsLock = true
	}
	admin_role_id, _ := strconv.ParseInt(c.PostForm("admin_role_id"), 10, 64)
	admin.AdminRoleId = admin_role_id
	admin.Email = c.PostForm("email")
	if len(c.PostForm("password")) > 0 {
		admin.Password = app.Strmd5(c.PostForm("password"))
	}
	admin.Username = c.PostForm("username")
	admin.Mobile = c.PostForm("mobile")
	_, err = databases.Orm.ID(id).UseBool("is_lock").Update(admin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	//endregion
	//删除缓存
	var redis_key string = "admin:info:" + strconv.FormatInt(admin.Id, 10)
	newredis.DelKey(redis_key)
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "保存成功",
		"url":    "/admin/list",
	})
}

//endregion

//region Remark:更改管理员状态 Author:tang
func GetAdminStatus(c *gin.Context) {
	id := c.Query("id")
	admin := new(models.BlogAdmin)
	databases.Orm.Id(id).Get(admin)
	if admin.IsLock == false {
		admin.IsLock = true //已锁定
	} else {
		admin.IsLock = false //未锁定
	}
	res, _ := databases.Orm.Id(id).Cols("is_lock").Update(admin)
	if res >= 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "设置成功",
			"url":    "/admin/list",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "设置失败",
			"url":    "/admin/list",
		})
	}
}

//endregion
