package chat

import (
	"html/template"
	"log"
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	type content struct {
		title string
	}
	//加载模板
	tpl, err := template.ParseFiles("views/home/layouts/chat/layout.html", "views/home/chat/index.html")
	if err != nil {
		log.Fatal(err)
		return
	}
	//执行模板
	err = tpl.ExecuteTemplate(w, "layout", "Hello world")
	if err != nil {
		log.Fatal(err)
		return
	}
}
