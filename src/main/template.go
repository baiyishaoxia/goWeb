package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//加载模板
		tpl, err := template.ParseFiles("html/demo.html")
		if err != nil {
			log.Fatal(err)
			return
		}

		//执行模板
		err = tpl.Execute(w, Person{Name: "test", Age: 18})
		if err != nil {
			log.Fatal(err)
			return
		}

	})

	server := http.Server{
		Addr:        ":9090",         //监听地址和端口
		Handler:     mux,             //Handle
		ReadTimeout: 5 * time.Second, //读取超时5s
	}

	//启动监听
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
