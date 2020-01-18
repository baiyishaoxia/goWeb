package job

import "fmt"

//监听任务计划
func HandleConcurrent() {
	for {
		select {
		case userLevel := <-UsereLevelChan:
			HandleUsereLevel() //升级逻辑
			fmt.Println("监听管道用户活跃度自动升级定时任务:",userLevel)
		case newsList := <-NewsChan:
			HandleNewsPull()
			fmt.Println("每天自动更新新闻定时任务:",newsList)
		}
	}
}
