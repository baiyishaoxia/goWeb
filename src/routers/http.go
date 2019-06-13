package routers

import (
	"app/channel/chat"
	chatRoom "app/controllers/home/chat"
	"net/http"
)

func InitHttpRouter() {
	// Create a simple file server
	//fs := http.FileServer(http.Dir("./html/home/chat"))
	//http.Handle("/", fs)
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs)) //静态资源
	http.HandleFunc("/v2/chat", chatRoom.GetIndex)            //聊天室

	http.HandleFunc("/v2/ws", chat.HandleConnections) // Configure webSocket route
}
