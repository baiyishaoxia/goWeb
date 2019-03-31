package background

import (
	"app"
	"app/models/background"
	"config"
	"databases"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//region Remark:会员管理 Author:tang
func GetMemberList(c *gin.Context) {
	keywords := c.Query("keywords")
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.ParseInt(models.ReadConfig("sys.paginate"), 10, 64)
	data, num, all, page := models.GetUsersList(page-1, int(limit), keywords)
	files, _ := ioutil.ReadDir("./uploads/file/template/") //下载模板文件
	c.HTML(http.StatusOK, "member/list", gin.H{
		"Title":    "Background Index",
		"Data":     data,
		"Keywords": keywords,
		"Num":      num,
		"DownPage": float64(page + 1),
		"Page":     float64(page),
		"UpPage":   float64(page - 1),
		"All":      all,
		"Files":    files,
		"Time": func(time time.Time) string {
			return time.Format("2006-01-02 15:04:05")
		},
	})
}

//region Remark:新增 Author:tang
func GetMemberCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "member/add", gin.H{
		"Title": "Background Index",
	})
}
func PostMemberCreate(c *gin.Context) {
	if c.PostForm("name") == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "用户名不能为空！",
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
	//保存数据
	users := new(models.Users)
	if c.PostForm("is_lock") == "1" {
		users.IsLock = true
	}
	users.Password = app.Strmd5(c.PostForm("password"))
	users.Name = c.PostForm("name")
	users.Sex, _ = strconv.ParseInt(c.PostForm("sex"), 10, 64)
	users.HeadImg = c.PostForm("head_img")
	users.Email = c.PostForm("email")
	users.Phone = c.PostForm("phone")
	_, err := databases.Orm.Insert(users)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "保存成功",
		"url":    "/admin/member/list",
	})
}

//endregion
func GetMemberEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := new(models.Users)
	databases.Orm.Id(id).Get(user)
	c.HTML(http.StatusOK, "member/edit", gin.H{
		"Title": "Background Index",
		"Data":  user,
	})
}
func PostMemberEdit(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if c.PostForm("name") == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "用户名不能为空！",
		})
		return
	}
	users := new(models.Users)
	if models.GetUserExits(c.PostForm("name"), id) {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "该用户已存在",
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
	if c.PostForm("is_lock") == "1" {
		users.IsLock = true
	}
	if len(c.PostForm("password")) > 0 {
		users.Password = app.Strmd5(c.PostForm("password"))
	}
	users.Name = c.PostForm("name")
	users.HeadImg = c.PostForm("head_img")
	users.Sex, _ = strconv.ParseInt(c.PostForm("sex"), 10, 64)
	users.Email = c.PostForm("email")
	users.Phone = c.PostForm("phone")
	_, err := databases.Orm.ID(id).UseBool("is_lock").Update(users)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "保存成功",
		"url":    "/admin/member/list",
	})
}
func GetMemberDelList(c *gin.Context) {
	c.HTML(http.StatusOK, "member/del", gin.H{
		"Title": "Background Index",
	})
}

//endregion

//region Remark:修改密码 Author:tang
func GetMemberPassword(c *gin.Context) {
	user := new(models.Users)
	databases.Orm.Where("id=?", c.Param("id")).Get(user)
	c.HTML(http.StatusOK, "member/change_password", gin.H{
		"Title": "Background Index",
		"Data":  user,
	})
}
func PostMemberPassword(c *gin.Context) {
	id := c.Param("id")
	password := c.PostForm("password")
	password_c := c.PostForm("password_confirmation")
	if password == "" || password_c == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "密码不能为空",
		})
		return
	}
	if password != password_c {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "密码确认错误",
		})
		return
	}
	user := new(models.Users)
	databases.Orm.Id(id).Get(user)
	if len(password) > 0 {
		user.Password = app.Strmd5(password)
	}
	_, err := databases.Orm.ID(id).Update(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "操作失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "修改成功",
		"url":    "/admin/member/list",
	})
}

