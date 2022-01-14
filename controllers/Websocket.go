package controllers

import (
	"MonitoringPlatform/models"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gorilla/websocket"
	"log"
)

type MyWebSocketController struct {
	beego.Controller
}

var upgrade = websocket.Upgrader{}
var (
	clients = make(map[*websocket.Conn]bool)
	//每设计有一个功能模块就添加一个channel监听信息，通过websocket发送给前端
	devBroadcast = make(chan models.DevWebsocketMsg)
	resBroadcast = make(chan models.ResWebsocketMsg)
)

func (this *MyWebSocketController) Get() {
	ws, err := upgrade.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	clients[ws] = true

	//如果从 socket 中读取数据有误，我们假设客户端已经因为某种原因断开。我们记录错误并从全局的 “clients” 映射表里删除该客户端，这样一来，我们不会继续尝试与其通信。
	//另外，HTTP 路由处理函数已经被作为 goroutines 运行。这使得 HTTP 服务器无需等待另一个连接完成，就能处理多个传入连接。
	for {
		var msg models.DevWebsocketMsg // Read in a new message as JSON and map it to a DevWebsocketMsg object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("页面可能断开啦 ws.ReadJSON error: %v", err)
			delete(clients, ws)
			break
		} else {
			fmt.Println("接受到从页面上反馈回来的信息 ", msg)
		}
	}
}
