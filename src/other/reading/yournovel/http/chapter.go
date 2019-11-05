package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"other/reading/yournovel/service/novel"
	"other/reading/yournovel/tool"
)

func NovelChapter(c *gin.Context) {

	webSiteUrl, exist := c.GetQuery("url")
	if !exist {
		tool.ErrorResponse(c, "源网址不存在", webSiteUrl)
		return
	}
	novelName, exist := c.GetQuery("novel_name")
	if !exist {
		c.Redirect(http.StatusMovedPermanently, webSiteUrl)
		return
	}
	novelChapter, err := novel.SearchChapterOfNovel(webSiteUrl, novelName)
	if err != nil {
		fmt.Println("reading loading error:",err)
		//c.Redirect(http.StatusMovedPermanently, webSiteUrl)
		return
	}
	c.HTML(http.StatusOK, "chapter_index.html", gin.H{
		"chapter": novelChapter,
		"head": "chapter_head",
	} )
}
