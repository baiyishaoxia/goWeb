package controller

import (
	"app/models/background"
	"strings"
)

type Trees struct {
	Id       int64
	ParentId int64
	Level    int
	Html     string
	Title    string
}

//region Remark:树形 Author:tang
func UnlimitedForLevel2(cate *[]models.Category, html string, pid int64, level int) []Trees {
	tree := make([]Trees, 1)
	treeArray := make([]Trees, 1)
	for key, val := range *cate {
		if val.ParentId == 0 {
			tree[key].ParentId = 0
		}
		if val.ParentId == pid {
			tree[key].Id = val.Id
			tree[key].Level = level + 1
			tree[key].Html = strings.Repeat(html, level)
			tree[key].Title = tree[key].Html + tree[key].Title
			treeArray = append(treeArray, tree[key])
			UnlimitedForLevel2(cate, html, val.Id, level+1)
		}
	}
	return treeArray
}

//endregion

//region Remark:数组格式化 Author:tang
func Array2ToArray2(data []Trees) []Trees {
	arr := make([]Trees, 1)
	for key, value := range data {
		arr[key].Id = value.Id
		arr[key].Title = value.Title
	}
	return arr
}

//endregion