//endregion

//region Remark:查看详情 Author:tang
func GetMemberShow(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user_info := models.GetUserInfoById(id)
	c.HTML(http.StatusOK, "member/show", gin.H{
		"Title": "Background Index",
		"Data":  user_info,
	})
}

//endregion

//region Remark:删除 Author:tang
func PostMemberDel(c *gin.Context) {
	ids := c.PostFormArray("id[]")
	var users = new(models.Users)
	_, err := databases.Orm.In("id", ids).Delete(users)
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
		"url":    "/admin/member/list",
	})
	return
}

//endregion

//region Remark:是否停用 Author:tang
func GetMemberStatus(c *gin.Context) {
	id := c.Query("id")
	user := new(models.Users)
	databases.Orm.Id(id).Get(user)
	if user.IsLock == false {
		user.IsLock = true //已锁定
	} else {
		user.IsLock = false //未锁定
	}
	res, _ := databases.Orm.Id(id).Cols("is_lock").Update(user)
	if res >= 1 {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpSuccess,
			"info":   "设置成功",
			"url":    "/admin/member/list",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "设置失败",
			"url":    "/admin/member/list",
		})
	}
}

//endregion

//region Remark:导入 Author:tang
func GetImportCsv(c *gin.Context) {
	c.HTML(http.StatusOK, "member/csv", gin.H{
		"Title": "导入",
	})
}
func PostImportCsv(c *gin.Context) {
	url := c.PostForm("csv_url")
	if url == "" {
		c.JSON(http.StatusOK, gin.H{
			"status": config.HttpError,
			"info":   "请上传需要导入的csv文件！",
		})
		return
	}
	suff := strings.Index(url, ".")
	if suff != -1 {
		if url[suff:] != ".zip" {
			c.JSON(http.StatusOK, gin.H{
				"status": config.HttpError,
				"info":   "请导入csv格式的文件！",
			})
			return
		}
	}
	//创建文件夹 并 解压文件
	//path := "/uploads/file/" + time.Now().Format("2006/0102/") + "/" + time.Now().Format("150403/")
	//directory.DirectoryMkdir(path)
	//app.Unzip("."+url, "."+path)
	file, err := os.Open("." + url)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	// 这个方法体执行完成后，关闭文件
	defer file.Close()
	reader := csv.NewReader(file)
	j := 0
	db := databases.Orm.NewSession()
	for {
		j++
		// Read返回的是一个数组，它已经帮我们分割了，
		record, err := reader.Read()
		// 如果读到文件的结尾，EOF的优先级居然比nil还高！
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("记录集错误:", err.Error())
			return
		}
		if j > 1 {
			//序号,姓名,性别,电话,邮箱,城市,密码
			data := new(models.Users)
			number, _ := strconv.ParseInt(record[0], 10, 64)
			name := ConvertToString(record[1], "gbk", "utf-8")
			has, _ := db.Where("name=?", name).Get(data)
			var sex int64
			sex_c := ConvertToString(record[2], "gbk", "utf-8")
			if sex_c == "男" {
				sex = 1
			}
			if sex_c == "女" {
				sex = 2
			}
			if sex_c == "保密" {
				sex = 3
			}
			phone := record[3]
			email := record[4]
			city := ConvertToString(record[5], "gbk", "utf-8")
			password := record[6]
			if !has {
				add := &models.Users{Name: name, Sex: sex, Phone: phone, Email: email, City: city, Password: app.Strmd5(password)}
				_, err := db.Insert(add)
				if err != nil {
					db.Rollback()
				}
			}
			fmt.Println("正在导入序号为:", number, "的数据")
		}
	}
	db.Commit()
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"info":   "导入成功",
		"url":    "/admin/member/list",
	})
	return

}
func GetImportCsvDownload(c *gin.Context) {
	zipName := "./uploads/file/template/" + c.Param("name")
	c.File(zipName)
}

//endregion
