package job

import (
	"app"
	"app/models"
	"databases"
	"fmt"
	"html/template"
	"time"
)

//每天自动更新新闻
var NewsChan = make(chan string, 10000)

func HandleNewsPull() {
	  db := databases.Orm.NewSession()
	  defer db.Close()
	  //读取配置
	  params:=app.NewMP()
	  url:="https://api.jisuapi.com/news/get"
	  params = map[string]interface{}{
	  	"channel":"科技", //频道
	  	"num":10,         //数量 默认10，最大40
	  	"start":models.ReadConfig("article_start"),  //起始位置，默认0 最大400
	  	"appkey":models.ReadConfig("article_appkey"), //appKey
	  }
	data,err:=app.UrlPost(url,params)
	if err != nil{
		fmt.Println("----------------------------",err.Error())
	}
	if data == nil{
		return
	}
	if data.Has("result") && data.Get("result") == nil{
		return
	}
	//开始入库
	result:=data.Get("result").(map[string]interface{})
	list,ok:=result["list"].([]interface{})
	if !ok{
		return
	}
	fmt.Println("正在更新"+ app.ToString(result["channel"]) +"类新闻，总更新数:",result["num"])
	if len(list) < 0{
		fmt.Println("-----------------------暂未拉取到新闻更新失败------------------------")
		return
	}
	//数据源
	item:=make([]models.Article,len(list))
	for key,val:=range list {
		v1:=val.(map[string]interface{})
		loc, _ := time.LoadLocation("Local")    //获取本地时区
		theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", app.ToString(v1["time"]), loc)
		item[key].Title = app.ToString(v1["title"]) //标题
		item[key].CreatedAt =  theTime			    //时间
		item[key].CateId = 15                       //科技类新闻
		item[key].Img = app.ToString(v1["pic"])     //封面图
		item[key].Content = template.HTML(app.ToString(v1["content"]))     //内容
		item[key].Source = app.ToString(v1["weburl"])                      //来源于
		item[key].Keywords = "每日头条"              //标签
		item[key].AuthorId = 2                      //作者
		item[key].ClickNum = 1000                   //浏览量
		item[key].Sort = 999                        //排序
		item[key].Status = 1                        //默认直接在本平台展示
	}
	//todo 在数据库中是否重复存在
	var removal[]interface{}
	for _,v:=range item {
		article:=new(models.Article)
		res,_err:=db.Where("title = ?",v.Title).Get(article)
		if _err != nil{
			fmt.Println("重复存在:",_err.Error())
		}
		if !res{
			removal = append(removal, v)
		}
	}
	//todo 开始批量入库
	res, err := db.Insert(removal)
	if res < 1  || err != nil{
		fmt.Println("-----------------------更新失败-------------------------------",len(removal),params["start"],err)
		return
	}
	//todo 更新配置
	config := []models.Config{
		models.Config{
			Name:  "article_start",
			Value: app.ToString(app.ToInt64(params["start"]) + app.ToInt64(params["num"])),
		},
	}
	models.SetConfig(config)
	fmt.Println("-----------------------拉取新闻更新成功------------------------",len(item))
}
