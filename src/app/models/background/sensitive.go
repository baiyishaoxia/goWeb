package models

import (
	"regexp"
	"strings"
)

//region Remark:是否存在敏感词汇 [return: true存在 false不存在]  Author:tang
func IsSensitive(str string) bool {
	//数据源
	array := []string{"笨蛋", "哈哈", "傻不啦", "啥玩意"}
	//去除html标签
	str = trimHtml(str)
	for _, val := range array {
		reg := regexp.MustCompile(str)
		//将敏感词汇替换为***字符
		res := reg.ReplaceAllString(val, "***")
		if res == "***" {
			return true
		}
	}
	return false
}

//endregion

//region Remark:去除html标签 Author:tang
func trimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

//endregion
