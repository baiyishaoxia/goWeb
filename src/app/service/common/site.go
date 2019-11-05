package common

import (
	"app"
	"app/controllers"
	"bufio"
	"config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	fileName ="./runtime/logs/web_people_count.log" //用户访问数到文件记录
)


//站点统计
func SiteStatistic(c *gin.Context)  {
	var (
		count int64 = 0                     //统计用户访问数[包括相同IP]
		real_count int64 = 0                //统计用户访问数[不包括相同IP]
		clickcount int64 =  1               //定义一个变量保存当前用户的访问次数
	)
	//假设用户访问，得到IP地址
	ip:=c.ClientIP()
	writeStr:=time.Now().Format("2006-01-02 15:04:05") +"|" + ip
	//按行读取文件
	first,_:=readLog(fileName,false,"",0)
	//判断当前有没有记录访问信息
	if first{
		//有数据，计算当前用户是第几次访问该网页
		_,clickcount =readLog(fileName,true,ip,clickcount)
		//获取已有数量
		count,real_count = getFileRows(fileName,'\n')
		writeStr = writeStr + "\n"
	}else{
		//当前用户是第一个来访问该网页
		count = 1
		real_count = 1
	}
	//写入数据到文件
	writeFile(fileName,writeStr)
	data:=map[string]int64{
		"num":count,
		"real_num":real_count,
		"user_count":clickcount,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": config.HttpSuccess,
		"msg":"Success",
		"data": data,
	})
	return
}


/* title  读取网站访问量的文件
 * @param  fileName     string  文件名称
 * @param  flag         bool    是否确认有文件数据信息
 * @param  ip    		string  ip
 * @param  clickcount   string  当前用户访问量
 */
func readLog(fileName string,flag bool,ip string,clickcount int64) (bool,int64) {
	file, err := os.Open(fileName)
	defer func(){file.Close()}()
	if err!=nil {
		if os.IsNotExist(err){
			file,_ = os.Create(fileName)
		}
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		if flag == false{
			if strings.Count(lineText,"")-1 >0{
				return true,clickcount
			}
			return false,clickcount
		}else{
			//第几次访问该网页
			item:=strings.Split(lineText,"|")
			//判读是不是当前用户查看的
			if item[1] == ip{
				//以前访问的记录与当前用户的ip相同
				clickcount ++
			}
		}
	}
	return true,clickcount
}

/* title  获取文件行数
 * @param  fileName string 文件名称
 * @param  flag     byte   分割符
 */
func getFileRows(fileName string,flag byte)  (int64,int64){
	var count int64 = 0
	var real_arr = make([]string,1)
	file,err := os.Open(fileName)
	if err != nil{
		fmt.Println(err.Error())
	}
	defer file.Close()
	fd:=bufio.NewReader(file)
	for {
		content,err := fd.ReadString(flag)
		if err!= nil{
			break
		}
		real_arr = append(real_arr, strings.Split(content,"|")[1])
		count++
	}
	fmt.Println(app.RemoveRepeatedElement(real_arr))
	return count+1,int64(len(app.RemoveRepeatedElement(real_arr)))
}

/* title  写入文件
 * @param  fileName string 文件名称
 * @param  content  string 文件内容
 */
func writeFile(fileName string,content string)  {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_,err=f.Write([]byte(content))
	}
}

//region   测试统一返回json格式   Author:tang
type SiteController struct {
  controllers.BaseController
}
func GetTest(a *gin.Context) {
	var c SiteController
	c.GinContext(a)
	c.SuccessJSON(200, "test",map[string]string{
	   "name": a.Query("name"),
	   "age":a.Query("age"),
	   "sex":a.Query("sex"),
	})
}
//endregion
