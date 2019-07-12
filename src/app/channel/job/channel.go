package job

import "fmt"

//监听任务计划
func HandleConcurrent() {
	for {
		select {
		case userLevel := <-UsereLevelChan:
			HandleUsereLevel() //升级逻辑
			fmt.Println(userLevel)
		}
	}
}
