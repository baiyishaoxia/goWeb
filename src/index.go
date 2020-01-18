package main

import (
	"app"
	"app/channel/chat"
	"app/channel/job"
	"app/grpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	readroute "other/reading/yournovel/routers"
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
	//当前进程中生成爬虫获取小说资源协程
	g.Go(func() error {
		return realdata().ListenAndServe()
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
	//当前进程中生成一个定时任务协程
	go job.HandleConcurrent()
	go cronData()
	//创建聊天室
	g.Go(func() error {
		routers.InitHttpRouter()

		// Start listening for incoming chat messages
		go chat.HandleMessages()
		go chat.IntoManager()                         // 启动管理员管理模块
		go chat.DataSent(&chat.Conns, chat.ToMessage) // 启动服务器广播线程

		// Start the server on localhost port 8000 and log any errors
		log.Println("http server started on :9092")
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

//加载爬虫小说
func realdata() *http.Server {
	gin.SetMode(gin.DebugMode)
	router := readroute.InitGetRouter()
	server := &http.Server{
		Addr:         ":9093",
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	fmt.Println("http://localhost" + server.Addr)
	return server
}

//启动定时任务
func cronData() {
	c := cron.New()
	spec := "0 */25 * * * ?" //(每25分钟执行)
	_err:=c.AddFunc(spec, func() {
		job.UsereLevelChan <- app.Uuid()
	})
	fmt.Println("----------------用户活跃度自动升级---------------",_err)
	_err2:=c.AddFunc("0 0 12 0 1,3,6 ?",func(){ //(每周一、周三、周六的12点更新) //0 0 12 0 1,3,6 ?
		job.NewsChan <- app.Uuid()
	})
	fmt.Println("----------------每天自动更新新闻---------------",_err2)
	//启动计划任务
	c.Start()
	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()
	select {}
}
