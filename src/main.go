package main

import (
	"app/grpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
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
