package main

import (
	"app/channel/chat"
	chatRoom "app/controllers/home/chat"
	"app/grpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"routers"
	"time"
)

var (
	g errgroup.Group
)

func main() {
	//当前进程中生成一个home协程
	g.Go(func() error {
		return home().ListenAndServe()
	})
	//当前进程中生成一个background协程
	g.Go(func() error {
		return background().ListenAndServe()
	})
	//当前进程中生成一个RPC通信协程
	g.Go(func() error {
		grpcInfo := new(grpc.GrpcInfo)
		rpc.Register(grpcInfo)
		listener, err := net.Listen("tcp", ":6666")
		if err != nil {
			fmt.Println("listen error:", err)
		}
		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			//新协程来处理--json
			go jsonrpc.ServeConn(conn)
		}
	})
	//创建聊天室
	g.Go(func() error {
		// Create a simple file server
		//fs := http.FileServer(http.Dir("./html/home/chat"))
		//http.Handle("/", fs)
		fs := http.FileServer(http.Dir("public"))
		http.Handle("/public/", http.StripPrefix("/public/", fs))
		http.HandleFunc("/chat", chatRoom.GetIndex) //聊天室

		// Configure webSocket route
		http.HandleFunc("/ws", chat.HandleConnections)

		// Start listening for incoming chat messages
		go chat.HandleMessages()
		go chat.IntoManager()                         // 启动管理员管理模块
		go chat.DataSent(&chat.Conns, chat.ToMessage) // 启动服务器广播线程

		// Start the server on localhost port 8000 and log any errors
		log.Println("http server started on :8000")
		err := http.ListenAndServe(":9092", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
		return err
	})
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}

//加载前台路由
func home() *http.Server {
	gin.SetMode(gin.DebugMode)
	router := routers.InitHomeRouter()
	server := &http.Server{
		Addr:         ":9090",
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	fmt.Println("http://localhost" + server.Addr)
	return server
}

//加载后台路由
func background() *http.Server {
	gin.SetMode(gin.DebugMode)
	router := routers.InitBackGroundRouter()
	server := &http.Server{
		Addr:         ":9091",
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	fmt.Println("http://localhost" + server.Addr + "/login")
	return server
}
