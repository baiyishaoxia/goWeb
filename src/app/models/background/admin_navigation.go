package models

import (
	"app"
	newredis "app/vendors/redis/models"
	"databases"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"html/template"
	"strconv"
	"strings"
	"time"
)

type AdminNavigation struct {
	Id            int64                `xorm:"pk autoincr BIGINT"`
	ParentId      int64                `xorm:"BIGINT"`
	SiteId        int64                `xorm:"BIGINT"`
	SiteChannelId int64                `xorm:"BIGINT"`
	Title         string               `xorm:"not null VARCHAR(255)"`
	Url           string               `xorm:"not null VARCHAR(255)"`
	Ico           string               `xorm:"VARCHAR(255)"`
	Sort          int64                `xorm:"not null default 99 BIGINT"`
	IsShow        bool                 `xorm:"not null default false BOOL"`
	IsSys         bool                 `xorm:"not null default false BOOL"`
	CreatedAt     time.Time            `xorm:"created"`
	UpdatedAt     time.Time            `xorm:"updated"`
	Node          *AdminNavigationNode `xorm:"- <- ->"`
	NodeTitles    template.HTML        `xorm:"- <- ->"`
	TitleHtml     template.HTML        `xorm:"- <- ->"`
	Level         int                  `xorm:"- <- ->"`
}

/**
导航节点编辑和添加--parentid不为空，SiteId和SiteChannelId为空
*/
type AdminNavigation1 struct {
	Id        int64    `xorm:"pk autoincr BIGINT"`
	ParentId  int64    `xorm:"BIGINT"`
	Title     string   `xorm:"not null VARCHAR(255)"`
	Url       string   `xorm:"not null VARCHAR(255)"`
	Ico       string   `xorm:"VARCHAR(255)"`
	Sort      int64    `xorm:"not null default 99 BIGINT"`
	IsShow    bool     `xorm:"not null default false BOOL"`
	IsSys     bool     `xorm:"not null default false BOOL"`
	CreatedAt app.Time `xorm:"created"`
	UpdatedAt app.Time `xorm:"updated"`
}

/**
导航节点编辑和添加--parentid为空，SiteId和SiteChannelId为空
*/
type AdminNavigation2 struct {
	Id        int64    `xorm:"pk autoincr BIGINT"`
	Title     string   `xorm:"not null VARCHAR(255)"`
	Url       string   `xorm:"not null VARCHAR(255)"`
	Ico       string   `xorm:"VARCHAR(255)"`
	Sort      int64    `xorm:"not null default 99 BIGINT"`
	IsShow    bool     `xorm:"not null default false BOOL"`
	IsSys     bool     `xorm:"not null default false BOOL"`
	CreatedAt app.Time `xorm:"created"`
	UpdatedAt app.Time `xorm:"updated"`
}

/**
字段Parameter为json反解析后的结构
*/
type Parameter struct {
	SiteId    int `json:"site_id,int"`
	ChannelId int `json:"channel_id,int"`
	PageId    int `json:"page_id,nt"`
}

///**
//后台的路由和参数生成URL路径
// */
func ParameterJsonDecode(action string, par string) string {
	parameter := &Parameter{}
	json.Unmarshal([]byte(par), parameter)
	url := action + "?"
	if par == "" {
		url = action
	}
	var i int = 0
	var connectpar string = "&"
	if parameter.SiteId != 0 {
		url = url + "site_id=" + strconv.Itoa(parameter.SiteId)
		//fmt.Println(url,parameter.SiteId,parameter.PageId,parameter.ChannelId,par)
		i++
	}
	if parameter.PageId != 0 {
		if i == 0 {
			connectpar = ""
		}
		url = url + connectpar + "page_id=" + strconv.Itoa(parameter.PageId)
		//fmt.Println(url,parameter.SiteId,parameter.PageId,parameter.ChannelId,par)
		i++
	}
	if parameter.ChannelId != 0 {
		if i == 0 {
			connectpar = ""
		}
		url = url + connectpar + "channel_id=" + strconv.Itoa(parameter.ChannelId)
		//fmt.Println(url,parameter.SiteId,parameter.PageId,parameter.ChannelId,par)
	}
	return "/admin" + url
}

