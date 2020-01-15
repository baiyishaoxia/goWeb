package chat

import (
	"encoding/binary"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
	"strings"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

// 声明客户端的结构体
var Conns = make(map[string]*websocket.Conn) // 声明成功登录之后的连接对象map
var messages = make(chan UserMessageList)    // 声明消息channel(用于显示当前用户、在线用户列表)
var adminList = make(map[string]string)      // 声明管理员列表(管理员的操作)

var ToMessage = make(chan string) // 声明消息channel，用于广播(私聊、群发、禁言、踢出)

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

type UserMessageList struct {
	CurrentUser string   `json:"current_user"`  //当前用户
	UserList    []string `json:"user_list"`     //用户列表,Array
	UserListStr string   `json:"user_list_str"` //用户列表,String
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("客服连接失败.....")
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	for {
		var (
			msg    Message
			meuser UserMessageList
		)
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(Conns, msg.Username) //从用户列表删除
			delete(clients, ws)
			break
		}
		fmt.Println("【"+msg.Username+"】发送了:", msg.Message)
		//Conns["【"+msg.Username+"】来自【"+RemoteIp(r)+"】"] = ws // 记录当前用户登录状态信息
		Conns["【"+msg.Username+"】来自【中国】"] = ws // 记录当前用户登录状态信息

		meuser.CurrentUser = msg.Username
		meuser.UserListStr, meuser.UserList = userList()
		messages <- meuser // 向用户发送已在线用户列表
		//messages <- "List-"+msg.Username+"-"+userList()                         // 向用户发送已在线用户列表
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

//监听读取
func HandleMessages() {
	for {
		select {
		// Grab the next message from the broadcast channel
		case msg := <-broadcast:
			// Send it out to every client that is currently connected
			for client := range clients {
				err := client.WriteJSON(msg)
				if err != nil {
					log.Printf("error: %v", err)
					client.Close()
					delete(clients, client)
				}
				//fmt.Println(msg, client)
			}
		case user := <-messages:
			//data := strings.Split(msg,"-")                         // 聊天数据分析
			fmt.Println(user)
			for key, _ := range Conns {
				err := Conns[key].WriteJSON(user)
				if err != nil {
					log.Printf("error: %v", err)
					Conns[key].Close()
					delete(Conns, key) //当用户失去连接从用户列表删除
				}
			}
		}
	}
}

/**
 * 获取用户列表
 * @param string   用户列表
 * @param array   用户列表
 **/
func userList() (string, []string) {
	var (
		userList  string = "当前在线用户列表："
		userArray []string
	)
	for user := range Conns {
		userList += "\r\n<br/>" + user
		userArray = append(userArray, user)
	}
	return userList, userArray
}

/**
 * 管理员注册程序
 **/
func adminReg() {
	var adminname, password string
	fmt.Println("请输入管理员用户名：")
	fmt.Scanln(&adminname)
	fmt.Println("请输入管理员密码：")
	fmt.Scanln(&password)
	adminList[adminname] = password //将注册的管理员姓名密码保存到adminList中，单次启动有效
	fmt.Println("注册成功！请登录")
	adminLog() //跳转到登录
}

/**
 * 管理员登录、注册
 **/
func IntoManager() {
	fmt.Println("请输入将要进行操作：1、管理员注册 2、管理员登录")
	var input string
LOOP:
	{
		fmt.Scanln(&input)
		switch input {
		case "1":
			adminReg()
		case "2":
			adminLog()
		default:
			goto LOOP
		}
	}
	admimManager(ToMessage)
}

/**
 * 管理员登录程序
 **/
func adminLog() {
	for {
		var adminname, password string
		fmt.Println("请输入管理员用户名：")
		fmt.Scanln(&adminname)
		fmt.Println("请输入管理员密码：")
		fmt.Scanln(&password)
		if pwd, ok := adminList[adminname]; !ok {
			fmt.Println("用户名或者密码错误")
		} else {
			if pwd != password {
				fmt.Println("用户名或者密码错误！")
			} else {
				fmt.Println("登录成功！")
				break
			}
		}
	}
}

/**
* 服务器向客户端发送消息数据解析与封装操作
* @param client  客户端连接信息
* @param messages 数据通道中的数据
	 群发数据格式：普通字符串
	 单对单发送格式："To-" + uname(发送用户) + "-" + objUser（目标用户） +"-" +input
	 用户列表："List-"+objUser（目标用户）+"-"+Listinfo
	 命令:objUser + "-" + command
**/
func DataSent(conns *map[string]*websocket.Conn, messages chan string) {

	for {
		msg := <-messages

		fmt.Println(msg)

		data := strings.Split(msg, "-") // 聊天数据分析:
		length := len(data)

		if length == 2 { // 管理员单个用户发送控制命令

			(*conns)[data[0]].WriteJSON(data[1])

		} else if length == 3 { // 用户列表

			(*conns)[data[1]].WriteJSON(data[2])

		} else if length == 4 { // 向单个用户发送数据

			msg = data[1] + " say to you : " + data[3]

			(*conns)[data[2]].WriteJSON(msg)

		} else {
			// 群发
			for _, value := range *conns {
				value.WriteJSON(msg)
			}
		}
	}
}

/**
 * 发送系统通知消息
 * @param messages  channel
 * @param info   channel
 **/
func notesInfo(messages chan string, info string) {
	messages <- info
}

/**
 * 管理员管理模块
 **/
func admimManager(messages chan string) {
	for {
		var input, objUser string
		fmt.Scanln(&input)
		switch input {
		case "/to":
			fmt.Println(userList())
			fmt.Println("请输入聊天对象：")
			fmt.Scanln(&objUser)
			if _, ok := Conns[objUser]; !ok {
				fmt.Println("不存在此用户!")
			} else {
				fmt.Println("请输入消息：")
				fmt.Scanln(&input)
				notesInfo(messages, "To-Manager-"+objUser+"-"+input)
			}

		case "/all":
			fmt.Println("请输入消息：")
			fmt.Scanln(&input)
			notesInfo(messages, "Manager say : "+input)

		case "/shield":
			fmt.Println(userList())
			fmt.Println("请输入屏蔽用户名：")
			fmt.Scanln(&objUser)
			notesInfo(messages, objUser+"-/shield")
			notesInfo(messages, "用户："+objUser+"已被管理员禁言！")

		case "/remove":
			fmt.Println(userList())
			fmt.Println("请输入踢出用户名：")
			fmt.Scanln(&objUser)
			notesInfo(messages, "用户："+objUser+"已被管理员踢出聊天室！")
			if _, ok := Conns[objUser]; !ok {
				fmt.Println("不存在此用户!")
			} else {
				Conns[objUser].Close() // 删除该用户的连接
				delete(Conns, objUser) // 从已登录的列表中删除该用户
			}
		}
	}
}

/**
 * 获取用户IP
 * @param string   用户IP
 **/
const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)

func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

// Ip2long 将 IPv4 字符串形式转为 uint32
func Ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}