func AdminNavigationMerge(data []AdminNavigation, html string, pid int64, level int) (res []AdminNavigation) {
	for _, v := range data {
		if v.ParentId == pid {
			v.Level = level + 1
			newhtml := strings.Repeat(html, v.Level)
			v.TitleHtml = template.HTML(newhtml + v.Title)
			res = append(res, v)
			res1 := AdminNavigationMerge(data, html, v.Id, level+1)
			for _, v1 := range res1 {
				res = append(res, v1)
			}
		}
	}
	return res
}

/**
后台导航菜单
*/
func NavHtml(admin *BlogAdmin) string {
	//判断redis是否有缓存数据
	var redis_key string = "admin:navlist:" + strconv.FormatInt(admin.Id, 10)
	if res, _ := newredis.Exists(redis_key); res == true {
		res, _ := redis.String(newredis.Get(redis_key))
		return res
	} else {
		var the_html string
		//如果是超级管理员
		if admin.Role.IsSuper {
			admin_navigations := make([]*AdminNavigation, 0)
			databases.Orm.Where("is_show =?", true).Asc("sort").Find(&admin_navigations)
			the_html = BuildNavHtml(admin_navigations)
		} else {
			admin_navigations := make([]*AdminNavigation, 0)
			databases.Orm.Where("is_show =?", true).Where("id in(select admin_navigation_id from admin_role_node where admin_role_id=?)", admin.Role.Id).Asc("sort").Find(&admin_navigations)
			the_html = BuildNavHtml(admin_navigations)
		}
		//写入到Redis
		newredis.Set(redis_key, the_html, 60*60)
		return the_html
	}
}
func BuildNavHtml(admin_navigations []*AdminNavigation) (html string) {
	navigation_list := make(map[int64][]*AdminNavigation)
	for _, navigation := range admin_navigations {
		if _, ok := navigation_list[navigation.ParentId]; !ok {
			navigation_list[navigation.ParentId] = make([]*AdminNavigation, 0)
		}
		navigation_list[navigation.ParentId] = append(navigation_list[navigation.ParentId], navigation)
	}
	var the_html string = ""
	var href string = ""
	for _, navigation_one := range navigation_list[0] {
		var ico1 string = ""
		if navigation_one.Ico != "" {
			ico1 = "<img src= />"
		}
		the_html = the_html + "<div class='list-group'><h1 title='" + navigation_one.Title + "'>" + ico1 + "</h1>"
		the_html = the_html + "<div class='list-wrap'>"
		the_html = the_html + "<h2>" + navigation_one.Title + "<i></i></h2>"
		the_html = the_html + "<ul>"
		for _, navigation_two := range navigation_list[navigation_one.Id] {
			href = navigation_two.Url
			the_html = the_html + "<li>"
			the_html = the_html + "<a href='" + href + "' navid='" + strconv.FormatInt(navigation_two.Id, 10) + "' target='mainframe'>"
			the_html = the_html + "<span>" + navigation_two.Title + "</span>"
			the_html = the_html + "</a>"
			if len(navigation_list[navigation_two.Id]) > 0 {
				the_html = the_html + "<ul>"
				for _, navigation_three := range navigation_list[navigation_two.Id] {
					href = navigation_three.Url
					the_html = the_html + "<li>"
					the_html = the_html + "<a href='" + href + "' navid='" + strconv.FormatInt(navigation_three.Id, 10) + "' target='mainframe'>"
					the_html = the_html + "<span>" + navigation_three.Title + "</span>"
					the_html = the_html + "</a>"
					if len(navigation_list[navigation_three.Id]) > 0 {
						the_html = the_html + "<ul>"
						for _, navigation_four := range navigation_list[navigation_three.Id] {
							href = navigation_four.Url
							the_html = the_html + "<li>"
							the_html = the_html + "<a href='" + href + "' navid='" + strconv.FormatInt(navigation_four.Id, 10) + "' target='mainframe'>"
							the_html = the_html + "<span>" + navigation_four.Title + "</span>"
							the_html = the_html + "</a>"
							if len(navigation_list[navigation_four.Id]) > 0 {
								the_html = the_html + "<ul>"
								for _, navigation_five := range navigation_list[navigation_four.Id] {
									href = navigation_five.Url
									the_html = the_html + "<li>"
									the_html = the_html + "<a href='" + href + "' navid='" + strconv.FormatInt(navigation_five.Id, 10) + "' target='mainframe'>"
									the_html = the_html + "<span>" + navigation_five.Title + "</span>"
									the_html = the_html + "</a>"
									if len(navigation_list[navigation_five.Id]) > 0 {
										the_html = the_html + "<ul>"
										for _, navigation_six := range navigation_list[navigation_five.Id] {
											//fmt.Println("---------",navigation_five.Title,navigation_five.Node.RouteAction,navigation_five.Node.Parameter)
											href = navigation_six.Url
											the_html = the_html + "<li>"
											the_html = the_html + "<a href='" + href + "' navid='" + strconv.FormatInt(navigation_six.Id, 10) + "' target='mainframe'>"
											the_html = the_html + "<span>" + navigation_six.Title + "</span>"
											the_html = the_html + "</a>"
											if len(navigation_list[navigation_six.Id]) > 0 {
												the_html = the_html + "<ul>"
												for _, navigation_seven := range navigation_list[navigation_six.Id] {
													href = navigation_seven.Url
													//fmt.Println("---------",navigation_six.Title,navigation_six.Node.RouteAction,navigation_six.Node.Parameter)
													the_html = the_html + "<li>"
													the_html = the_html + "<a href='" + href + "' navid='" + strconv.FormatInt(navigation_seven.Id, 10) + "' target='mainframe'>"
													the_html = the_html + "<span>" + navigation_seven.Title + "</span>"
													the_html = the_html + "</a>"
												}
												the_html = the_html + "</ul>"
											}
										}
										the_html = the_html + "</ul>"
									}
								}
								the_html = the_html + "</ul>"
							}
						}
						the_html = the_html + "</ul>"
					}
				}
				the_html = the_html + "</ul>"
			}
		}
		the_html = the_html + "</ul>"
		the_html = the_html + "</div>"
		the_html = the_html + "</div>"
	}
	return the_html
}

/**
角色分配添加和编辑页面权限节点列表
*/
func NavigationRoleData() (data []AdminNavigation) {
	//判断redis是否有缓存数据
	var redis_key string = "admin:navlist:role"
	if res, _ := newredis.Exists(redis_key); res == true {
		valueBytes, _ := redis.Bytes(newredis.Get(redis_key))
		json.Unmarshal(valueBytes, &data)
	} else {
		//获取数据
		admin_navigations := make([]AdminNavigation, 0)
		databases.Orm.Asc("sort").Find(&admin_navigations)
		for k, v := range admin_navigations {
			admin_navigations[k].Title = "<span class='folder-open'></span>" + v.Title
			admin_navigation_nodes := make([]AdminNavigationNode, 0)
			databases.Orm.Cols("title", "id").Where("admin_navigation_id =?", v.Id).Find(&admin_navigation_nodes)
			admin_navigations[k].NodeTitles = "<span class='cbllist'>"
			for _, v1 := range admin_navigation_nodes {
				_id := strconv.FormatInt(v1.Id, 10)
				admin_navigations[k].NodeTitles = admin_navigations[k].NodeTitles + template.HTML("<input id='ck"+_id+"' name='route[]' type='checkbox' value="+_id+">"+"<label for='ck"+_id+"'>"+v1.Title+"</label>")
			}
			admin_navigations[k].NodeTitles = admin_navigations[k].NodeTitles + "</span>"
		}
		data = AdminNavigationMerge(admin_navigations, "<span class='folder-line'></span>", 0, 0)
		//缓存到redis
		value, _ := json.Marshal(data)
		newredis.Set(redis_key, value, 60*60)
	}
	return data
}

/**
导航添加和编辑页面父菜单选择
*/
func NavigationRoleTreeData() (data []AdminNavigation) {
	//判断redis是否有缓存数据
	var redis_key string = "admin:navlist:tree"
	if res, _ := newredis.Exists(redis_key); res == true {
		valueBytes, _ := redis.Bytes(newredis.Get(redis_key))
		json.Unmarshal(valueBytes, &data)
	} else {
		admin_navigations := make([]AdminNavigation, 0)
		databases.Orm.Asc("sort").Find(&admin_navigations)
		data = AdminNavigationMerge(admin_navigations, "|--", 0, 0)
		data = append(data, AdminNavigation{Id: 0, TitleHtml: "无父级菜单"})
		//缓存到redis
		value, _ := json.Marshal(data)
		newredis.Set(redis_key, value, 60*60)
	}
	return data
}
